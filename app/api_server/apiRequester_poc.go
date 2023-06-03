// スキーマを返す公開メソッド。
// 呼び出しの認証はJWTが常套手段？
package main

import (
	"context"
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
)

var result interface{}

type Stat struct {
	GetStatsData struct {
		StatisticalData struct {
			DataInf struct {
				Value []struct {
					Time string `json:"@time"`
					Unit string `json:"@unit"`
					V    string `json:"$"`
				} `json:"VALUE"`
			} `json:"DATA_INF"`
		} `json:"STATISTICAL_DATA"`
	} `json:"GET_STATS_DATA"`
}

type server struct {
	pb.UnimplementedSuicideServiceServer
}

// func ApiRequester() {
func main() {

	appId := os.Getenv("ESTATAPPID")

	client := &http.Client{}

	params := "&lang=J&statsDataId=0003411679&metaGetFlg=Y&cntGetFlg=N&limit=5&explanationGetFlg=N&annotationGetFlg=N&replaceSpChars=0&sectionHeaderFlg=1&cdTab=10100&cdCat02=20X6010X840000000000"
	url := "http://api.e-stat.go.jp/rest/3.0/app/json/getStatsData?appId=" + appId + params

	resp, err := client.Get(url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")

	resp, err = client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var data Stat

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}
	// メモリでデータを共有(ごめんなさい)
	result = data.GetStatsData.StatisticalData.DataInf.Value[0].V
	fmt.Println(result)

	lis, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSuicideServiceServer(s, &server{})
	s.Serve(lis)
}

func (s *server) SuicideRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := &pb.Suicide{
		Latest: result.(string),
	}
	return &pb.Response{Suicide: result}, nil
}

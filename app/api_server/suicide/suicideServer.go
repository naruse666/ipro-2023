// スキーマを返す公開メソッド。
// 呼び出しの認証はJWTが常套手段？
package main

import (
	"context"
	"encoding/json"

	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
)

type Suicide struct {
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

func main() {
	lis, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSuicideServiceServer(s, &Suicide{})
	s.Serve(lis)
}

func (s *Suicide) SuicideRequest(ctx context.Context, in *pb.Request) (*pb.Suicide, error) {
	appId := os.Getenv("ESTATAPPID")

	client := &http.Client{}

	params := "&lang=J&statsDataId=0003411679&metaGetFlg=Y&cntGetFlg=N&limit=5&explanationGetFlg=N&annotationGetFlg=N&replaceSpChars=0&sectionHeaderFlg=1&cdTab=10100&cdCat02=20X6010X840000000000"
	url := "http://api.e-stat.go.jp/rest/3.0/app/json/getStatsData?appId=" + appId + params

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to get: %v", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to new request: %v", err)
	}

	req.Header.Set("accept", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response *pb.Suicide

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}
	
	return response, nil
}

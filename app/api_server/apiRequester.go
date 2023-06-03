// スキーマを返す公開メソッド。
// 呼び出しの認証はJWTが常套手段？
package main

import (
  //"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/grpc"
	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
)

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

//func ApiRequester() {
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

	type server struct {
		pb.unimplementedsuicideserviceserver 
	}


	func (s *server) suiciderequest(ctx context.context, in *pb.request) (*pb.response, error){
		return &pb.response{
			suicide: "12345",
		}, nil
	}

	s:= grpc.newserver()
	pb.registersuicideserviceserver(s, &server{})
//	sample := data.getstatsdata.statisticaldata.datainf.value[0]

  router := gin.default()
  router.GET("/api", func(c *gin.Context) {
  	//c.JSON(http.StatusOK, data)
  	server
  })
  router.Run(":8010")
}

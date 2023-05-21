package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/gin-gonic/gin"
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

func main() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	// set appid
	var appId string
	if env == "prod" {
		param, err := getSSMParameter("e-stat-appId")
		if err != nil {
			panic(err)
		}
		appId = param
	} else {
		appId = os.Getenv("ESTATAPPID")
	}

	//request モジュール書く
	// "https://api.e-stat.go.jp/rest/<バージョン>/app/json/getStatsData?<パラメータ群>"
	client := &http.Client{}

	params := "&lang=J&statsDataId=0003411679&metaGetFlg=Y&cntGetFlg=N&limit=5&explanationGetFlg=N&annotationGetFlg=N&replaceSpChars=0&sectionHeaderFlg=1&cdTab=10100&cdCat02=20X6010X840000000000"
	url := "http://api.e-stat.go.jp/rest/3.0/app/json/getStatsData?appId=" + appId + params
	//url := "https://api.e-stat.go.jp/rest/3.0/app/json/getStatsData?appId=" + *res.Parameter.Value + "&statsDataId=0000010201&limit=10"
	// call api
	resp, err := client.Get(url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")

	resp, err = client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	// output resp type
	//	fmt.Printf("%+v\n", body)
	// fmt.Println(string(body))

	var data Stat

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data.GetStatsData.StatisticalData.DataInf.Value[0])
	sample := data.GetStatsData.StatisticalData.DataInf.Value[0]

	//
	// html render
	//
	router := gin.Default()
	// use LoadHTMLGlob or LoadHTMLFiles function
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "Main website",
			"time":  sample.Time[:4],
			"value": sample.V,
			"unit":  sample.Unit,
		})
	})
	router.Run(":8080")
}

func getSSMParameter(paramKey string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)

	if err != nil {
		errMsg := fmt.Errorf("unable to load SDK config, %v", err)
		return "", errMsg
	}

	svc := ssm.NewFromConfig(cfg)

	// get with decryption
	res, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(paramKey),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		panic(err)
	}
	return *res.Parameter.Value, nil
}

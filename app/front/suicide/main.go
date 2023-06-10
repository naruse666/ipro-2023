package main

import (
	"context"
//	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
)

type Stat struct {
	GetStatsData *GetStatsData `json:"GET_STATS_DATA"`
}
type GetStatsData struct {
	StatisticalData *StatisticalData `json:"STATISTICAL_DATA"`
}
type StatisticalData struct {
	DataInf *DataInf `json:"DATA_INF"`
}

type DataInf struct {
	Value []*Value `json:"VALUE"`
}

type Value []struct {
	Time string `json:"@time"`
	Unit string `json:"@unit"`
	V    string `json:"$"`
}

type server struct {
	pb.UnimplementedSuicideServiceServer
}

func main() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	// request to grpc
	conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSuicideServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SuicideRequest(ctx, &pb.Request{Request: "suicide"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// r include data
	data := r.GetStatsData.StatisticalData.DataInf.Value[0]
//	fmt.Println(r)

	//
	// html render
	//
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "Main website",
			"time":  data.Time[:4],
			"value": data.V,
			"unit":  data.Unit,
		})
	})
	router.Run(":8020")
}

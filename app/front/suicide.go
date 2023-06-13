package main

import (
	"context"
	"time"
	"log"
	"net/http"

	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
  "github.com/gin-gonic/gin"
)

func suicideRequest(host string, gctx *gin.Context) {
	// request to suicide grpc
	conn, err := grpc.Dial(host + ":8010", grpc.WithInsecure(), grpc.WithBlock())
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

	data := r.GetStatsData.StatisticalData.DataInf.Value[0]

	gctx.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
		"time":  data.Time[:4],
		"value": data.V,
		"unit":  data.Unit,
	})
}

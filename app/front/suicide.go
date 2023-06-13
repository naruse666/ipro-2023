package main

import (
	"time"
	"log"
	"net/http"

  "github.com/gin-gonic/gin"
)

func suicideRequest(c *gin.Context) {
	// request to suicide grpc
	conn, err := grpc.Dial(hosts.SuicideHost + ":8010", grpc.WithInsecure(), grpc.WithBlock())
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

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
		"time":  data.Time[:4],
		"value": data.V,
		"unit":  data.Unit,
	})
}

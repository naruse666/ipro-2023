package main

import ( 
	"context"
	"os"

	"github.com/gin-gonic/gin"
	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
)


func main() {
	hosts := NewBackendHosts()

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}


	// r include data
	data := r.GetStatsData.StatisticalData.DataInf.Value[0]

	//
	// html render
	//
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", suicideRequest())
	router.Run(":8020")
}

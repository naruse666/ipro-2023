package main

import ( 
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	hosts := NewBackendHosts()

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	//
	// html render
	//
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func (c *gin.Context) {
		suicideRequest(hosts.SuicideHost, c)
	})
	router.Run(":8020")
}

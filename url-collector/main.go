package main

import (
	"os"
	"url-collector/url-collector/api"
	"url-collector/url-collector/executor"
	"url-collector/url-collector/fetcher"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	defaultPort = "8080"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	jobExec := executor.NewLimitAwareExecutor()
	nasaFetcher := fetcher.NewNasaFetcher(jobExec)

	picturesController := api.NewPicturesController(nasaFetcher)

	router.GET("/pictures", picturesController.GetImages)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.Run(":" + port)
}

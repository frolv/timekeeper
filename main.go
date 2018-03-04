package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"timekeeper/api/v1"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	apiv1.Initialize(router)

	router.Run()
}

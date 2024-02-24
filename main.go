package main

import (
	"apica/lru/lru"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	server.Use(cors.New(config))

	lruController := lru.NewLruController(lru.NewLruService())

	server.GET("/get/:key", lruController.Get)
	server.POST("/set", lruController.Set)

	server.Run(":8080")
}

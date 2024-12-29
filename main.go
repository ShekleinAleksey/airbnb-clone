package main

import (
	"airbnb-clone/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ads", handlers.GetAds)
	r.POST("/ads", handlers.CreateAd)
	r.PUT("/ads/:id", handlers.UpdateAd)
	r.DELETE("ads/:id", handlers.DeleteAd)

	r.Run(":8080")
}

package main

import (
	"airbnb-clone/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/listings", handlers.GetListing)
	r.POST("/listings", handlers.CreateListing)
	r.PUT("/listings/:id", handlers.UpdateListing)
	r.DELETE("listings/:id", handlers.DeleteListing)

	r.Run(":8080")
}

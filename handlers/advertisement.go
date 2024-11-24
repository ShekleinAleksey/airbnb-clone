package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Listing struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Price       float64 `json:"price"`
}

func GetListing(c *gin.Context) {
	fmt.Println("GetListing")
}

func CreateListing(c *gin.Context) {

}

func UpdateListing(c *gin.Context) {

}

func DeleteListing(c *gin.Context) {

}

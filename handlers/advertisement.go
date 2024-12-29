package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ad struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Price       float64 `json:"price"`
}

func GetAds(c *gin.Context) {
	var ads []Ad

	c.JSON(http.StatusOK, ads)
}

func CreateAd(c *gin.Context) {
	var ad Ad

	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(http.StatusCreated, ad)
}

func UpdateAd(c *gin.Context) {
	// id := c.Param("id")
	// var ad Ad

	// c.JSON(http.StatusOK, ad)
}

func DeleteAd(c *gin.Context) {
	// id := c.Param("id")

	// c.Status(http.StatusNoContent)
}

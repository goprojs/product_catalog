package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goprojs/product_catalog/pkg/catalog"
)

func getCakes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, catalog.Cakes)
}

func getCakeByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range catalog.Cakes {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cake not found"})
}
func postCakeByID(c *gin.Context) {
	var newCake catalog.Cake

	if err := c.BindJSON(&newCake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	catalog.Cakes = append(catalog.Cakes, newCake)

	c.IndentedJSON(http.StatusCreated, newCake)
}

// Delete a cake by ID
func deleteCakeByID(c *gin.Context) {
	id := c.Param("id")

	// Find the cake with the specified ID and remove it
	for i, item := range catalog.Cakes {
		if item.ID == id {
			// Remove the cake from the slice
			catalog.Cakes = append(catalog.Cakes[:i], catalog.Cakes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "cake deleted"})
			return
		}
	}

	// If the cake wasn't found, return a 404 error
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cake not found"})
}

func main() {
	fmt.Println("ready to go")

	// GET /cakes
	router := gin.Default()
	router.GET("/cakes", getCakes)
	router.GET("/cake/:id", getCakeByID)
	router.POST("/cakes", postCakeByID)
	router.DELETE("/cake/:id", deleteCakeByID)

	router.Run("localhost:8080")

}

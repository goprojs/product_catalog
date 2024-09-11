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

func main() {
	fmt.Println("ready to go")

	// GET /cakes
	router := gin.Default()
	router.GET("/cakes", getCakes)

	router.Run("localhost:8080")
}

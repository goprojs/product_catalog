package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goprojs/product_catalog/pkg/catalog"
	"github.com/stretchr/testify/assert"
)

// Set up the Gin router with routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/cakes", getCakes)
	router.GET("/cake/:id", getCakeByID)
	router.POST("/cakes", postCakeByID)
	router.DELETE("/cake/:id", deleteCakeByID)
	return router
}

// Mock the data to avoid using actual data from catalog
var mockCakes = []catalog.Cake{
	{ID: "1", Title: "Chocolate Cake", Description: "Delicious chocolate cake", Category: "Birthday", Price: 400, Weight: 500, Image: "http://image1.png"},
	{ID: "2", Title: "Vanilla Cake", Description: "Tasty vanilla cake", Category: "Anniversary", Price: 300, Weight: 450, Image: "http://image2.png"},
}

// Test GET /cakes
func TestGetCakes(t *testing.T) {
	// Mock the catalog.Cakes with mock data
	catalog.Cakes = mockCakes

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/cakes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)                      // Verify status code
	assert.Contains(t, w.Body.String(), "Chocolate Cake") // Verify response contains mock data
	assert.Contains(t, w.Body.String(), "Vanilla Cake")
}

// Test GET /cake/:id
func TestGetCakeByID(t *testing.T) {
	// Mock the catalog.Cakes with mock data
	catalog.Cakes = mockCakes

	router := setupRouter()

	// Test for existing cake
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/cake/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)                      // Verify status code
	assert.Contains(t, w.Body.String(), "Chocolate Cake") // Verify the correct cake is retrieved

	// Test for non-existing cake
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/cake/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)                      // Verify status code for not found
	assert.Contains(t, w.Body.String(), "cake not found")
}

// Test POST /cakes
func TestPostCakeByID(t *testing.T) {
	router := setupRouter()

	// Mock the new cake
	newCake := catalog.Cake{
		ID:          "3",
		Title:       "Strawberry Cake",
		Description: "Sweet strawberry cake",
		Category:    "Birthday",
		Price:       350,
		Weight:      600,
		Image:       "http://image3.png",
	}

	// Convert newCake to JSON
	cakeJSON, _ := json.Marshal(newCake)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/cakes", bytes.NewBuffer(cakeJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)                          // Verify status code
	assert.Contains(t, w.Body.String(), "Strawberry Cake") // Verify response contains the new cake
}

// Test DELETE /cake/:id
func TestDeleteCakeByID(t *testing.T) {
	// Mock the catalog.Cakes with mock data
	catalog.Cakes = mockCakes

	router := setupRouter()

	// Test deleting an existing cake
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/cake/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)                         // Verify status code
	assert.Contains(t, w.Body.String(), "cake deleted") // Verify cake was deleted

	// Test deleting a non-existing cake
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/cake/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)                          // Verify status code for not found
	assert.Contains(t, w.Body.String(), "cake not found") // Verify response contains correct message
}

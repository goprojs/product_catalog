package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/goprojs/product_catalog/pkg/catalog"
)

// MongoDB client
var client *mongo.Client
var cakeCollection *mongo.Collection

// Initialize MongoDB client and connect to MongoDB Atlas
func initMongoDB() error {
	// Replace with your MongoDB Atlas connection string
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return fmt.Errorf("MONGODB_URI is not set")
	}
	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB Atlas
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Connect to the "catalog" database and "cakes" collection
	cakeCollection = client.Database("catalog").Collection("cakes")
	fmt.Println("Connected to MongoDB Atlas!")
	return nil
}

func getCakes(c *gin.Context) {
	var cakes []catalog.Cake

	// Fetch all cakes from the MongoDB collection
	cursor, err := cakeCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cakes"})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a Cake struct
	for cursor.Next(context.Background()) {
		var cake catalog.Cake
		if err := cursor.Decode(&cake); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding cake data"})
			return
		}
		cakes = append(cakes, cake)
	}

	// Respond with the list of cakes
	c.IndentedJSON(http.StatusOK, cakes)
}

func getCakeByID(c *gin.Context) {
	id := c.Param("id")

	// Convert the id string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the cake with the specified ID
	var cake catalog.Cake
	err = cakeCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&cake)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "cake not found"})
		return
	}

	// Respond with the cake data
	c.IndentedJSON(http.StatusOK, cake)
}

func postCakeByID(c *gin.Context) {
	var newCake catalog.Cake

	// Bind the JSON body to the newCake struct
	if err := c.BindJSON(&newCake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the new cake into the MongoDB collection
	result, err := cakeCollection.InsertOne(context.Background(), newCake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting cake"})
		return
	}

	// Respond with the inserted document's ID
	c.IndentedJSON(http.StatusCreated, gin.H{"insertedID": result.InsertedID})
}

// Delete a cake by ID
func deleteCakeByField(c *gin.Context) {
	// Get the field and value from query parameters
	field := c.Query("field") // Field to filter on (e.g., "id" or "name")
	value := c.Query("value") // Value of the field to delete

	if field == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field and value parameters are required"})
		return
	}

	// Create a filter based on the specified field and value
	filter := bson.M{field: value}

	// Delete the cake with the specified field and value
	result, err := cakeCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting cake: %+v", err)})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cake not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Cake deleted"})
}

func main() {
	// Initialize MongoDB client with MongoDB Atlas
	if err := initMongoDB(); err != nil {
		fmt.Println("Failed to connect to MongoDB Atlas:", err)
		return
	}
	defer client.Disconnect(context.Background()) // Ensure the client disconnects when the application closes

	// Create Gin router and define routes
	router := gin.Default()
	router.GET("/cakes", getCakes)
	router.GET("/cake/:id", getCakeByID)
	router.POST("/cakes", postCakeByID)
	router.DELETE("/cake", deleteCakeByField) // Updated route for deleting by field

	// Start the server on localhost:8080
	router.Run("localhost:8080")
}

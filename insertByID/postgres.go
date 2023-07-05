package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add more fields as needed
}

func setupDB() (*sql.DB, error) {
	// Replace the connection parameters with your own
	db, err := sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func insertDataByID(c *gin.Context) {
	// Extract the ID from the URL parameter
	id := c.Param("id")

	// Parse the ID into an integer
	dataID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Parse the request body into a Data struct
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get a database connection
	db, err := setupDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Perform the insertion query
	query := "INSERT INTO your_table (id, name) VALUES ($1, $2)"
	_, err = db.Exec(query, dataID, data.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}
func main() {
	r := gin.Default()

	// Register the route
	r.POST("/data/:id", insertDataByID)

	// Start the server
	r.Run(":8080")
}

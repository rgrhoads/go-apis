package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "fmt"
type TestData struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("/", welcome)
	router.GET("/test-data", testData)

	router.Run(":3773")
}

func welcome(c *gin.Context) {
	log.Printf("Welcome to Modren APIs!")
}

func testData(c *gin.Context) {
	response := TestData{"12345", "Hello World!"}

	c.JSON(http.StatusOK, response)
}

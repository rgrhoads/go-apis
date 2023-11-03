package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// import "fmt"

func main() {
	router := gin.Default()

	router.GET("/", welcome)

	router.Run(":3773")
}

func welcome(c *gin.Context) {
	log.Printf("Welcome to Modren APIs!")
}

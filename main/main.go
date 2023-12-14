package main

import (
	"go-apis/rest/router"
)

type TestData struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := router.New()

	router.Run(":3773")
}

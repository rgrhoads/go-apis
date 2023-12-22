package main

import (
	"go-apis/rest/router"
)

func main() {
	router := router.New()

	router.Run(":3773")
}

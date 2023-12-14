package scraper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScrapeMessage(ctx *gin.Context) {
	msg := "Web Scraper Under Constrction..."
	response := gin.H{
		"success": "true",
		"message": msg,
	}
	ctx.JSON(http.StatusOK, response)
}

func ScrapeWebsite(ctx *gin.Context) {
	website := ctx.Param("website")
	fmt.Printf("Scraping the Website: %s", website)
}

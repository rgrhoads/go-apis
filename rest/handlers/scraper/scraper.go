package scraper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
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
	// website := ctx.Param("website")
	// fmt.Printf("Scraping the Website: %s", website)

	url := "https://www.trackingdifferences.com/ETF/ISIN/IE00B1XNHC34"

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

	c.OnHTML("h1.page-title", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit(url)
}

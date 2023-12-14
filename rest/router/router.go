package router

import (
	"go-apis/rest/handlers/scraper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()

	group := router.Group("/")
	webScraperGroup := router.Group("/scrape")

	setDefaultHandlers(group)
	setWebScraperHandlers(webScraperGroup)

	return router
}

func setDefaultHandlers(group *gin.RouterGroup) {
	group.GET("", welcomeMessage)
	group.GET("/test-data", testData)
}

func setWebScraperHandlers(group *gin.RouterGroup) {
	group.GET("", scraper.ScrapeMessage)
	group.GET(":website", scraper.ScrapeWebsite)
}

func testData(ctx *gin.Context) {
	testData := []int{1, 2, 3, 4, 5}
	response := gin.H{
		"success": "true",
		"data":    testData,
	}

	ctx.JSON(http.StatusOK, response)
}

func welcomeMessage(ctx *gin.Context) {
	wMsg := "Welcome to the Modren GO API Microservice!"
	log.Print(wMsg)

	response := gin.H{
		"success": "true",
		"message": wMsg,
	}
	ctx.JSON(http.StatusOK, response)
}

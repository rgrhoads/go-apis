package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-apis/rest/handlers/scraper"
	"go-apis/rest/handlers/websocket"
)

func New() *gin.Engine {
	router := gin.Default()

	group := router.Group("/")
	webScraperGroup := router.Group("/scrape")
	webSocketGroup := router.Group("/ws")

	setDefaultHandlers(group)
	setWebScraperHandlers(webScraperGroup)
	setWebSocketHandlers(webSocketGroup)

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

func setWebSocketHandlers(group *gin.RouterGroup) {
	group.GET("/message", websocket.SendMessage)
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

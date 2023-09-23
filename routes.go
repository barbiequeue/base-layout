package main

import "github.com/gin-gonic/gin"

func setRoutes(engine *gin.Engine) {
	engine.POST("/webhook", webhookHandler)
}

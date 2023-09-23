package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func webhookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

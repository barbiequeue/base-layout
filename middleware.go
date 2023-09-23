package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !conf.useWebhookSecureKey.Bool() {
			c.Next()
		}

		reqSecKey := c.Request.Header.Get("X-Secure-Key")
		if reqSecKey != conf.webhookSecureKey.String() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "wrong secure header",
			})
			return
		}

		c.Next()
	}
}

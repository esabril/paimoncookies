package middleware

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckApiAppKey middleware for API Request protection
func CheckApiAppKey(conf *service.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		appKey, _ := c.GetQuery("appKey")

		if appKey != conf.Api.AppKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized request",
			})

			return
		}

		c.Next()
	}
}

package server

import (
	"github.com/esabril/paimoncookies/internal/api"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetRouter Configure and handle http-requests
func GetRouter(s *service.Service) *gin.Engine {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Printf("Error occurred while setting trusted proxies to Gin Routed: %s\n", err.Error())
	}

	configRoutes(r, s)

	return r
}

// Routes for application
func configRoutes(r *gin.Engine, s *service.Service) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", handle(api.Health, s))

	ag := r.Group("/api")
	v1 := ag.Group("/v1")

	v1.GET("/agenda", handle(api.GetAgenda, s))
}

// Auxiliary function for setting additional arguments to API handlers
func handle(f func(c *gin.Context, s *service.Service), s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c, s)
	}
}

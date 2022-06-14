package server

import (
	"github.com/esabril/paimoncookies/internal/api"
	"github.com/esabril/paimoncookies/internal/server/middleware"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetRouter Configure and handle http-requests
func GetRouter(s *service.Service, c *service.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	if !c.Api.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		r.Use(gin.Logger())
	}

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Printf("Error occurred while setting trusted proxies to Gin Routed: %s\n", err.Error())
	}

	configRoutes(r, s, c)

	return r
}

// Routes for application
func configRoutes(r *gin.Engine, s *service.Service, c *service.Config) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", handle(api.Health, s))

	ag := r.Group("/api")
	ag.Use(middleware.CheckApiAppKey(c))

	v1 := ag.Group("/v1")
	v1.GET("/agenda", handle(api.GetAgenda, s))
	v1.GET("/characters", handle(api.GetCharacters, s))
}

// Auxiliary function for setting additional arguments to API handlers
func handle(f func(c *gin.Context, s *service.Service), s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		s.SetTodayWeekday()
		f(c, s)
	}
}

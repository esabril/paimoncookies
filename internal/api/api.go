package api

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

// Health Application healthcheck
func Health(c *gin.Context, _ *service.Service) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAgenda Agenda: schedule of day's resources and other game statuses
func GetAgenda(c *gin.Context, s *service.Service) {
	today := strings.ToLower(time.Now().Weekday().String())
	world, err := s.World.CreateAgenda(today)

	if err != nil {
		log.Printf("Error occurred while getting Agenda")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения расписания для сегодняшнего дня",
		})
	}

	c.JSON(http.StatusOK, world)
}

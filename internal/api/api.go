package api

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Health Application healthcheck
func Health(c *gin.Context, _ *service.Service) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetAgenda Agenda: schedule of day's resources and other game statuses
func GetAgenda(c *gin.Context, s *service.Service) {
	world, err := s.World.CreateAgenda(s.TodayWeekday)

	if err != nil {
		log.Printf("Error occurred while getting Agenda: %s\n", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения расписания для сегодняшнего дня",
		})

		return
	}

	c.JSON(http.StatusOK, world)
}

// GetCharacters list of characters and elements
func GetCharacters(c *gin.Context, s *service.Service) {
	elements, characters, err := s.Characters.GetInitialCharactersList()

	if err != nil {
		log.Println("Error occurred while getting Characters list:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения списка персонажей",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"elements":   elements,
		"characters": characters,
	})
}

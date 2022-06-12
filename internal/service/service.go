package service

import (
	"fmt"
	"github.com/esabril/paimoncookies/internal/service/world"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"time"
)

// Service common application service
type Service struct {
	TodayWeekday string
	World        *world.World
}

// NewService creates new service
func NewService(c *Config) *Service {
	db, err := sqlx.Open(
		c.Database.DriverName,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s database=%s sslmode=disable",
			c.Database.Host,
			c.Database.Port,
			c.Database.Username,
			c.Database.Password,
			c.Database.Database,
		),
	)

	if err != nil {
		log.Fatalf("Unable to connect to database: %s\n", err.Error())
	}

	return &Service{
		TodayWeekday: strings.ToLower(time.Now().Weekday().String()),
		World:        world.NewService(db),
	}
}

package service

import (
	"fmt"
	"github.com/esabril/paimoncookies/internal/service/characters"
	"github.com/esabril/paimoncookies/internal/service/world"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"time"
)

// Service common application service
type Service struct {
	Config       *Config
	TodayWeekday string
	World        *world.World
	Characters   *characters.Characters
}

// NewService creates new service
func NewService(c *Config) *Service {
	db, err := sqlx.Open(
		c.Database.DriverName,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s database=%s sslmode=%s",
			c.Database.Host,
			c.Database.Port,
			c.Database.Username,
			c.Database.Password,
			c.Database.Database,
			c.Database.SslMode,
		),
	)

	if err != nil {
		log.Fatalf("Unable to connect to database: %s\n", err.Error())
	}

	return &Service{
		Config: c,
		World:  world.NewService(db),
		Characters:   characters.NewService(db),
	}
}

// SetTodayWeekday refresh and set today weekday for requests
func (s *Service) SetTodayWeekday() {
	location, err := time.LoadLocation(s.Config.Timezone)
	if err != nil {
		log.Printf("Error while loading timezone: %s\n", err.Error())
	}

	now := time.Now()

	s.TodayWeekday = strings.ToLower(now.In(location).Weekday().String())
}

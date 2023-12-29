package service

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/esabril/paimoncookies/internal/service/archive"
	"github.com/esabril/paimoncookies/internal/service/characters"
	"github.com/esabril/paimoncookies/internal/service/world"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Service common application service
type Service struct {
	Config       *Config
	TodayWeekday string
	World        *world.World
	Characters   *characters.Characters
	Archive      *archive.Archive
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

	w := world.NewService(db)
	ch := characters.NewService(db)

	return &Service{
		Config:     c,
		World:      w,
		Characters: ch,
		Archive:    archive.NewArchive(w, ch),
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

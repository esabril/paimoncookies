package commands

import (
	"log"
)

const CommandAgenda = "agenda"

// GetAgenda Render agenda message
// TODO: Sometimes for specific user
func (c *Commander) GetAgenda() string {
	agenda, err := c.service.World.CreateAgenda(c.service.TodayWeekday)
	if err != nil {
		log.Printf("Unable to create Agenda: %s\n", err.Error())

		return c.renderer.RenderError("расписание дня")
	}

	result, err := c.renderer.Render("agenda.tpl", agenda)
	if err != nil {
		log.Printf("Unable to render Agenda template: %s\n", err.Error())

		return c.renderer.RenderError("расписание дня")
	}

	return result
}

package commands

import (
	"github.com/esabril/paimoncookies/internal/service/world/model"
	"log"
	"strings"
)

func (c *Commander) GetGemInfo(gem model.Gem) string {
	gem.Title = c.renderer.AddEmojiToGem(gem.Title)

	result, err := c.renderer.Render("gem.tpl", gem)
	if err != nil {
		log.Printf("Unable to render Gem template: %s\n", err.Error())

		return c.renderer.RenderError(gem.Title)
	}

	return result
}

// We can have three types of Gem message
// - {emoji} Нефрит Шивада
// - Нефрит Шивада
// - нефрит (case insensitive)
func (c *Commander) isGem(reply string) (model.Gem, bool) {
	parts := strings.Split(reply, " ")
	gemName := ""

	switch len(parts) {
	case 1, 2:
		gemName = reply
	case 3:
		gemName = strings.Join(parts[1:3], " ")
	default:
		return model.Gem{}, false
	}

	gem, err := c.service.World.FindGemByTitle(gemName)
	if err != nil {
		return model.Gem{}, false
	}

	return gem, true
}

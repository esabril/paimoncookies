package commands

import (
	"fmt"
	"log"
	"strings"
)

func (c *Commander) GetCharacterMenuRules(element string) string {
	showExtendedRules := false

	if element != "" {
		element = fmt.Sprintf("%s", c.renderer.AddEmojiToElement(element))
	} else {
		showExtendedRules = true
	}

	params := struct {
		Element           string
		ShowExtendedRules bool
	}{
		Element:           element,
		ShowExtendedRules: showExtendedRules,
	}

	result, err := c.renderer.Render("character_rules.tpl", params)
	if err != nil {
		log.Printf("Unable to render Character Rules template: %s\n", err.Error())

		return "Я не могу напомнить тебе правила поиска персонажей... но ты ведь и так их помнишь, правда?"
	}

	return result
}

func (c *Commander) GetCharacterInfo(name string) (result string, element string) {
	character, err := c.service.Archive.GetCharacterInfo(name)
	if err != nil {
		return c.renderer.RenderError(fmt.Sprintf("информацию о персонаже %s", name)), ""
	}

	element = character.Element
	character.Element = c.renderer.GetEmojiToElement(character.Element)
	today := c.service.World.GetWeekdayTranslation(c.service.TodayWeekday)

	for i, wd := range character.Materials.TalentBook.Weekdays {
		if wd == today {
			character.Materials.TalentBook.Weekdays[i] = fmt.Sprintf("📍 *%s*", wd)
			break
		}
	}

	result, err = c.renderer.Render("character.tpl", character)
	if err != nil {
		log.Printf("Unable to render Character template: %s\n", err.Error())

		return c.renderer.RenderError(fmt.Sprintf("информацию о персонаже %s", name)), element
	}

	return result, element
}

func (c *Commander) isCharacter(reply string) bool {
	return c.service.Characters.CheckCharacter(reply)
}

func (c *Commander) isElement(reply string) bool {
	return c.service.Characters.CheckElement(c.getElementFromReply(reply))
}

// We can have four types of reply message:
// - "🔥 Пиро ➡"️([element emoji, space, element name, space, arrow])
// - "⬅ 🔥 Пиро" ([arrow, space, element emoji, space, element name])
// - "🔥 Пиро" (element emoji, space, element name])
// - Any text message. Simple "Крио" will give us valid element checking
func (c *Commander) getElementFromReply(reply string) string {
	data := strings.Split(reply, " ")

	// We possibly have case "🔥 Пиро"
	if len(data) == 2 {
		return data[1]
	}

	if len(data) == 3 {
		if data[0] == c.renderer.PreviousPageEmoji {
			// check case "⬅ 🔥 Пиро"
			return data[2]
		} else {
			// check case "🔥 Пиро ➡"
			return data[1]
		}
	}

	// Another message
	return reply
}

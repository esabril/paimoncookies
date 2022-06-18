package commands

import "log"

const CommandStart = "start"

func (c *Commander) GetStart() string {
	result, err := c.renderer.Render("start.tpl", nil)
	if err != nil {
		log.Printf("Unable to render Start template: %s\n", err.Error())

		return c.renderer.RenderError("приветствие")
	}

	return result
}

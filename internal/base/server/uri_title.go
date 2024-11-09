package server

import (
	"github.com/gofiber/fiber/v2"
	"shortlink/internal/base"
	"shortlink/internal/base/toolkit"
)

type UriTitleApi struct {
}

func NewUriTitleApi(router fiber.Router) {
	server := &UriTitleApi{}

	prefix := base.GetConfig().Server.BasePath

	router.Get(prefix+"/get-title", server.GetTitleByUrl)
}

func (h UriTitleApi) GetTitleByUrl(c *fiber.Ctx) (err error) {
	url := c.Query("url")

	var title string
	if title, err = toolkit.GetTitleByUrl(url); err != nil {
		return err
	}
	err = c.JSON(fiber.Map{
		"title": title,
	})
	return
}

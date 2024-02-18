package server

import (
	"url-shortener/internal/utils"
	"url-shortener/internal/views"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HomeView)
	s.App.Get("/health", s.healthHandler)

}

func (s *FiberServer) HomeView(c *fiber.Ctx) error {
	return utils.Render(c, views.HomeView())
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

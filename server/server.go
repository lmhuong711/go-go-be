package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/lmhuong711/go-go-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("go-go-be")
	prometheus.RegisterAt(app, "/metrics")

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(prometheus.Middleware)

	app.Use(func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})

	api := app.Group("/api/v1")
	routes.SetupRoutes(api)

	return app
}

package main

import (
	"github.com/amedinac/jobs/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")


	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "Job Search Project!! 😉",
		})
	})

	routes.ApplicantRoute(api.Group("/applicants"))
}





func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)


	

	app.Listen(":3000")
}
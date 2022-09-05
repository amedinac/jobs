package routes

import (
	"github.com/amedinac/jobs/controllers"
	"github.com/gofiber/fiber/v2"
)

func ApplicantRoute(route fiber.Router) {
	route.Get("", controllers.GetApplicants)
}


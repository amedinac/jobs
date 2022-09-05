package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Applicant struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var applicants = []*Applicant{
	{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@mail.com",
		Phone:     "1234567890",
	},
	{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@mail.com",
		Phone:     "1234567890",
	},
}

func GetApplicants(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{	
		"success":  true,
		"data": fiber.Map{	
			"applicants": applicants,
			},
		})
}



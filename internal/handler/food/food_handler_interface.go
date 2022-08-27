package food

import "github.com/gofiber/fiber/v2"

type FoodHandler interface {
	GetAll(c *fiber.Ctx) error
	FindMatches(c *fiber.Ctx) error
}

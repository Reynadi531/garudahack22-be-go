package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	GetProfile(c *fiber.Ctx) error
}

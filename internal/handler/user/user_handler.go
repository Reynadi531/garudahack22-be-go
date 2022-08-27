package user

import (
	"gh22-go/internal/entities/web"
	userService "gh22-go/internal/service/user"
	"gh22-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService userService.UserService
}

func (u userHandler) GetProfile(c *fiber.Ctx) error {
	token, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Unable to parse the token",
		})
	}

	result := u.userService.GetProfile(token.UserID.String())
	return c.Status(result.StatusCode).JSON(result)
}

func NewUserHandler(userService userService.UserService) UserHandler {
	return &userHandler{userService: userService}
}

package user

import (
	"gh22-go/internal/entities/web"
	userRepo "gh22-go/internal/repository/user"
	"github.com/gofiber/fiber/v2"
)

type userService struct {
	userRepo userRepo.UserRepository
}

func (u userService) GetProfile(userID string) web.Response {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Error:      err,
		}
	}

	if user.AppToken == "" {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Error:      "User not found",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Success get profile",
		Data: map[string]interface{}{
			"name":      user.Name,
			"email":     user.Email,
			"app_token": user.AppToken,
		},
		Error: nil,
	}
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

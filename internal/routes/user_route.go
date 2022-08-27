package routes

import (
	"gh22-go/internal/handler/user"
	userRepo "gh22-go/internal/repository/user"
	userService "gh22-go/internal/service/user"
	"gh22-go/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	userGroup := app.Group("/api/v1/user")

	userGroup.Get("/profile", middleware.JWTProtected(), userHandler.GetProfile)
}

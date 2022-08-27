package routes

import (
	authHandler "gh22-go/internal/handler/auth"
	tokenRepo "gh22-go/internal/repository/token"
	userRepo "gh22-go/internal/repository/user"
	authService "gh22-go/internal/service/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	tokenRepo := tokenRepo.NewTokenRepository(db)
	authService := authService.NewAuthService(userRepo, tokenRepo)
	authHandler := authHandler.NewAuthHandler(authService)

	authGroup := app.Group("/api/v1/auth")

	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/refresh", authHandler.Refresh)
}

package routes

import (
	food2 "gh22-go/internal/handler/food"
	"gh22-go/internal/repository/food"
	foodService2 "gh22-go/internal/service/food"
	"github.com/gofiber/fiber/v2"
)

func RegisterFoodRoutes(app *fiber.App, foods []string) {
	foodRepo := food.NewFoodRepository(foods)
	foodService := foodService2.NewFoodService(foodRepo)
	foodHandler := food2.NewFoodHandler(foodService)

	foodGroup := app.Group("/api/v1/")
	foodGroup.Get("/foods", foodHandler.GetAll)
	foodGroup.Get("/foods/search", foodHandler.FindMatches)
}

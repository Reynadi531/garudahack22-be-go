package food

import (
	foodService "gh22-go/internal/service/food"
	"github.com/gofiber/fiber/v2"
)

type foodHandler struct {
	foodService foodService.FoodService
}

func (f foodHandler) GetAll(c *fiber.Ctx) error {
	result := f.foodService.GetAll()
	return c.Status(result.StatusCode).JSON(result)
}

func (f foodHandler) FindMatches(c *fiber.Ctx) error {
	result := f.foodService.FindMatches(c.Query("q"))
	return c.Status(result.StatusCode).JSON(result)
}

func NewFoodHandler(foodService foodService.FoodService) FoodHandler {
	return &foodHandler{foodService: foodService}
}

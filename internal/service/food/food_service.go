package food

import (
	"gh22-go/internal/entities/web"
	foodRepo "gh22-go/internal/repository/food"
	"github.com/gofiber/fiber/v2"
)

type foodService struct {
	foodRepo foodRepo.FoodRepository
}

func (f foodService) GetAll() web.Response {
	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Success get all foods",
		Data:       f.foodRepo.GetAll(),
		Error:      nil,
	}
}

func (f foodService) FindMatches(s string) web.Response {
	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Success get all foods",
		Data:       f.foodRepo.FindMatches(s),
		Error:      nil,
	}
}

func NewFoodService(foodRepo foodRepo.FoodRepository) FoodService {
	return &foodService{foodRepo: foodRepo}
}

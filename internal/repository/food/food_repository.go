package food

import "github.com/lithammer/fuzzysearch/fuzzy"

type FoodRepository interface {
	GetAll() []string
	FindMatches(string) []string
}

type foodRepository struct {
	foods []string
}

func (f foodRepository) GetAll() []string {
	return f.foods
}

func (f foodRepository) FindMatches(s string) []string {
	return fuzzy.Find(s, f.foods)
}

func NewFoodRepository(foods []string) FoodRepository {
	return &foodRepository{foods: foods}
}

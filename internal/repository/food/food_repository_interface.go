package food

type FoodRepository interface {
	GetAll() []string
	FindMatches(string) []string
}

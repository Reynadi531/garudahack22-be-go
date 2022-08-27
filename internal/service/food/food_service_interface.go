package food

import "gh22-go/internal/entities/web"

type FoodService interface {
	GetAll() web.Response
	FindMatches(s string) web.Response
}

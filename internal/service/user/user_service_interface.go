package user

import "gh22-go/internal/entities/web"

type UserService interface {
	GetProfile(userID string) web.Response
}

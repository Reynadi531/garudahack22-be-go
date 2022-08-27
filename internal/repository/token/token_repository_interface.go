package token

import (
	"gh22-go/internal/entities/model"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Create(userID uuid.UUID, token string) error
	GetActiveToken(userID uuid.UUID) (model.Token, error)
	Revoke(userID uuid.UUID) error
}

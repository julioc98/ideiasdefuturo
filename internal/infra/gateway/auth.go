// Package gateway ...
package gateway

import (
	"github.com/google/uuid"
	"github.com/julioc98/ideiasdefuturo/internal/domain"
)

// Auth ....
type Auth struct{}

// GenerateToken ....
func (a *Auth) GenerateToken(user *domain.User) (string, error) {
	return uuid.NewString(), nil
}

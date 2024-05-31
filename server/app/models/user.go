package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID                int32
	Phone             string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type UserClaims struct {
	Sub int32
	jwt.RegisteredClaims
}

func (user *User) GenerateJwtClaims() (claims jwt.Claims) {
	return UserClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

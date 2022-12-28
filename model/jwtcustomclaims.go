package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTCustomClaims struct {
	UserID  uuid.UUID
	Email   string
	IsAdmin bool
	jwt.StandardClaims
}

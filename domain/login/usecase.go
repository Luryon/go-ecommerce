package login

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/luryon/go-ecommerce/model"
	"time"
)

type Login struct {
	useCaseUser UseCaseUser
}

func New(uc UseCaseUser) Login {
	return Login{uc}
}

func (l Login) Login(email, password, jwtSecretKey string) (model.User, string, error) {
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}

	claims := model.JWTCustomClaims{
		UserID:  user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	data, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "toekn.SignedString()", err)
	}

	user.Password = ""

	return user, data, nil
}

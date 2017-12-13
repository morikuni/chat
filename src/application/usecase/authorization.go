package usecase

import (
	"context"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/morikuni/chat/src/application/dto"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/pkg/errors"
)

type Authorization interface {
	GenerateAccessToken(ctx context.Context, email, password string) (dto.AccessToken, error)
}

func NewAuthorization(
	authentication Authentication,
) Authorization {
	return authorization{
		authentication,
	}
}

type authorization struct {
	authentication Authentication
}

type token struct {
	jwt.StandardClaims
	UserID model.UserID `json:"user_id"`
}

func (a authorization) GenerateAccessToken(ctx context.Context, email, password string) (dto.AccessToken, error) {
	userID, err := a.authentication.Login(ctx, email, password)
	if err != nil {
		return dto.AccessToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		UserID: userID,
	})
	at, err := token.SignedString([]byte("secret"))
	if err != nil {
		return dto.AccessToken{}, errors.Wrap(err, "cannot generate access token")
	}
	return dto.AccessToken{at}, nil
}

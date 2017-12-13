package usecase

import (
	"context"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/morikuni/chat/src/application"
	"github.com/morikuni/chat/src/application/dto"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/infra"
	"github.com/pkg/errors"
)

type Authorization interface {
	GenerateAccessToken(ctx context.Context, email, password string) (dto.AccessToken, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (model.UserID, error)
}

func NewAuthorization(
	authentication Authentication,
	clock infra.Clock,
) Authorization {
	return authorization{
		authentication,
		clock,
	}
}

type authorization struct {
	authentication Authentication
	clock          infra.Clock
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
			ExpiresAt: a.clock.Now().Add(time.Hour).Unix(),
		},
		UserID: userID,
	})
	at, err := token.SignedString([]byte("secret"))
	if err != nil {
		return dto.AccessToken{}, errors.Wrap(err, "cannot generate access token")
	}
	return dto.AccessToken{at}, nil
}

func (a authorization) ValidateAccessToken(ctx context.Context, accessToken string) (model.UserID, error) {
	var token token
	_, err := jwt.ParseWithClaims(accessToken, &token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("access token has invalid method: %v", t.Method.Alg())
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", application.RaiseInvalidCredentialError(err.Error())
	}
	return token.UserID, nil
}

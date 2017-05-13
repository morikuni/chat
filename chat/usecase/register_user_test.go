package usecase

import (
	"context"
	"testing"

	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	repo := user.NewRepository()
	ru := NewRegisterUser(repo)

	id, err := ru.Register(ctx, "mario", "me@email.mail", "password")
	assert.Nil(err)
	assert.NotEmpty(id)

	u, err := repo.Find(ctx, id)

	assert.Nil(err)
	assert.EqualValues("mario", u.Name())
	assert.Nil(u.Authenticate("me@email.mail", "password"))
}

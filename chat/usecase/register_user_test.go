package usecase

import (
	"testing"

	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)

	repo := user.NewRepository()
	ru := NewRegisterUser(repo)

	id, err := ru.Register("mario", "me@email.mail", "password")
	assert.Nil(err)
	assert.NotEmpty(id)

	u, err := repo.Find(id)

	assert.Nil(err)
	assert.EqualValues("mario", u.Name())
	assert.Nil(u.Authenticate("me@email.mail", "password"))
}

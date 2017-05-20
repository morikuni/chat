package user

import (
	"testing"

	"github.com/morikuni/chat/chat/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)

	name := model.UserName("user")
	email := model.Email("email@email.com")
	raw := model.Password("password")
	user := New(name, email, raw)
	assert.Regexp("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", user.ID())
	assert.Equal(name, user.Name())
	assert.Equal(uint64(1), user.Version())

	assert.Nil(user.Authenticate(email, raw))
	assert.NotNil(user.Authenticate(model.Email("hoge"), raw))
	assert.NotNil(user.Authenticate(email, model.Password("hoge")))

	user.UpdateProfile(model.UserName("updated"))
	assert.Equal(model.UserName("updated"), user.Name())
	assert.Equal(uint64(2), user.Version())
}

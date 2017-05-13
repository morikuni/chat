package user

import (
	"testing"

	"github.com/morikuni/chat/chat/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestHashedPassword(t *testing.T) {
	assert := assert.New(t)

	raw := model.Password("password")
	pass := newHashedPassword(raw)
	assert.True(pass.Equal(raw))
}

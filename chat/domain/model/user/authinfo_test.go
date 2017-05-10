package user

import (
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashedPassword(t *testing.T) {
	assert := assert.New(t)

	raw := model.Password("password")
	pass := newHashedPassword(raw)
	assert.True(pass.Equal(raw))
}

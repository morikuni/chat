package model

import (
	"testing"

	"github.com/morikuni/chat/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert := assert.New(t)

	password, verr := ValidatePassword("")
	assert.Equal(domain.RaiseValidationError("password", "cannot be empty"), verr)
	assert.Zero(password)

	password, verr = ValidatePassword("password")
	assert.NoError(verr)
	hash, err := password.Hash()
	assert.NoError(err)

	tester, verr := ValidatePassword("hogehoge")
	assert.NoError(verr)
	assert.Error(hash.Equal(tester))

	tester, err = ValidatePassword("password")
	assert.NoError(err)
	assert.NoError(hash.Equal(tester))
}

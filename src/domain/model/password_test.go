package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert := assert.New(t)

	password, verr := ValidatePassword("")
	assert.EqualError(verr, "cannot be empty")
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

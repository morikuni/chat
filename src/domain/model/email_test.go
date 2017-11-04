package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	email, verr := ValidateEmail("")
	assert.EqualError(verr, "invalid format")
	assert.Zero(email)

	email, verr = ValidateEmail("@email@mail.com")
	assert.EqualError(verr, "invalid format")
	assert.Zero(email)

	email, verr = ValidateEmail("email@mail.com")
	assert.NoError(verr)
	assert.EqualValues("email@mail.com", email)
}

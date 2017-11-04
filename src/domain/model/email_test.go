package model

import (
	"testing"

	"github.com/morikuni/chat/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	email, verr := ValidateEmail("")
	assert.Equal(domain.RaiseValidationError("invalid format"), verr)
	assert.Zero(email)

	email, verr = ValidateEmail("@email@mail.com")
	assert.Equal(domain.RaiseValidationError("invalid format"), verr)
	assert.Zero(email)

	email, verr = ValidateEmail("email@mail.com")
	assert.NoError(verr)
	assert.EqualValues("email@mail.com", email)
}

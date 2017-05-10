package password

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	password, err := Validate("aaaaaa")
	assert.EqualValues("aaaaaa", password)
	assert.Nil(err)

	password, err = Validate("aaaaa")
	assert.EqualValues("", password)
	assert.EqualError(errors.Cause(err), "password must be longer than 5 characters")
}

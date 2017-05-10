package email

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	email, err := Validate("ho-ge.bar+baz1@foo2.com")
	assert.EqualValues("ho-ge.bar+baz1@foo2.com", email)
	assert.Nil(err)

	email, err = Validate("@foo.com")
	assert.EqualValues("", email)
	assert.EqualError(errors.Cause(err), "invalid format")

	email, err = Validate("foo.com@")
	assert.EqualValues("", email)
	assert.EqualError(errors.Cause(err), "invalid format")

	email, err = Validate("")
	assert.EqualValues("", email)
	assert.EqualError(errors.Cause(err), "invalid format")
}

package taskqueue

import (
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	type Input struct {
		Event event.Event
	}
	type Expect struct {
	}

	assert := assert.New(t)

	e := event.AccountCreated{
		UserID: 12345,
		Email:  "hoge@email.com",
	}

	task, err := createTask(e)

	assert.NoError(err)
	assert.Equal("POST", task.Method)

	values, err := url.ParseQuery(string(task.Payload))
	assert.NoError(err)
	assert.Equal("account_created", values.Get("name"))
	payload, err := base64.StdEncoding.DecodeString(values.Get("payload"))
	assert.NoError(err)
	assert.JSONEq(`{"user_id": 12345, "email": "hoge@email.com"}`, string(payload))
}

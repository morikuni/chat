package user

import (
	"crypto/sha256"
	"io"
	"math/rand"
	"time"

	"github.com/morikuni/chat/chat/domain/model"
	"github.com/pkg/errors"
)

func newAuthInfo(email model.Email, password model.Password) AuthInfo {
	return AuthInfo{
		email,
		newHashedPassword(password),
	}
}

type AuthInfo struct {
	Email    model.Email
	Password HashedPassword
}

func (ai AuthInfo) Authenticate(email model.Email, password model.Password) error {
	if ai.Email != email {
		return errors.New("email does not match")
	}
	if !ai.Password.Equal(password) {
		return errors.New("password does not match")
	}
	return nil
}

func newHashedPassword(raw model.Password) HashedPassword {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(r, salt); err != nil {
		panic(errors.Wrap(err, "failed to read salt"))
	}
	return newHashedPasswordWithSalt(raw, salt)
}

func newHashedPasswordWithSalt(raw model.Password, salt []byte) HashedPassword {
	data := append([]byte(raw), salt...)
	hash := sha256.Sum256(data)
	return HashedPassword{
		hash[:],
		salt,
	}
}

type HashedPassword struct {
	Hash []byte
	Salt []byte
}

func (hp HashedPassword) Equal(raw model.Password) bool {
	other := newHashedPasswordWithSalt(raw, hp.Salt)
	return string(hp.Hash) == string(other.Hash) && string(hp.Salt) == string(other.Salt)
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

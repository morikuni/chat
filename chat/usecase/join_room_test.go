package usecase

// import (
// 	"context"
// 	"testing"
//
// 	"github.com/morikuni/chat/chat/domain/model/category"
// 	"github.com/morikuni/chat/chat/domain/model/room"
// 	"github.com/morikuni/chat/chat/domain/model/roommember"
// 	"github.com/morikuni/chat/chat/domain/model/user"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestJoinRoom(t *testing.T) {
// 	assert := assert.New(t)
// 	ctx := context.Background()
//
// 	userRepo := user.NewRepository()
// 	roomRepo := room.NewRepository()
// 	roomMemberRepo := roommember.NewRepository()
//
// 	u := user.New("name", "email", "password")
// 	assert.Nil(userRepo.Save(ctx, u))
//
// 	c := category.New("name")
// 	r := room.New(c, "name", "description")
// 	assert.Nil(roomRepo.Save(ctx, r))
//
// 	jr := NewJoinRoom(userRepo, roomRepo, roomMemberRepo)
//
// 	id, err := jr.Join(ctx, string(u.ID()), string(r.ID()))
// 	assert.Nil(err)
// 	assert.NotEmpty(id)
//
// 	rm, err := roomMemberRepo.Find(ctx, id)
//
// 	assert.Nil(err)
// 	assert.NotEmpty(rm)
// }

package infra

//
// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"testing"
//
// 	"github.com/lib/pq"
// 	"github.com/pkg/errors"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestAppend(t *testing.T) {
// 	assert := assert.New(t)
//
// 	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/chat?sslmode=disable")
// 	assert.NoError(err)
//
// 	repo := postgresRepo{}
// 	rcp := repositoryContextProvider{db}
// 	ctx, err := rcp.WithTx(context.Background())
// 	assert.NoError(err)
// 	ctx, err = rcp.WithTx(ctx)
// 	assert.NoError(err)
//
// 	obj, err := json.Marshal(map[string]interface{}{
// 		"name": "tarou",
// 	})
// 	assert.NoError(err)
//
// 	err = repo.Append(ctx, "user", "hogehoeg", obj)
// 	if err != nil {
// 		err := errors.Cause(err)
// 		if e, ok := err.(*pq.Error); ok {
// 			t.Log(e.Code.Name())
// 		}
// 	}
// 	assert.NoError(err)
//
// 	version, obj, err := repo.Find(ctx, "user", "hogehoeg")
// 	assert.NoError(err)
// 	assert.Equal(0, version)
// 	var i interface{}
// 	json.Unmarshal(obj, &i)
//
// 	obj, err = json.Marshal(map[string]interface{}{
// 		"name": "hoge",
// 		"age":  16,
// 	})
// 	err = repo.Save(ctx, "user", "hogehoeg", 0, obj)
// 	assert.NoError(err)
//
// 	version, obj, err = repo.Find(ctx, "user", "hogehoeg")
// 	assert.NoError(err)
// 	assert.Equal(1, version)
// 	json.Unmarshal(obj, &i)
//
// 	ctx.Commit()
// }

package main

import (
	"net/http"

	"github.com/morikuni/chat/chat/api"
	"github.com/morikuni/chat/chat/di"
	"github.com/morikuni/yacm"
)

func main() {
	env := yacm.NewEnv().
		AppendCatcherFuncs(api.Catcher).
		WithShutterFunc(api.Shutter)

	apis := []api.API{
		di.NewSignUp(),
	}

	for _, api := range apis {
		http.Handle(api.Path(), env.Serve(api))
	}

	http.ListenAndServe(":8081", nil)
}

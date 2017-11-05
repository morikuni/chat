package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morikuni/chat/src/di"
)

func init() {
	router := mux.NewRouter()

	api := di.InjectAPI()
	router.HandleFunc("/chats", api.GetChats).Methods("GET")
	router.HandleFunc("/chats", api.PostChats).Methods("POST")

	th := di.InjectTaskHandler()
	router.Handle("/internal/event", th)

	http.Handle("/", router)
}

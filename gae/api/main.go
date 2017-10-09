package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morikuni/chat/src/di"
)

func init() {
	api := di.InjectAPI()

	router := mux.NewRouter()

	router.HandleFunc("/chats", api.GetChats).Methods("GET")
	router.HandleFunc("/chats", api.PostChats).Methods("POST")

	http.Handle("/", router)
}

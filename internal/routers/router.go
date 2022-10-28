package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosa12978/ElonScrewedUp/internal/endpoints"
)

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("", endpoints.GetTweets).Methods("GET")
	router.HandleFunc("{uuid}", endpoints.GetTweetByUUID).Methods("GET")
	router.HandleFunc("", endpoints.CreateTweet).Methods("POST")
	router.HandleFunc("{uuid}", endpoints.UpdateTweet).Methods("PUT")
	router.HandleFunc("{uuid}", endpoints.DeleteTweet).Methods("DELETE")
	return router
}

package server

import (
	"net/http"

	"github.com/yosa12978/ElonScrewedUp/internal/routers"
)

func Run() {
	mux := routers.NewRouter()
	server := http.Server{
		Addr:    "0.0.0.0:8989",
		Handler: mux,
	}
	server.ListenAndServe()
}

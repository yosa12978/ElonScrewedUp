package server

import (
	"net/http"

	db "github.com/yosa12978/ElonScrewedUp/internal/gorm"
	"github.com/yosa12978/ElonScrewedUp/internal/routers"
)

func Run() {
	db.Connect("database.db")
	mux := routers.NewRouter()
	server := http.Server{
		Addr:    "0.0.0.0:8989",
		Handler: mux,
	}
	server.ListenAndServe()
}

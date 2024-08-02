package server

import (
	"discord/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func New() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: SetupRouter(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to init server")
	}
}

func SetupRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/get-items", controllers.GetById)

	return r
}

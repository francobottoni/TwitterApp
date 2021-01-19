package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/francobottoni/TwitterApp/middleware"
	"github.com/francobottoni/TwitterApp/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

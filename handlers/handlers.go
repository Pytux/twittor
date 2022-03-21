package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/pytux/twittor/middlewares"
	"github.com/pytux/twittor/routers"
)

func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/profile/edit", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Printf("Server running in PORT %s", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

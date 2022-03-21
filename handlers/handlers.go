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
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.CreateTweets))).Methods("POST")
	router.HandleFunc("/tweet/delete", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetTweets))).Methods("GET")

	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/relationship", middlewares.CheckDB(middlewares.ValidateJWT(routers.CreateRelationship))).Methods("GET")
	router.HandleFunc("/relationship/query", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetRelationship))).Methods("GET")
	router.HandleFunc("/relationship/delete", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteRelationship))).Methods("DELETE")

	router.HandleFunc("/userList", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetAllUsers))).Methods("GET")

	router.HandleFunc("/gettweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetFollowedTweets))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Printf("Server running in PORT %s", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

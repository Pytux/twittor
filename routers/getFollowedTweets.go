package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pytux/twittor/db"
)

func GetFollowedTweets(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page must be higher than 0.", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page must be higher than 0.", http.StatusBadRequest)
		return
	}

	response, error := db.GetFollowedTweets(UserID, page)
	if !error {
		http.Error(w, "Error getting tweets.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

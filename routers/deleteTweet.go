package routers

import (
	"net/http"

	"github.com/pytux/twittor/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You need send the parameter ID", 400)
		return
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "Can't delete tweet", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

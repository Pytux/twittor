package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pytux/twittor/db"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You need send the parameter ID", 400)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You need send the parameter page", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You need send the parameter page", 400)
		return
	}

	pag := int64(page)

	response, success := db.GetTweets(ID, pag)
	if !success {
		http.Error(w, "Error to get tweets", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

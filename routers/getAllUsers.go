package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pytux/twittor/db"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page must be higher than 0.", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := db.GetAllUsers(UserID, pag, search, typeUser)
	if !status {
		http.Error(w, "Error trying to get users. ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

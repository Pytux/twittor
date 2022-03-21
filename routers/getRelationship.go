package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func GetRelationship(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You need send the parameter ID.", http.StatusBadRequest)
		return
	}

	var relationship models.Relationship
	relationship.UserID = UserID
	relationship.UserRelationshipID = ID

	var response models.QueryRelationship

	status, err := db.GetRelationship(relationship)
	if err != nil || !status {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

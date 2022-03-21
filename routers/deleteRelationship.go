package routers

import (
	"net/http"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func DeleteRelationship(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You need send the parameter ID.", http.StatusBadRequest)
		return
	}

	var relationship models.Relationship
	relationship.UserID = UserID
	relationship.UserRelationshipID = ID

	status, err := db.DeleteRelationship(relationship)
	if err != nil {
		http.Error(w, "An error occurred. Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't delete relationship.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

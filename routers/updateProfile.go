package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid body "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateUser(user, UserID)

	if err != nil {
		http.Error(w, "An error occurred while updating user. Wait a minute and try again."+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't make the changes in the user.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

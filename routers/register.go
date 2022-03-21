package routers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error while trying decoding, "+err.Error(), 400)
		return
	}

	_, err = mail.ParseAddress(user.Email)

	if err != nil {
		http.Error(w, "The email field is not valid.", 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "The email is required.", 400)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Minimum password length is 8.", 400)
		return
	}

	_, exist, _ := db.CheckIfEmailIsInUse(user.Email)

	if exist {
		http.Error(w, "The email is used by other user.", 400)
		return
	}

	_, status, err := db.SaveUser(user)
	if err != nil {
		http.Error(w, "An error ocurred when trying to register the user, "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't save the user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

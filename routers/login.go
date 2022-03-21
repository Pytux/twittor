package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/jwt"
	"github.com/pytux/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "applicatopm/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Email or Password  invalid.", 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required.", 400)
		return
	}

	user, exist := db.Login(user.Email, user.Password)

	if !exist {
		http.Error(w, "The user not exist.", 400)
		return
	}

	jwtKey, err := jwt.GenerateToken(user)
	if err != nil {
		http.Error(w, "A an error occurred while token was generated.", 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

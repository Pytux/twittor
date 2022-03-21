package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/pytux/twittor/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You need send the parameter ID.", http.StatusBadRequest)
		return
	}

	perfil, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error, cant copy image.", http.StatusBadRequest)
	}
}

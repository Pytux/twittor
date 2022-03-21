package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")

	var ext = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/avatar/" + UserID + "." + ext

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading avatar. "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error, can't copy avatar. "+err.Error(), http.StatusBadRequest)
		return
	}

	var (
		user   models.User
		status bool
	)

	user.Avatar = UserID + "." + ext

	status, err2 := db.UpdateUser(user, UserID)
	if err2 != nil || !status {
		http.Error(w, "Error, can't save avatar. "+err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

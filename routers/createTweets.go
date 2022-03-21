package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

func CreateTweets(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Invalid body "+err.Error(), 400)
		return
	}

	register := models.Tweet{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error occurred. Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Can't save tweet.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

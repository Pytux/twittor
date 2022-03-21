package middlewares

import (
	"net/http"

	"github.com/pytux/twittor/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Cant connect with the database", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}

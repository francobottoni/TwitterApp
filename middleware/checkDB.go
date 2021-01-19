package middleware

import (
	"net/http"

	"github.com/francobottoni/TwitterApp/database"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConection() == false {
			http.Error(w, "Conection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/francobottoni/TwitterApp/database"
	"github.com/francobottoni/TwitterApp/models"
)

// Is a func that create a register in a bd
func Login(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in recibes values", 400)
		return
	}

	if t.Email == "" {
		http.Error(w, "Email is required ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "A password of at least 6 characters is required", 400)
	}

	_, found, _ := database.IfUserAlreadyExist(t.Email)

	if found == true {
		http.Error(w, "This user already exist in the App", 400)
		return
	}

	_, status, err, u := database.InsertUser(t)

	if err != nil {
		http.Error(w, "Error when insert a new user", 400)
		return
	}

	if status == false {
		http.Error(w, "Can't insert a User", 400)
		return
	}

	j, _ := json.Marshal(u)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

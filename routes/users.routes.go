package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Okrams/go-gorm-api/db"
	"github.com/Okrams/go-gorm-api/models"
	"github.com/gorilla/mux"
)

func GetUsersHanlder(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

}

func GetUserHanlder(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHanlder(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHanlder(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}

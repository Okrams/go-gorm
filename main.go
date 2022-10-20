package main

import (
	"net/http"

	"github.com/Okrams/go-gorm-api/db"
	"github.com/Okrams/go-gorm-api/models"
	"github.com/Okrams/go-gorm-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHanlder)
	usersRoute := router.PathPrefix("/users").Subrouter()

	usersRoute.HandleFunc("", routes.GetUsersHanlder).Methods("GET")
	usersRoute.HandleFunc("/{id}", routes.GetUserHanlder).Methods("GET")
	usersRoute.HandleFunc("", routes.PostUserHanlder).Methods("POST")
	usersRoute.HandleFunc("", routes.DeleteUserHanlder).Methods("DELETE")

	taskRoute := router.PathPrefix("/tasks").Subrouter()
	taskRoute.HandleFunc("", routes.GetTasksHandler).Methods("GET")
	taskRoute.HandleFunc("/{id}", routes.GetTaskHandler).Methods("GET")
	taskRoute.HandleFunc("", routes.CreateTaskHandler).Methods("POST")
	taskRoute.HandleFunc("/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3003", router)
}

package handler

import (
	"Go-sumon/structure"
	"net/http"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []structure.User
	GenericGetAllHandler(w, r, "user", &users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	UserCreateHandler(w, r, "user")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	GenericGetHandler(w, r, "user")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	GenericUpdateHandler(w, r, "user")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	GenericDeleteHandler(w, r, "user")
}

func FindUserHandler(w http.ResponseWriter, r *http.Request) {
	GenericFindHandler(w, r, "user")
}

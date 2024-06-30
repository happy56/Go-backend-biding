package handler

import (
    "net/http"
	"Go-sumon/structure"
)

func GetAllSPHandler(w http.ResponseWriter, r *http.Request) {
    var users []structure.Client
    GenericGetAllHandler(w, r, "serviceProvider", &users)
}

func CreateSPHandler(w http.ResponseWriter, r *http.Request) {
    SPCreateHandler(w, r, "user", "serviceProvider")
}

func GetSPHandler(w http.ResponseWriter, r *http.Request) {
    GenericGetHandler(w, r, "serviceProvider")
}

func UpdateSPHandler(w http.ResponseWriter, r *http.Request) {
    GenericUpdateHandler(w, r, "serviceProvider")
}

func DeleteSPHandler(w http.ResponseWriter, r *http.Request) {
    GenericDeleteHandler(w, r, "serviceProvider")
}

func FindSPHandler(w http.ResponseWriter, r *http.Request) {
    GenericFindHandler(w, r, "serviceProvider")
}
package handler

import (
    "net/http"
	"Go-sumon/structure"
)

func GetAllClientHandler(w http.ResponseWriter, r *http.Request) {
    var client []structure.Client
    GenericGetAllHandler(w, r, "client", &client)
}

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
    ClientCreateHandler(w, r, "user", "client")
}

func GetClientHandler(w http.ResponseWriter, r *http.Request) {
    GenericGetHandler(w, r, "client")
}

func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
    GenericUpdateHandler(w, r, "client")
}

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
    GenericDeleteHandler(w, r, "client")
}

func FindClientHandler(w http.ResponseWriter, r *http.Request) {
    GenericFindHandler(w, r, "client")
}
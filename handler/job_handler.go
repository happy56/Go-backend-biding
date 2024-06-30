package handler

import (
    "net/http"
	"Go-sumon/structure"
)

func GetAllJobHandler(w http.ResponseWriter, r *http.Request) {

    var job []structure.Review
    GenericGetAllHandler(w, r, "job", &job)
}

func CreateJobHandler(w http.ResponseWriter, r *http.Request) {
   
    var job structure.Review
    GenericCreateHandler(w, r, "job", &job)
}

func GetJobHandler(w http.ResponseWriter, r *http.Request) {
    GenericGetHandler(w, r, "job")
}

func UpdateJobHandler(w http.ResponseWriter, r *http.Request) {
    GenericUpdateHandler(w, r, "job")
}

func DeleteJobHandler(w http.ResponseWriter, r *http.Request) {
    GenericDeleteHandler(w, r, "job")
}

func FindJobHandler(w http.ResponseWriter, r *http.Request) {
    GenericFindHandler(w, r, "job")
}
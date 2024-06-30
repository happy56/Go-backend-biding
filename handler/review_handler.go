package handler

import (
    "net/http"
	"Go-sumon/structure"
)

func GetAllReviewHandler(w http.ResponseWriter, r *http.Request) {

    var reviews []structure.Review
    GenericGetAllHandler(w, r, "review", &reviews)
}

func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
   
    var review structure.Review
    GenericCreateHandler(w, r, "review", &review)
}

func GetReviewHandler(w http.ResponseWriter, r *http.Request) {
    GenericGetHandler(w, r, "review")
}

func UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
    GenericUpdateHandler(w, r, "review")
}

func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
    GenericDeleteHandler(w, r, "review")
}

func FindReviewHandler(w http.ResponseWriter, r *http.Request) {
    GenericFindHandler(w, r, "review")
}
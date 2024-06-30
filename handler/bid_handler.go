package handler

import (
    "net/http"
    "Go-sumon/structure"
)

func GetAllBidHandler(w http.ResponseWriter, r *http.Request) {
    var bids []structure.Bid
    GenericGetAllHandler(w, r, "bid", &bids)
}

func CreateBidHandler(w http.ResponseWriter, r *http.Request) {
    var bid structure.Bid
    GenericCreateHandler(w, r, "bid", &bid)
}

func GetBidHandler(w http.ResponseWriter, r *http.Request) {
    GenericGetHandler(w, r, "bid")
}

func UpdateBidHandler(w http.ResponseWriter, r *http.Request) {
    GenericUpdateHandler(w, r, "bid")
}

func DeleteBidHandler(w http.ResponseWriter, r *http.Request) {
    GenericDeleteHandler(w, r, "bid")
}

func FindBidHandler(w http.ResponseWriter, r *http.Request) {
    GenericFindHandler(w, r, "bid")
}

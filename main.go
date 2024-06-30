package main

import (
	"Go-sumon/fileuploader"
	"Go-sumon/handler"
	"log"
	"net/http"
)

func main() {
	// Register HTTP handlers for bid routes
	http.HandleFunc("/bid", enableCors(handler.GetAllBidHandler))
	http.HandleFunc("/bid/create", enableCors(handler.CreateBidHandler))
	http.HandleFunc("/bid/{_id}", enableCors(handler.GetBidHandler))
	http.HandleFunc("/bid/{_id}/update", enableCors(handler.UpdateBidHandler))
	http.HandleFunc("/bid/{_id}/delete", enableCors(handler.DeleteBidHandler))
	http.HandleFunc("/bid/{_id}/find", enableCors(handler.FindBidHandler))

	// Register HTTP handlers for review routes
	http.HandleFunc("/review", enableCors(handler.GetAllReviewHandler))
	http.HandleFunc("/review/create", enableCors(handler.CreateReviewHandler))
	http.HandleFunc("/review/{_id}", enableCors(handler.GetReviewHandler))
	http.HandleFunc("/review/{_id}/update", enableCors(handler.UpdateReviewHandler))
	http.HandleFunc("/review/{_id}/delete", enableCors(handler.DeleteReviewHandler))
	http.HandleFunc("/review/{_id}/find", enableCors(handler.FindReviewHandler))

	// Start the HTTP server
	go func() {
		log.Println("Server listening on port 5000...")
		if err := http.ListenAndServe(":5000", nil); err != nil {
			log.Fatalf("Error starting server on port 5000: %v", err)
		}
	}()

	// Setup file upload routes
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/upload", enableCors(fileuploader.UploadFile))
	log.Println("File upload server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting file upload server on port 8080: %v", err)
	}
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request initiated")
	defer log.Println("Request finished")
	select {
	case <-time.After(5 * time.Second):
		//Print at comand line
		log.Println("Request successful")
		//Print at browser
		w.Write([]byte("Request completed"))

	case <-ctx.Done():
		//Print at comand line
		log.Println("Request cancelled")
		//Print at browser
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
	}
}

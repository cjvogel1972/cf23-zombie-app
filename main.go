package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func readiness(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	if dt.Minute()%10 < 5 {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, CodeFest!")
}

func main() {
	http.HandleFunc("/probe/readiness", readiness)
	http.HandleFunc("/probe/liveness", liveness)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

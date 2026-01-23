package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}

	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello %s from %s\n", target, hostname)
}

func main() {
	fmt.Println("starting server on :8081")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

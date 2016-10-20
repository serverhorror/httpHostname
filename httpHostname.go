package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func root(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", hostname)
}

func main() {
	log.Print("Starting")
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		fmt.Println("`LISTEN_PORT' must be set")
		return
	}
	log.Printf("listenPort: %v", listenPort)

	http.HandleFunc("/", root)
	http.ListenAndServe(":"+listenPort, nil)
}

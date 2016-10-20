package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%v", r)
	fmt.Fprintf(w, "%s\n", hostname)
}

func ping() {
	for {
		log.Print(".")
		time.Sleep(time.Second * 3)
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	go ping()

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

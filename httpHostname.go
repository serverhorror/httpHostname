package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type hostName struct {
	Hostname string `json:"hostname"`
}

func (h *hostName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("\n%s\n", dump)

	n, err := os.Hostname()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Hostname = n

	b, err := json.Marshal(h)

	_, err = w.Write(b)
	if err != nil {
		log.Print(err)
		return
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Print("Starting")
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		fmt.Println("`LISTEN_PORT' must be set")
		return
	}
	log.Printf("listenPort: %v", listenPort)

	http.Handle("/", &hostName{})
	http.ListenAndServe(":"+listenPort, nil)
}

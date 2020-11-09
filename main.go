package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

// Message type...
type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to Simple API"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}

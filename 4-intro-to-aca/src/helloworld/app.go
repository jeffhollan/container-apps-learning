package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	log.Print("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Call URL/{name} to get a greeting
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

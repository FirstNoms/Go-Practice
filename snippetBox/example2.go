package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! welcome to Chinomso's E-wears")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)

	http.ListenAndServe(":8080", mux)

}

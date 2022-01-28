package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*This is the First Go Web Application example.
when writing a Go web application:
	1-> Create a Handler function that receives the HTTP request from the client
	2-> Then Serves static assets like javascript, css etc by pointing it to a url path.
	3-> Then Accepts connections from the internet by listening/connecting to the port
*/

//Handler Function
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Nomso! Welcome to my website🥰")
}

func main() {
	/*	http.HandleFunc("/", handler)

		fs:= http.FileServer(http.Dir("static/"))
		http.Handle("/static", http.StripPrefix("/static", fs))

		http.ListenAndServe(":80", nil)
	*/

	r := mux.NewRouter()

	r.HandleFunc("books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(writer, "youve reequested the book: %s on page %s\n", title, page)
	})
	http.ListenAndServe(":80", r)

}

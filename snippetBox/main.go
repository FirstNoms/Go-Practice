package main

/* 	When building a Go Web application, you need:
-> A Handler: serves as the controller. Executes your Application logic, writes the HTTP response to the User
-> Router(serve mux): This maps your application URL and the corresponding handlers.
-> Web Server: established by go. Listens for incoming HTTP request.

*/
import (
	"log"
	"net/http"
)

//Define a Handler function called "home", writes a byte slice and contains "hello from Snippet box" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippet box 🥰😘"))
}

//Second Handler:
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific Snippet..."))
}

//Third Handler:
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//Header().Set() method adds an 'Allow: POST' header to the response header map.
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(405)
		//This method below is used to send the 405 status code anthe string message as the response body.
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new Snippet....."))
}

func main() {
	//Use the http.NewServeMux() function to initialize a new server mux, then register the home function as the handler for the "/" URL pattern. This is the ROUTER
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//use the http.ListenAndServe()function to start a new WEB SERVER. we pass two parameters: TCP network address to listen on(":port number") and the router we just created. if http.ListenAndServe()returns an error, we use the log.fatal()function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

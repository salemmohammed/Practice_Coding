package main

import (
	user "RestAPI/api"
	"log"
	"net/http"
)

func main() {
	//register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})      // an object to implement a handler interface
	log.Fatal(http.ListenAndServe(":8080", nil)) // if i cannt lesson log.fatal.

	/*
		ListenAndServe starts an HTTP server with a given address and handler. The handler is usually
		nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

		http.Handle("/foo", fooHandler)

		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})

		log.Fatal(http.ListenAndServe(":8080", nil))

	*/

	/*
		More control over the server's behavior is available by creating a custom Server:
			s := &http.Server{
				Addr:           ":8080",
				Handler:        myHandler,
				ReadTimeout:    10 * time.Second,
				WriteTimeout:   10 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}
			log.Fatal(s.ListenAndServe())
	*/

}

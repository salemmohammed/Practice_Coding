package main

import (
	"fmt"
	"log"
	"net/http"
)

var userCounter int

type reportCounter struct {
	counter int
}

func main() {
	http.HandleFunc("/users", userHandleFunc)
	var rc reportCounter
	http.Handle("/reports", &rc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (rc *reportCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /reports")
	rc.counter++
	s := fmt.Sprintf("/reports API call count: %v", rc.counter)
	fmt.Fprintf(w, s)
}
func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the page")
	fmt.Fprintf(w, "hi, thanks %v", r.Method)
}

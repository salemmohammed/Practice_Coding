package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func GetRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "GET" {
		data := Response{Message: "Hello World From - GET"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	} else {
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func PostRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "POST" {
		data := Response{Message: "Hello World From - POST"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	} else {
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func main() {
	http.HandleFunc("/get", GetRequest)
	http.HandleFunc("/post", PostRequest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

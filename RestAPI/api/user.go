package user

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UserAPI struct{}

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db = []*User{}
var nextUserID uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("Got call for %v using HTTP METHOD %v \n", r.URL, r.Method)
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodDelete:
		doDelete(w, r)
	case http.MethodPut:
		doPut(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to %v \n", r.Method, r.URL)
		log.Printf("Unsupported methond '%v' to %v\n", r.Method, r.URL)
	}
}

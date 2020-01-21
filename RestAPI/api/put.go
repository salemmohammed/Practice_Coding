package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func doPut(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Got HTTP method '%v' to %v \n", r.Method, r.URL)

	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to update user %v", id)

	lock.Lock()
	//nextUserID++

	var tmp = []*User{}
	for _, u := range db {
		if id == u.ID {
			continue
		}
		tmp = append(tmp, u)
	}
	db = tmp
	//aUser.ID = nextUserID
	db = append(db, aUser)
	lock.Unlock()

	respUser := User{ID: aUser.ID, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}

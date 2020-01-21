package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func doDelete(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Got HTTP method '%v' to %v \n", r.Method, r.URL)
	// w.Header().Set("Connect-Type", "application/json")
	// jd := json.NewDecoder(r.Body)
	// aUser := &User{}
	// err := jd.Decode(aUser)
	// if nil != err {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to delete user %v", id)

	lock.Lock()
	var tmp = []*User{}
	for _, u := range db {
		if id == u.ID {
			continue
		}
		tmp = append(tmp, u)
	}
	db = tmp
	lock.Unlock()
}

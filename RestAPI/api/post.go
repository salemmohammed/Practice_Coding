package user

import (
	"encoding/json"
	"net/http"
)

func doPost(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Got HTTP method '%v' to %v \n", r.Method, r.URL)
	w.Header().Set("Connect-Type", "application/json")
	jd := json.NewDecoder(r.Body)
	aUser := &User{}
	err := jd.Decode(aUser)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lock.Lock()
	nextUserID++
	aUser.ID = nextUserID
	db = append(db, aUser)
	lock.Unlock()

	respUser := User{ID: aUser.ID, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}

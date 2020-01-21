package user

import (
	"encoding/json"
	"net/http"
)

func doGet(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Got HTTP method '%v' to %v \n", r.Method, r.URL)
	w.Header().Set("Connect-Type", "application/json")
	je := json.NewEncoder(w)
	je.Encode(db)
}

package main
import (
	"log"
	"io"
	"bytes"
	"encoding/json"
	"os"
	"fmt"
)
// Simple Structs
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}
type UserDb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func createJsonFile(){
	users := []User{
		{Username:"salem alqhatani", Password:"1000",Email:"salem055@gmail.com"},
		{Username:"salem alqhatani", Password:"1000",Email:"salem055@gmail.com"},
	}

	db := UserDb{Users:users,Type:"simple"}
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)

	enc.Encode(db)
	filename, err := os.Create("user.db.json")
	if err != nil{
		log.Fatalln(err)
	}
	defer filename.Close()
	io.Copy(filename, buf)
}
func main(){
	// Two options, create or open file name
	file, err := os.Open("user.db.json")
	if err != nil{
		log.Fatalln(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	db := UserDb{} 
	dec.Decode(&db)
	fmt.Println(db)

	//createJsonFile()
}
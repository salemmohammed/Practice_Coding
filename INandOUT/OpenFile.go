package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var f, err = os.Open("data.txt") // open return a pointer or an error
	if err != nil {
		log.Fatalln(err) // kill my porgram if there is an errors
	}
	defer f.Close()       // close the file
	io.Copy(os.Stdout, f) // standard out
}

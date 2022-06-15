package main

import (
	"family/father"
	"family/father/son"
	"fmt"
)

func main() {
	f := new(father.Father)
	fmt.Println(f.Data("Mr. Jeremy Maclin"))

	c := new(son.Son)
	fmt.Println(c.Data("Riley Maclin"))
}

// an implementation on how a struct called in custom package
// https://www.golangprograms.com/golang-package-examples/how-to-import-structs-from-another-package-in-go.html

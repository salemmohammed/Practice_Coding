package main

import (
	"fmt"
	"salem/consistent"
)

func main() {
	ch := consistent.NewConsistentHashing()
	ch.Add("server1")
	ch.Add("server2")
	ch.Add("server3")
	
	//[server1 server2 server3]

	fmt.Println(ch.ListNodes())
	targetObj := []string{"client1", "client2", "client3", "client4", "client5", "client6"}
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}
//client: client1 in server: server1
//client: client2 in server: server3
//client: client3 in server: server1
//client: client4 in server: server3
//client: client5 in server: server1
//client: client6 in server: server3

	ch.Add("server4")
	ch.Add("server5")
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}
}

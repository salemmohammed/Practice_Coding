package main

import (
	"log"
	"struct_Examples/car"
)

func main() {
	car2 := car.Newcar()
	//car2 := car.Car{Brand: "Toyota", Year: 2023}
	car3, err := car2.Display()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(car2)
	log.Print(car3)
}

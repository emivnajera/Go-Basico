package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 200}
	fmt.Println(myCar)

	//Otra Manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}

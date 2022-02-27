package main

import (
	"emivnajera/go-basico/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	mypackage.PrintMessage()
}

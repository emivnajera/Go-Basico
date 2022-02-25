package main

import (
	"emivnajera/go-basico/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)
}

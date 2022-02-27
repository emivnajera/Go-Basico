package mypackage

import "fmt"

//CarPublic

type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage() {
	fmt.Println("Hola")
}

package main

import "fmt"

func main() {
	//Println
	hellomessage := "Hello"
	Worldmessage := "World"
	fmt.Println(hellomessage, Worldmessage)

	//Printf
	nombre := "Emiliano"
	edad := 21
	fmt.Printf("%s tiene %d anios\n", nombre, edad)
	fmt.Printf("%v tiene %v anios\n", nombre, edad)

	//Sprintf
	message := fmt.Sprintf("%s tiene %d anios\n", nombre, edad)
	fmt.Println(message)

	//Tipo Datos
	fmt.Printf("hellomessage: %T", hellomessage)
}

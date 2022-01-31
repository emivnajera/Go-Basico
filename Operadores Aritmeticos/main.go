package main

import "fmt"

func main() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multipliacion", result)

	//Division
	result = y / x
	fmt.Println("Division: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremento
	x++
	fmt.Println("Incremento: ", x)

	//Decremento
	x--
	fmt.Println("Decremento: ", x)
}

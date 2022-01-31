package main

import "fmt"

func main() {
	// Declaracion de Constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	//Declaracion de Enteros
	base := 12
	var altura int = 14
	var area int

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Hola Mundo")
	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(base, altura, area)
	fmt.Println(a, b, c, d)

	//Area de un Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)
}

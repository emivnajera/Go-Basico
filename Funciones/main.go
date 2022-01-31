package main

import "fmt"

func normalFunction() {
	fmt.Println("Hola Mundo")
}

func mensaje(message string) {
	fmt.Println(message)
}

func suma(b int, c int) int {
	resultado := b + c
	return resultado
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction()
	mensaje("Hello World")
	fmt.Println(suma(5, 10))
	value1, value2 := dobleReturn(4)
	fmt.Println(value1, value2)
}

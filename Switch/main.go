package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	//Switch sin condicion

	value := 200
	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("Sin condicion")
	}
}

package main

import "fmt"

func main() {
	silce := []string{"Hola", "que", "hace"}

	for i, valor := range silce {
		fmt.Println(i, valor)
	}
}

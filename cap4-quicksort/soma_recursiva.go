package main

import "fmt"

func Soma(lista []int) int {

	if len(lista) == 0 { // caso base
		return 0
	}

	return lista[0] + Soma(lista[1:])
}

func main() {

	numeros := []int{2, 4, 6}

	resultado := Soma(numeros)

	fmt.Println("Soma:", resultado)
}

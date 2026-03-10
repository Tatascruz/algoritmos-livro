package main

import "fmt"

func Maior(lista []int) int {

	if len(lista) == 1 {
		return lista[0]
	}

	maiorDoResto := Maior(lista[1:])

	if lista[0] > maiorDoResto {
		return lista[0]
	}

	return maiorDoResto
}

func main() {

	numeros := []int{4, 6, 2, 9, 3}

	fmt.Println("Maior:", Maior(numeros))
}

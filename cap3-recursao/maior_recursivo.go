package main

import "fmt"

func MaiorRecursivo(lista []int) int {
	if len(lista) == 1 {
		return lista[0]
	}

	maiorDoResto := MaiorRecursivo(lista[1:])

	if lista[0] > maiorDoResto {
		return lista[0]
	}

	return maiorDoResto
}

func main() {
	fmt.Println(MaiorRecursivo([]int{2, 4, 6, 1}))
}

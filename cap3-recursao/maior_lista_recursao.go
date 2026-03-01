package main

import "fmt"

// MaiorRecursivo retorna o maior valor de uma lista usando recursão.
// Requisito: a lista não pode estar vazia.
func MaiorRecursivo(lista []int) int {
	if len(lista) == 1 {
		return lista[0]
	}

	maiordoResto := MaiorRecursivo(lista[1:])

	if lista[0] > maiordoResto {
		return lista[0]
	}

	return maiordoResto
}

func main() {
	fmt.Println(MaiorRecursivo([]int{3, 7, 2, 9, 5}))
	fmt.Println(MaiorRecursivo([]int{8, 4, 4}))
	fmt.Println(MaiorRecursivo([]int{5, 1, 7}))
}

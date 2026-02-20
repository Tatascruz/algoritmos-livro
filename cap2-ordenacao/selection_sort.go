package main

import "fmt"

// SelectionSort ordena a lista usando o algoritmo de ordenação por seleção (Selection Sort)
func SelectionSort(lista []int) []int {
	n := len(lista)

	for i := 0; i < n-1; i++ {
		menorIndice := i

		for j := i + 1; j < n; j++ {
			if lista[j] < lista[menorIndice] {
				menorIndice = j
			}
		}

		// Troca o menor encontrado para a posição i
		lista[i], lista[menorIndice] = lista[menorIndice], lista[i]
	}

	return lista
}

func main() {
	lista := []int{29, 10, 14, 37, 13}

	fmt.Println("Lista original:", lista)
	fmt.Println("Lista ordenada:", SelectionSort(lista))
}

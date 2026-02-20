package main

import "fmt"

// BuscaBinaria retorna o índice do item se achar, se não -1
func BuscaBinaria(lista []int, alvo int) int {
	esq := 0
	dir := len(lista) - 1

	for esq <= dir {
		meio := (esq + dir) / 2

		if lista[meio] == alvo {
			return meio
		}

		if lista[meio] < alvo {
			esq = meio + 1
		} else {
			dir = meio - 1
		}

	}

	return -1
}

func main() {
	lista := []int{3, 7, 12, 19, 25, 31, 44}
	alvo := 19

	fmt.Println("Busca bínaria:", BuscaBinaria(lista, alvo))
}

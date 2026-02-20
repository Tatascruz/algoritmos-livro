package main

import "fmt"

func BuscaLinear(lista []int, alvo int) (int, int) {
	tentativas := 0

	for i, v := range lista {
		tentativas++
		if v == alvo {
			return i, tentativas
		}
	}

	return -1, tentativas
}

func BuscaBinaria(lista []int, alvo int) (int, int) {
	esq := 0
	dir := len(lista) - 1
	tentativas := 0

	for esq <= dir {
		tentativas++
		meio := (esq + dir) / 2

		if lista[meio] == alvo {
			return meio, tentativas
		}

		if lista[meio] < alvo {
			esq = meio + 1
		} else {
			dir = meio - 1
		}
	}
	return -1, tentativas
}

func main() {
	lista := []int{10, 20, 30, 40, 50, 60, 70}
	alvo := 60

	indiceLinear, tentativasLinear := BuscaLinear(lista, alvo)
	indiceBinaria, tentativasBinaria := BuscaBinaria(lista, alvo)

	fmt.Println("Busca Linear:")
	fmt.Println("Índice:", indiceLinear)
	fmt.Println("Tentaivas:", tentativasLinear)

	fmt.Println("\nBusca Binária:")
	fmt.Println("Índice:", indiceBinaria)
	fmt.Println("Tentativas:", tentativasBinaria)
}

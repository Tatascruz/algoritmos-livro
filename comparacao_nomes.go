package main

import "fmt"

func BuscaLinearNomes(lista []string, alvo string) (int, int) {
	tentativas := 0

	for i, v := range lista {
		tentativas++
		if v == alvo {
			return i, tentativas
		}
	}

	return -1, tentativas
}

func BuscaBinariaNomes(lista []string, alvo string) (int, int) {
	esq := 0
	dir := len(lista) - 1
	tentativas := 0

	for esq <= dir {
		meio := (esq + dir) / 2
		tentativas++

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
	nomes := []string{
		"Anna",
		"Bruno",
		"Carlos",
		"Daniela",
		"Eduardo",
		"Fernada",
		"Gabriel",
	}

	alvo := "Anna"

	indiceLinear, tentativasLinear := BuscaLinearNomes(nomes, alvo)
	indiceBinaria, tentativasBinaria := BuscaBinariaNomes(nomes, alvo)

	fmt.Println("Busca Linear (Nomes):")
	fmt.Println("Índice:", indiceLinear)
	fmt.Println("Tentativas:", tentativasLinear)

	fmt.Println("\nBusca Binária (Nomes):")
	fmt.Println("Índice:", indiceBinaria)
	fmt.Println("Tentativas:", tentativasBinaria)
}

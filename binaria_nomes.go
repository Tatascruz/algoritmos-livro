package main

import "fmt"

func BuscaBinariaNomes(lista []string, alvo string) int {
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
	nomes := []string{
		"Anna",
		"Bruno",
		"Carlos",
		"Daniela",
		"Eduardo",
		"Fernanda",
		"Gabriel",
	}

	alvo := "Eduardo"

	resultado := BuscaBinariaNomes(nomes, alvo)

	if resultado != -1 {
		fmt.Println("Nome encontrado no índice:", resultado)
	} else {
		fmt.Println("Nome não encontrado")
	}
}

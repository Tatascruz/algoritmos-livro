package main

import "fmt"

func BuscaBinariaRecursiva(lista []int, alvo int) bool {
	if len(lista) == 0 {
		return false
	}

	meio := len(lista) / 2

	if lista[meio] == alvo {
		return true
	}

	if alvo < lista[meio] {
		return BuscaBinariaRecursiva(lista[:meio], alvo)
	}

	return BuscaBinariaRecursiva(lista[meio+1:], alvo)
}

func main() {
	numeros := []int{1, 3, 5, 7, 9}

	fmt.Println(BuscaBinariaRecursiva(numeros, 7))

	fmt.Println(BuscaBinariaRecursiva(numeros, 2))
}

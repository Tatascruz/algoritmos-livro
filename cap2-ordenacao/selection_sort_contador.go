package main

import "fmt"

type Resultado struct {
	Ordenado    []int
	Comparacoes int
	Trocas      int
}

func SelectionSortComContador(lista []int) Resultado {
	n := len(lista)
	comparacoes := 0
	trocas := 0

	for i := 0; i < n-1; i++ {
		menorIndice := i

		for j := i + 1; j < n; j++ {
			comparacoes++
			if lista[j] < lista[menorIndice] {
				menorIndice = j
			}
		}

		if menorIndice != i {
			lista[i], lista[menorIndice] = lista[menorIndice], lista[i]
			trocas++
		}
	}

	return Resultado{
		Ordenado:    lista,
		Comparacoes: comparacoes,
		Trocas:      trocas,
	}
}

func main() {
	lista := []int{29, 10, 14, 37, 13, 16, 45, 63, 55, 7}

	fmt.Println("Lista original:", lista)

	res := SelectionSortComContador(lista)

	fmt.Println("Lista ordenada :", res.Ordenado)
	fmt.Println("Comparacoes :", res.Comparacoes)
	fmt.Println("Trocas :", res.Trocas)
}

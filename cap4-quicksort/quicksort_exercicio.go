package main

import "fmt"

func quicksort(lista []int) []int {
	if len(lista) < 2 {
		return lista
	}

	pivo := lista[0]

	var menores []int
	var maiores []int

	for _, valor := range lista[1:] {

		if valor <= pivo {
			menores = append(menores, valor)
		} else {
			maiores = append(maiores, valor)
		}
	}

	menoresOrdenados := quicksort(menores)
	maioresOrdenados := quicksort(maiores)

	return append(append(menoresOrdenados, pivo), maioresOrdenados...)
}

func main() {

	lista := []int{8, 3, 5, 2, 9}

	resultado := quicksort(lista)

	fmt.Println("Lista Ordenada:", resultado)
}

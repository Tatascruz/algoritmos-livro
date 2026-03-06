package main

import "fmt"

func quicksort(lista []int) []int {

	// caso base
	if len(lista) < 2 {
		return lista
	}

	// escolhe pivô (aqui vou usar o primeiro para simplificar)
	pivo := lista[0]

	var menores []int
	var maiores []int

	// separa os elementos
	for _, valor := range lista[1:] {
		if valor <= pivo {
			menores = append(menores, valor)
		} else {
			maiores = append(maiores, valor)
		}
	}

	// recursão
	menoresOrdenados :=
		quicksort(menores)
	maioresOrdenados :=
		quicksort(maiores)

	// junta tudo
	return append(append(menoresOrdenados, pivo), maioresOrdenados...)
}

func main() {
	lista := []int{50, 3, 40, 12, 18, 9, 1}

	resultado := quicksort(lista)

	fmt.Println(resultado)
}

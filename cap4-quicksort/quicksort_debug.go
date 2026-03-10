package main

import "fmt"

func quicksort(lista []int) []int {

	fmt.Println("Lista recebida:", lista)

	if len(lista) < 2 {
		fmt.Println("Caso base:", lista)
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

	fmt.Println("Pivô:", pivo)
	fmt.Println("Menores:", menores)
	fmt.Println("Maiores:", maiores)
	fmt.Println("----", pivo)

	menoresOrdenados := quicksort(menores)
	maioresOrdenados := quicksort(maiores)

	resultado := append(append(menoresOrdenados, pivo), maioresOrdenados...)

	fmt.Println("Resultado parcial:", resultado)

	return resultado
}

func main() {

	lista := []int{8, 3, 5, 2, 9}

	fmt.Println("Lista original:", lista)
	fmt.Println("---------------")

	resultado := quicksort(lista)

	fmt.Println("--------------")
	fmt.Println("Lista ordenada:", resultado)
}

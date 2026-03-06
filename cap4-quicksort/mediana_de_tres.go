package main

import "fmt"

func medianaDeTres(lista []int) int {
	primeiro := lista[0]
	meio := lista[len(lista)/2]
	ultimo := lista[len(lista)-1]

	valores := []int{primeiro, meio, ultimo}

	//ordenando os três
	for i := 0; i < len(valores); i++ {
		for j := i + 1; j < len(valores); j++ {
			if valores[i] > valores[j] {
				valores[i], valores[j] = valores[j], valores[i]
			}
		}
	}

	return valores[1] // valor do meio pivô

}

func main() {
	lista := []int{30, 2, 18, 7, 25, 11, 4}

	pivo := medianaDeTres(lista)

	fmt.Println("Pivô escolhido:", pivo)
}

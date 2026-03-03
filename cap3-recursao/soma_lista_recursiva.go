package main

import "fmt"

func Soma(lista []int) int {
	if len(lista) == 0 {
		return 0
	}

	return lista[0] + Soma(lista[1:])
}

func main() {
	fmt.Println(Soma([]int{2, 4, 6, 8}))
}

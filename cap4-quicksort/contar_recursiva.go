package main

import "fmt"

func Contar(lista []int) int {
	if len(lista) == 0 {
		return 0
	}

	return 1 + Contar(lista[1:])
}

func main() {
	fmt.Println(Contar([]int{2, 4, 6}))

	fmt.Println(Contar([]int{10}))

	fmt.Println(Contar([]int{}))
}

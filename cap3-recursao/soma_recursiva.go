package main

import "fmt"

func SomaRecursiva(lista []int) int {
	if len(lista) == 0 {
		return 0
	}
	return lista[0] +
		SomaRecursiva(lista[1:])
}

func main() {
	fmt.Println(SomaRecursiva([]int{2, 4, 6}))

	fmt.Println(SomaRecursiva([]int{5}))

	fmt.Println(SomaRecursiva([]int{}))
}

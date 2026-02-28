package main

import "fmt"

func ContarRecursivo(lista []int) int {
	if len(lista) == 0 {
		return 0
	}
	return 1 + ContarRecursivo(lista[1:])
}

func main() {
	fmt.Println(ContarRecursivo([]int{2, 4, 6}))

	fmt.Println(ContarRecursivo([]int{5}))

	fmt.Println(ContarRecursivo([]int{}))
}

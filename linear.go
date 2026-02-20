package main

import "fmt"

// BuscaLinear retorna o índice do item se achar, senão -1
func BuscaLinear(lista []int, alvo int) int {
	for i, v := range lista {
		if v == alvo {
			return i
		}
	}
	return -1
}

func main() {
	lista := []int{3, 7, 12, 19, 25, 31, 44}
	alvo := 19

	fmt.Println("BuscaLinear:", BuscaLinear(lista, alvo))
}

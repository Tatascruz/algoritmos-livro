package main

import "fmt"

func ConategemComPilha(n int) {
	if n == 0 {
		fmt.Println("Chegou no caso base (0)")
		return
	}

	fmt.Println("Descendo:", n)

	ConategemComPilha(n - 1)

	fmt.Println("Subindo:", n)
}

func main() {
	fmt.Println("Início da execução")
	ConategemComPilha(3)
	fmt.Println("Fim da execução")
}

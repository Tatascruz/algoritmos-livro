package main

import "fmt"

func verificarEleitor(eleitor string, votaram map[string]bool) {
	if votaram[eleitor] {
		fmt.Println("Mande embora! Eleitor já votou")
	} else {
		votaram[eleitor] = true
		fmt.Println("Pode votar")
	}
}

func main() {
	votaram := make(map[string]bool)

	verificarEleitor("Tata", votaram)
	verificarEleitor("Anna", votaram)
	verificarEleitor("Tata", votaram)
	verificarEleitor("Tete", votaram)
	verificarEleitor("Anna", votaram)
}

package main

import "fmt"

type SistemaVotacao struct {
	votaram map[string]bool
}

func NovoSistemaVotacao() *SistemaVotacao {
	return &SistemaVotacao{
		votaram: make(map[string]bool),
	}
}

func (s *SistemaVotacao) verificarEleitor(nome string) string {
	if s.votaram[nome] {
		return "Mande embora! Eleitor ja votou."
	}

	s.votaram[nome] = true
	return "Pode votar!"
}

func main() {
	sistema := NovoSistemaVotacao()

	eleitores := []string{"Tata", "Anna", "Tata", "Tete", "Anna"}

	for _, eleitor := range eleitores {
		fmt.Printf("%s: %s\n", eleitor, sistema.verificarEleitor(eleitor))
	}
}

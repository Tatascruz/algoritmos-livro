package main

import (
	"fmt"
	"strings"
)

type SistemaVotacao struct {
	votaram map[string]bool
}

func NovoSistemaVotacao() *SistemaVotacao {
	return &SistemaVotacao{
		votaram: make(map[string]bool),
	}
}

func normalizerNome(nome string) string {
	return strings.ToLower(strings.TrimSpace(nome))
}

func (s *SistemaVotacao) VerificarEleitor(nome string) (bool, string) {
	nomeNormalizado := normalizerNome(nome)

	if nomeNormalizado == "" {
		return false, "Nome inválido."
	}

	if s.votaram[nomeNormalizado] {
		return false, "Manda embora! Eleitor ja votou."
	}

	s.votaram[nomeNormalizado] = true
	return true, "Pode votar!"
}

func main() {
	sistema := NovoSistemaVotacao()

	eleitores := []string{"Tata", "Anna", "tata", " ", "Tete", "anna"}

	for _, eleitor := range eleitores {
		podeVotar, mensagem := sistema.VerificarEleitor(eleitor)
		fmt.Printf("Eleitor: %q | Pode votar: %t | Mensagem: %s\n", eleitor, podeVotar, mensagem)
	}
}

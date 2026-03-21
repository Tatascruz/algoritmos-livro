package main

import (
	"fmt"
	"strings"
)

type Usuario struct {
	Nome  string
	Email string
	Idade int
}

type SistemaUsuarios struct {
	usuarios map[string]Usuario
}

func NovoSistemaUsuarios() *SistemaUsuarios {
	return &SistemaUsuarios{
		usuarios: make(map[string]Usuario),
	}
}

func normalizarChave(nome string) string {
	return strings.ToLower(strings.TrimSpace(nome))
}

func (s *SistemaUsuarios) Cadastrar(usuario Usuario) string {
	chave := normalizarChave(usuario.Nome)

	if chave == "" {
		return "Nome inválido."
	}

	if _, existe := s.usuarios[chave]; existe {
		return "Usuário ja cadastrado."
	}

	s.usuarios[chave] = usuario
	return "Usuário cadastrado com sucesso."
}

func (s *SistemaUsuarios) Buscar(nome string) (Usuario, bool) {
	chave := normalizarChave(nome)

	usuario, existe := s.usuarios[chave]
	return usuario, existe
}

func (s *SistemaUsuarios) Atualizar(nome string, novoEmail string, novaIdade int) string {
	chave := normalizarChave(nome)

	usuario, existe := s.usuarios[chave]
	if !existe {
		return "Usuário não encontrado."
	}

	usuario.Email = novoEmail
	usuario.Idade = novaIdade

	s.usuarios[chave] = usuario
	return "Usuário atualizado com sucesso."
}

func (s *SistemaUsuarios) Remover(nome string) string {
	chave := normalizarChave(nome)

	if _, existe := s.usuarios[chave]; !existe {
		return "Usuário não encontrado."
	}

	delete(s.usuarios, chave)
	return "Usuário removido com sucesso."
}

func (s *SistemaUsuarios) Listar() {
	if len(s.usuarios) == 0 {
		fmt.Println("Nenhum usuário cadastrado.")
		return
	}

	fmt.Println("Usuários cadastrados:")
	for _, usuario := range s.usuarios {
		fmt.Printf("- Nome: %s | Email: %s | Idade: %d\n", usuario.Nome, usuario.Email, usuario.Idade)
	}
}

func main() {
	sistema := NovoSistemaUsuarios()

	fmt.Println(sistema.Cadastrar(Usuario{
		Nome:  "Tata",
		Email: "tata@email.com",
		Idade: 40,
	}))

	fmt.Println(sistema.Cadastrar(Usuario{
		Nome:  "Anna",
		Email: "anna@email.com",
		Idade: 18,
	}))

	fmt.Println(sistema.Cadastrar(Usuario{
		Nome:  "Tata",
		Email: "outro@email.com",
		Idade: 31,
	}))

	fmt.Println()

	sistema.Listar()

	fmt.Println()

	usuario, encontrado := sistema.Buscar("anna")
	if encontrado {
		fmt.Printf("Usuário encontrado: %+v\n", usuario)
	} else {
		fmt.Println("Usuário não encontrado.")
	}

	fmt.Println()

	fmt.Println(sistema.Atualizar("Tata", "novo_tata@gmail", 31))
	fmt.Println(sistema.Remover("Anna"))

	fmt.Println()

	sistema.Listar()
}

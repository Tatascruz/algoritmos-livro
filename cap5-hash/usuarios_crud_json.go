package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Usuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

type SistemaUsuarios struct {
	usuarios map[string]Usuario
	arquivo  string
}

func NovoSistemaUsuarios(nomeArquivo string) *SistemaUsuarios {
	sistema := &SistemaUsuarios{
		usuarios: make(map[string]Usuario),
		arquivo:  nomeArquivo,
	}

	sistema.carregar()

	return sistema
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
	s.salvar()

	return "Usuário cadastrado com sucesso"
}

func (s *SistemaUsuarios) buscar(nome string) (Usuario, bool) {
	chave := normalizarChave(nome)

	usuario, existe := s.usuarios[chave]
	return usuario, existe
}

func (s *SistemaUsuarios) Atualizar(nome string, novoEmail string, novaIdade int) string {
	chave := normalizarChave(nome)

	usuario, existe := s.usuarios[chave]
	if !existe {
		return "Usuário não enocntrado."
	}

	usuario.Email = novoEmail
	usuario.Idade = novaIdade

	s.usuarios[chave] = usuario
	s.salvar()

	return "Usuário atualizado com sucesso."
}

func (s *SistemaUsuarios) Remover(nome string) string {
	chave := normalizarChave(nome)

	if _, existe := s.usuarios[chave]; !existe {
		return "Usuário não encontrado."
	}

	delete(s.usuarios, chave)
	s.salvar()

	return "Usuário removido com sucesso."
}

func (s *SistemaUsuarios) Listar() {
	if len(s.usuarios) == 0 {
		fmt.Println("Nenhum usuário cadastrado.")
		return
	}

	fmt.Println("Usuários cadastrado:")
	for _, usuario := range s.usuarios {
		fmt.Printf("- Nome: %s | Email: %s | Idade: %d\n", usuario.Nome, usuario.Email, usuario.Idade)
	}
}

func (s *SistemaUsuarios) salvar() {
	dados, err := json.MarshalIndent(s.usuarios, "", " ")
	if err != nil {
		fmt.Println("Erro ao gerar json:", err)
		return
	}

	err = os.WriteFile(s.arquivo, dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
	}
}

func (s *SistemaUsuarios) carregar() {
	dados, err := os.ReadFile(s.arquivo)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}

	err = json.Unmarshal(dados, &s.usuarios)
	if err != nil {
		fmt.Println("Erro ao carregar JSON:", err)
	}
}

func main() {
	sistema := NovoSistemaUsuarios("usuarios.json")

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

	fmt.Println()

	sistema.Listar()

	fmt.Println()

	fmt.Println(sistema.Atualizar("Tata", "novo_tata@gmail", 41))
	fmt.Println(sistema.Remover("Anna"))

	fmt.Println()

	sistema.Listar()
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		return "Usuário cadastrado com sucesso."
	}

	s.usuarios[chave] = usuario
	s.salvar()

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
		fmt.Println("Nenhum usuario cadastrado.")
		return
	}

	fmt.Println("\nUsuários cadastrados:")
	for _, usuario := range s.usuarios {
		fmt.Printf("- Nome: %s | Email: %s | Idade: %d\n", usuario.Nome, usuario.Email, usuario.Idade)
	}
}

func (s *SistemaUsuarios) salvar() {
	dados, err := json.MarshalIndent(s.usuarios, "", "")
	if err != nil {
		fmt.Println("Erro ao gerar JSON:", err)
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
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	err = json.Unmarshal(dados, &s.usuarios)
	if err != nil {
		fmt.Println("Erro ao carregar JSON:", err)
	}
}

func lerLinha(reader *bufio.Reader) string {
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func main() {
	sistema := NovoSistemaUsuarios("usuarios.json")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU DE USUÁRIOS ===")
		fmt.Println("1 - Cadastrar usuário")
		fmt.Println("2 - Listar usuário")
		fmt.Println("3 - Buscar usuário")
		fmt.Println("4 - Atualizar usuário")
		fmt.Println("5 - Remover usuário")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao := lerLinha(reader)

		switch opcao {
		case "1":
			fmt.Print("Nome: ")
			nome := lerLinha(reader)

			fmt.Print("Email: ")
			email := lerLinha(reader)

			fmt.Print("Idade: ")
			idadeTexto := lerLinha(reader)

			idade, err := strconv.Atoi(idadeTexto)
			if err != nil {
				fmt.Println("Idade inválida.")
				continue
			}

			mensagem := sistema.Cadastrar(Usuario{
				Nome:  nome,
				Email: email,
				Idade: idade,
			})
			fmt.Println(mensagem)

		case "2":
			sistema.Listar()

		case "3":
			fmt.Print("Digite o nome para buscar: ")
			nome := lerLinha(reader)

			usuario, encontrado := sistema.Buscar(nome)
			if !encontrado {
				fmt.Println("Usuário não enontrado.")
				continue
			}

			fmt.Printf("Usuário encontrado: Nome=%s | Email=%s | Idade=%d\n", usuario.Nome, usuario.Email, usuario.Idade)

		case "4":
			fmt.Print("Nome do usuário para atualizar: ")
			nome := lerLinha(reader)

			fmt.Print("Novo email: ")
			email := lerLinha(reader)

			fmt.Print("Nova idade: ")
			idadeTexto := lerLinha(reader)
			idade, err := strconv.Atoi(idadeTexto)
			if err != nil {
				fmt.Println("Idade inválida.")
				continue
			}

			fmt.Println(sistema.Atualizar(nome, email, idade))

		case "5":
			fmt.Print("Nome do usuário para remover: ")
			nome := lerLinha(reader)

			fmt.Println(sistema.Remover(nome))

		case "0":
			fmt.Println("Saindo do sistema...")
			return

		default:
			fmt.Println("Opção inválida.")
		}

	}

}

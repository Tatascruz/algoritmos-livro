# Capítulo 3 - Recursão

## O que é recursão?

Recursão é quando uma função chama a si mesma para resolver um problema até atingir
uma condição de parada chamada **caso base**.

Toda função recursiva precisa de duas partes fundamentais:

- **Caso base** -> condição que faz a função para
- **Passo recursivo** -> chamada que aproxina do caso base

Sem um caso base a função entra em chamadas infinitas,

---

## Estrutura Geral de uma Função Recursiva

```go
func Exemplo(n int) {
    if condicaoBase {
        return
    }

    Exemplo(n - 1)
}


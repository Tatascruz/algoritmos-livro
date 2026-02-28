# Capítulo 3 - Recursão

## O que é recursão?

Recursão é quando uma função chama a si mesma para resolver um problema até atingir
uma condição de parada chamada **caso base**.

Toda função recursiva precisa de duas partes fundamentais:

- **Caso base** -> condição que faz a função para
- **Passo recursivo** -> chamada que aproxina do caso base

Sem um caso base a função entra em chamadas infinitas,

---

## Conceitos aprendidos

- Caso base 
- Chamadas recursivas
- Pilhas de execusão (call stack)

## Exemplo implementados

## Contagem Recursiva

```go
func ContarRecursivo(lista [] int) {
    if len(lista) == 0 {
        return 0
    }

    return 1 + ContarRecursivo(lista[1:])
}
```
Este exemplo mostra como contar elementos de uma lista usando recursão.

## Visualização da pilha 

Descendo:

Descendo: 2
Descendo: 1
Chegou no caso base (0)

Subindo:

Subindo: 2
Subindo: 1
Fim da execusão

Isso mostra como a função entra e sai da pilha de execução.

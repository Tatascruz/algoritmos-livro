# Caítulo 2 - Ordenação

Este diretório contém os estudos e implementação do Capítulo 2 do livro *Entendendo Algoritmos*

## Conceitos estudados

- Ordenação de lista
- Selection Sort
- Comparações e trocas
- Análise de complexidade
- Diferença entre algoritmos O(n) e O(n²)

---

## O que aprendi

### Selection Sort

OSelection Sort percorre a lista buscando o menor elemento e o coloca na posição correta.
Ele repete esse processo até que toda a lista esteja ordenada.

### Complexidade

- Tempo **O(n²)**
- Espaço **O(1)**

Mesmo funcinando bem para a lista pequenas, o número de comparações cresce rapida conforme a lista aumenta.

---

## Experimentos realizados

- Contagem de comparações 
- Contagem de trocas
- Teste com diferentes tamanhos de lista

Exemplo observado:
- Lista com 10 elementos -> 45 comparações

Isso mostra como o número de operações cresce rapidamente.

---

## Arquivos deste capítulo

- 'selection_sort.go'
- 'selection_sort_contador.go'

---

## Próximos passos

- Estudar algoritmos de ordenação mais eficientes
- Comparar com algoritmos O(n log n)

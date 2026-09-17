# Sliding Window

**Sliding Window** é uma técnica usada para trabalhar com **subarrays ou substrings contíguos**, mantendo uma "janela" sobre uma parte do Array e movendo essa janela conforme percorremos os elementos.

Ela é muito comum em problemas de:

- Arrays
    
- Strings
    
- Subarrays
    
- Substrings
    
- Soma de elementos
    
- Maior/menor intervalo
    
- Frequência de elementos
    
- Elementos distintos
    

---

# 1. A ideia

Imagine:

```text
[2, 1, 5, 1, 3, 2]
```

Queremos encontrar a **maior soma de 3 elementos consecutivos**.

A primeira janela é:

```text
[2, 1, 5, 1, 3, 2]
 └─────┘
   2+1+5 = 8
```

Depois deslizamos a janela uma posição:

```text
[2, 1, 5, 1, 3, 2]
    └─────┘
    1+5+1 = 7
```

Depois:

```text
[2, 1, 5, 1, 3, 2]
       └─────┘
       5+1+3 = 9
```

Depois:

```text
[2, 1, 5, 1, 3, 2]
          └─────┘
          1+3+2 = 6
```

A maior soma é:

```text
9
```

com:

```text
[5, 1, 3]
```

---

# 2. O problema da abordagem ingênua

Uma abordagem simples seria calcular cada janela novamente:

```text
2 + 1 + 5
1 + 5 + 1
5 + 1 + 3
1 + 3 + 2
```

Observe que estamos recalculando elementos.

Por exemplo:

```text
2 + 1 + 5
    ↓ ↓
    1 + 5 + 1
```

`1` e `5` foram calculados novamente.

Sliding Window evita isso.

---

# 3. A janela "desliza"

Em vez de calcular tudo novamente, fazemos:

```text
nova soma = soma atual
             - elemento que saiu
             + elemento que entrou
```

Primeira janela:

```text
[2, 1, 5] 1 3 2
```

Soma:

```text
8
```

Movemos:

```text
2 [1, 5, 1] 3 2
```

Saiu:

```text
2
```

Entrou:

```text
1
```

Então:

```text
8 - 2 + 1 = 7
```

Próxima:

```text
2 1 [5, 1, 3] 2
```

Sai:

```text
1
```

Entra:

```text
3
```

Então:

```text
7 - 1 + 3 = 9
```

Essa é a essência da técnica.

---

# 4. Implementação em Go

Problema:

> Encontrar a maior soma de `k` elementos consecutivos.

```go
func maxWindowSum(numbers []int, k int) int {
    windowSum := 0

    for i := 0; i < k; i++ {
        windowSum += numbers[i]
    }

    maxSum := windowSum

    for right := k; right < len(numbers); right++ {
        left := right - k

        windowSum -= numbers[left]
        windowSum += numbers[right]

        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}
```

Uso:

```go
numbers := []int{2, 1, 5, 1, 3, 2}

result := maxWindowSum(numbers, 3)

fmt.Println(result) // 9
```

---

# 5. Implementação em Node/TypeScript

```ts
function maxWindowSum(numbers: number[], k: number): number {
  let windowSum = 0;

  for (let i = 0; i < k; i++) {
    windowSum += numbers[i];
  }

  let maxSum = windowSum;

  for (let right = k; right < numbers.length; right++) {
    const left = right - k;

    windowSum -= numbers[left];
    windowSum += numbers[right];

    maxSum = Math.max(maxSum, windowSum);
  }

  return maxSum;
}
```

Uso:

```ts
const numbers = [2, 1, 5, 1, 3, 2];

const result = maxWindowSum(numbers, 3);

console.log(result); // 9
```

---

# 6. Complexidade

A parte interessante:

```text
Array:
[2, 1, 5, 1, 3, 2]
```

Percorremos cada elemento essencialmente uma vez.

Portanto:

```text
Tempo → O(n)
Espaço → O(1)
```

Compare:

```text
Abordagem ingênua
→ recalcula cada janela
→ O(n × k)
```

versus:

```text
Sliding Window
→ adiciona/remove elementos da janela
→ O(n)
```

Se `k` for grande, a diferença pode ser enorme.

---

# 7. Fixed Window vs Dynamic Window

Existem **dois padrões principais** de Sliding Window.

## Fixed-size Window

O tamanho da janela é fixo.

Exemplo:

> Maior soma de `k` elementos consecutivos.

```text
[2, 1, 5, 1, 3, 2]
 └─────┘
    k=3
```

Sempre temos:

```text
window size = 3
```

---

## Dynamic Window

O tamanho da janela pode aumentar ou diminuir de acordo com uma condição.

Exemplo:

> Encontrar o menor subarray cuja soma seja pelo menos `7`.

Array:

```text
[2, 3, 1, 2, 4, 3]
```

Começamos expandindo:

```text
[2]
[2, 3]
[2, 3, 1]
[2, 3, 1, 2]
```

Quando a soma chega a pelo menos `7`, começamos a diminuir a janela:

```text
[2, 3, 1, 2] → soma 8
 ↑
remove 2

[3, 1, 2] → soma 6
```

A janela volta a crescer.

O padrão é:

```text
expand → condição atingida → shrink → condição perdida → expand
```

---

# 8. O padrão `left` e `right`

Sliding Window normalmente utiliza dois ponteiros:

```text
left
right
```

Por exemplo:

```text
[2, 3, 1, 2, 4, 3]
 ↑           ↑
left        right
```

A região entre eles representa nossa janela:

```text
[2, 3, 1, 2]
 └─────────┘
   window
```

Movemos `right` para expandir:

```text
[2, 3, 1, 2, 4]
 └───────────┘
```

Movemos `left` para diminuir:

```text
   [3, 1, 2, 4]
    └────────┘
```

Isso é muito parecido com **Two Pointers**.

---

# 9. Sliding Window vs Two Pointers

Essa distinção é importante.

**Two Pointers** é uma técnica mais geral:

```text
left ───────→
            ←──── right
```

Pode ser usada em vários tipos de problemas.

**Sliding Window** normalmente representa uma **região contínua entre dois ponteiros**:

```text
left
 ↓
[ A B C D ]
         ↑
       right
```

Então podemos pensar:

```text
Two Pointers
     │
     └── Sliding Window
```

Mas nem todo problema de Two Pointers é Sliding Window.

---

# 10. Sliding Window + Hash Map

Agora começa uma combinação muito poderosa.

Imagine o problema:

> Encontrar a maior substring sem caracteres repetidos.

Exemplo:

```text
"abcabcbb"
```

Podemos manter:

```text
left
right
```

e uma estrutura para saber quais caracteres estão na janela:

```text
Set
```

A janela começa:

```text
[a]
```

Depois:

```text
[a b]
```

Depois:

```text
[a b c]
```

Ao encontrar outro `a`:

```text
[a b c a]
 ↑       ↑
```

Temos uma repetição.

Movemos `left` até a janela voltar a ser válida:

```text
[a b c a]
   ↑     ↑
 left   right
```

Depois:

```text
[b c a]
```

Esse padrão aparece muito em problemas de entrevistas.

---

# 11. Quando pensar em Sliding Window?

Quando o problema mencionar algo parecido com:

> **subarray/substr contínuo + maior/menor/máximo/mínimo/quantidade + uma condição**

comece a considerar Sliding Window.

Exemplos:

```text
Maior soma de k elementos consecutivos
Menor subarray com soma >= target
Maior substring sem caracteres repetidos
Maior substring contendo no máximo k caracteres distintos
Quantidade de elementos dentro de uma janela
```

A pergunta mental é:

```text
Existe um intervalo contínuo
que posso representar com
left e right?
        ↓
       SIM
        ↓
Posso mover a janela sem
recalcular tudo?
        ↓
       SIM
        ↓
Sliding Window
```

---

# 12. A conexão com o que você estudou

E existe uma diferença interessante:

```text
Prefix Sum
→ pré-processa o Array
→ consultas de soma O(1)

Sliding Window
→ mantém uma janela durante o percurso
→ geralmente O(n)

Kadane
→ mantém a melhor soma terminando na posição atual
→ O(n)
```


> **Sliding Window é uma técnica para processar subarrays ou substrings contíguos mantendo uma janela entre dois ponteiros (`left` e `right`). A janela pode ter tamanho fixo ou variável. Em vez de recalcular os elementos a cada movimento, atualizamos o estado da janela incrementalmente, frequentemente reduzindo uma solução de O(n × k) para O(n).**
# Prefix Sum

**Prefix Sum** é uma técnica usada para pré-calcular somas acumuladas de um Array, permitindo responder rapidamente a consultas de soma de intervalos.

Diferente do **Kadane**, que procura a maior soma de um subarray, o **Prefix Sum** é principalmente uma técnica para **calcular somas de intervalos de forma eficiente**.

---

## 1. O problema

Imagine o Array:

```text
[2, 4, 1, 5, 3]
```

E queremos descobrir várias vezes a soma de determinados intervalos.

Por exemplo:

```text
soma do índice 1 até 3

[2, 4, 1, 5, 3]
    └──────┘
    4 + 1 + 5 = 10
```

Uma forma simples seria percorrer o intervalo:

```text
4 + 1 + 5
```

Isso custa `O(n)` para cada consulta no pior caso.

Se tivermos **milhares de consultas**, isso pode ficar caro.

Prefix Sum resolve isso fazendo um pré-processamento.

---

# 2. Construindo o Prefix Sum

Array original:

```text
[2, 4, 1, 5, 3]
```

Calculamos a soma acumulada:

```text
Índice:       0   1   2   3   4
Original:     2   4   1   5   3
              ↓   ↓   ↓   ↓   ↓
Prefix Sum:   2   6   7  12  15
```

Porque:

```text
prefix[0] = 2

prefix[1] = 2 + 4 = 6

prefix[2] = 2 + 4 + 1 = 7

prefix[3] = 2 + 4 + 1 + 5 = 12

prefix[4] = 2 + 4 + 1 + 5 + 3 = 15
```

Então:

```text
prefix[i] = soma de todos os elementos de 0 até i
```

---

# 3. A grande vantagem

Agora queremos:

```text
soma do índice 1 até 3
```

Original:

```text
[2, 4, 1, 5, 3]
    └──────┘
     10
```

No Prefix Sum:

```text
prefix[3] = 12
prefix[0] = 2
```

Então:

```text
12 - 2 = 10
```

Ou seja:

```text
sum(1, 3) = prefix[3] - prefix[0]
```

Conseguimos a resposta sem percorrer `4`, `1` e `5`.

---

# 4. A fórmula

Para um intervalo:

```text
[left, right]
```

usamos:

```text
sum(left, right) =
    prefix[right] - prefix[left - 1]
```

Por exemplo:

```text
Array:
[2, 4, 1, 5, 3]

Prefix:
[2, 6, 7, 12, 15]

left  = 1
right = 3
```

Então:

```text
prefix[3] - prefix[0]

12 - 2

= 10
```

---

# 5. O caso `left = 0`

Aqui existe uma pequena complicação.

Se quisermos:

```text
sum(0, 3)
```

não podemos fazer:

```text
prefix[3] - prefix[-1]
```

Por isso existe uma implementação ainda melhor: criar um Prefix Sum com **um elemento extra no início**.

Array:

```text
[2, 4, 1, 5, 3]
```

Prefix:

```text
[0, 2, 6, 7, 12, 15]
 ↑
 extra
```

Agora temos:

```text
prefix[i + 1] = prefix[i] + array[i]
```

A fórmula fica muito mais simples:

```text
sum(left, right) =
    prefix[right + 1] - prefix[left]
```

---

# 6. Exemplo

Queremos:

```text
sum(1, 3)
```

Array:

```text
[2, 4, 1, 5, 3]
    └──────┘
```

Prefix:

```text
[0, 2, 6, 7, 12, 15]
```

Aplicando:

```text
prefix[right + 1] - prefix[left]

prefix[4] - prefix[1]

12 - 2

= 10
```

Perfeito.

---

# 7. Implementação em Go

### Construindo o Prefix Sum

```go
func buildPrefixSum(numbers []int) []int {
    prefix := make([]int, len(numbers)+1)

    for i := 0; i < len(numbers); i++ {
        prefix[i+1] = prefix[i] + numbers[i]
    }

    return prefix
}
```

Uso:

```go
numbers := []int{2, 4, 1, 5, 3}

prefix := buildPrefixSum(numbers)

fmt.Println(prefix)
// [0 2 6 7 12 15]
```

### Consultando um intervalo

```go
func rangeSum(prefix []int, left int, right int) int {
    return prefix[right+1] - prefix[left]
}
```

Uso:

```go
result := rangeSum(prefix, 1, 3)

fmt.Println(result)
// 10
```

---

# 8. Implementação em Node/TypeScript

### Construindo

```ts
function buildPrefixSum(numbers: number[]): number[] {
  const prefix = new Array(numbers.length + 1).fill(0);

  for (let i = 0; i < numbers.length; i++) {
    prefix[i + 1] = prefix[i] + numbers[i];
  }

  return prefix;
}
```

Uso:

```ts
const numbers = [2, 4, 1, 5, 3];

const prefix = buildPrefixSum(numbers);

console.log(prefix);
// [0, 2, 6, 7, 12, 15]
```

### Consultando

```ts
function rangeSum(
  prefix: number[],
  left: number,
  right: number
): number {
  return prefix[right + 1] - prefix[left];
}
```

Uso:

```ts
const result = rangeSum(prefix, 1, 3);

console.log(result);
// 10
```

---

# 9. Complexidade

Aqui está a grande sacada do Prefix Sum.

### Construção

Precisamos percorrer o Array uma vez:

```text
O(n)
```

### Consulta

Depois que o Prefix Sum foi construído:

```text
prefix[right + 1] - prefix[left]
```

São apenas operações constantes:

```text
O(1)
```

Então:

|Operação|Complexidade|
|---|--:|
|Construir Prefix Sum|**O(n)**|
|Consultar intervalo|**O(1)**|
|Espaço adicional|**O(n)**|

A troca é:

```text
Mais memória
     ↓
O(n) espaço
     ↓
Consultas muito mais rápidas
     ↓
O(1)
```

---

# 10. Por que isso é tão útil?

Imagine:

```text
Array com 1.000.000 elementos
```

E você recebe:

```text
100.000 consultas
```

Sem Prefix Sum:

```text
cada consulta → O(n)
100.000 consultas → potencialmente O(n × q)
```

Com Prefix Sum:

```text
construção → O(n)

cada consulta → O(1)

total → O(n + q)
```

Isso é uma diferença enorme.

---

# 11. Prefix Sum vs Kadane

Como você acabou de estudar **Kadane**, é bom diferenciar:

||Prefix Sum|Kadane|
|---|---|---|
|Objetivo|Somar intervalos rapidamente|Encontrar maior soma de subarray|
|Ideia|Soma acumulada|Melhor soma terminando em cada posição|
|Tempo|`O(n)` + consultas `O(1)`|`O(n)`|
|Espaço|`O(n)`|`O(1)`|
|Principal uso|Range Sum|Maximum Subarray|

Exemplo:

```text
[2, 4, 1, 5, 3]
```

Prefix Sum pergunta:

> "Qual é a soma entre `left` e `right`?"

Kadane pergunta:

> "Qual é o subarray contíguo com a maior soma?"

São problemas diferentes.

---

# 12. Onde você vai encontrar Prefix Sum

Depois de dominar o conceito básico, ele aparece em problemas como:

```text
Range Sum
Subarray Sum
Subarray Sum Equals K
Contagem de subarrays
Difference Array
Prefix XOR
Prefix Product
2D Prefix Sum
```

E existe uma conexão muito importante com Hash Map:

```text
Prefix Sum + Hash Map
```

que permite resolver problemas de **subarray** de maneira muito eficiente.

---

anotação

> **Prefix Sum é uma técnica que pré-calcula somas acumuladas de um Array. Após o pré-processamento em O(n), a soma de qualquer intervalo pode ser obtida em O(1) através da diferença entre duas somas prefixadas. A técnica troca espaço adicional O(n) por consultas de intervalo mais eficientes.**

E **Prefix Sum é especialmente importante antes de começar `Subarray Sum`**, porque muitos problemas de entrevistas/LeetCode usam exatamente essa ideia.
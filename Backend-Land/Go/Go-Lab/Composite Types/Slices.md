# Conceitos Fundamentais

## Slices vs Arrays:
    - Podemos delcar um slice sem especificar o tamanho: var x = []int{10, 20, 30}
    - Usando [...] cria um array; usando [] cria um slice
    - A grande vantagem é que slcies podem crescer conforme necessário.

    Uma slice pode ser considerado um array dinâmico em go?

## Nil Slices:
    - Um slice declarado sem valor recebe nil como zero value
    - nil é diferente de outros linguagens (tipo null)
    - Um slice nil contém nada e len(nil_slice) retorna 0

# Operações Principais

## len - tamanho do slice
    - Funciona como em arrays
    - Retorna 0 para nil slices

## append - adiciona elementos
    - Retorna um novo slice e deve ser atribuído de volta: x = append(x, 10)
    - Pode adicionar um ou múltiplos valores: x = append(x, 5, 6, 7)
    - Pode combinar slices: x = append(x, y...)

## Capacity
    - Todo slice tem uma length (elementos atribuídos) e capacity (espaço reservado)
    - Quando length == capacity e você faz append, Go aloca um novo backing array maior
    - Go copia os dados antigos para um novo array (comportamento de array dinâmico)

## make - criar slices pré-alocados
    - make([]int, 5) - cria um slice com 5 = 5
    - make([]int, 5, 10) - cria um slice com 5 para 10
    - make([]int, 0, 10) - cria um slice com 0 para 10 (útil para append)

## copy - copia elementos entre slices

## Comparação:
    - Não pode usar == para comparar slices (erro de compilação)
    - Pode comparar com nil
    - Go 1.21+ oferece slices.Equal() e slices.EqualFunc()

## Limpeza:
    - Go 1.21+ adicionou clear(s) que zera todos os elementos

# Boas Práticas de Declaração

1. Nil slice: quando o slice pode nunca crescer
2. Slice literal: quando tem valores iniciais
3. make com capacity: quando sabe o tamanho aproximado
4. make com length 0 e capacity: preferível para usar com append (evita zero values surpresa)

# Array -> Slice

## Sintaxe básica:
```go
xArray := [4]int{5, 6, 7, 8}
xSlice :=Array[:] // Converte o array inteiro em slice
```
## Convertendo um subconjunto:
```go
x := [4]int{5,6, 7, 8}
y := x[:2] // Primeiros 2 elementos
z := x[2:] // Últimos 2 elementos
```
**Importante**: Compartilhamento de Memória

Quando você cria um slice a partir de um array, elas compartilham a mesma memória.

Alterações em uma afetam a outra:
```go
x := [4]int{5, 6, 7, 8}
y := x[:2] // y = [5, 6] com capacity 4
z := x[2:] // z = [7, 8]

x[0] = 10 // alteração no array original

fmt.Println("x:", x) // [10 6 7 8]
fmt.Println("y:", y) // [10, 6]
fmt.Println("z:", z) // [7 8]
```

# Slice -> Array

## Sintaxe:
```go
xSlice := []int{1, 2, 3, 4}
xArray := [4]int(xSlice) //Converte em array de 4 elementos
```
**Importante**: Dados são COPIADOS

Diferente de array -> slice, quando converte slice -> array, os dados são copiados para uma nova memória. Alterações não afetam uma à outra:
```go
xSlice := []int{1, 2, 3, 4}
xArray := [4]int(xSlice)

xSlice[0] = 10

fmt.Println(xSlice) // [10 2 3 4] <- mudou
fmt.Println(xArray) // [1 2 3 4] <- não mudou (cópia)
```
## Convertendo um subconjunto:
```go
xSlice := []int{1, 2, 3, 4}
smallArray := [2]int(xSlice) // Copia apenas os 2 primeiros

fmt.Println(smalArray) // [1 2]
```
## Restrições:

1. O tamanho deve ser conhecido em tempo de compilação
2. Tamanho do array <= tamanho do slice
    - array menor que slice -> funciona, copia parcialmente
    - array maior que slice -> PANIC em runtime

```go
xSlice := []int{1, 2, 3, 4}
panicArray := [5]int(xSlice) // Panic
// panic: runtime error: cannot convert slice with length 4
// to array or pointer to array with length 5
```

# Conversão de Pointer de Array

Podemos também converter slice em pointer para array, o que compartilha memória:
```go
xSlice := []int{1, 2, 3, 4}
xArrayPointer := (*[4]int)(xSlice)

XSlice[0] = 10
xArrayPointer[1] = 20

fmt.Println(xSlice)
fmt.Println(xArrayPointer) // &[10 20 3 4] <- mesma memória
```

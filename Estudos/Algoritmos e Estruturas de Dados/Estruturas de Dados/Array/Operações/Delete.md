Remover um elemento

No final

Go
```
numbers := []int{10, 20, 30, 40, 50}

numerbs = numbers[:len(numbers)-1] // -1 remove o último índice
```
Typescript
```
const numbers = [10, 20, 30, 40, 50]

numers.pop() //vazio remove o último indice
```

Complexidade: O(1)

No meio
Go
```
numbers := []int{10, 20, 30, 40, 50}

index := 2
numbers := append(numbers[:index], numbers[index+1:]...)

//Resultado: [10, 20, 40, 50]
```
Typescript
```
const numbers = [10, 20, 30, 40, 50]

numbers.pop(2)

//Resultado: [10, 20, 40, 50]
```

Complexidade: O(n)

[[Acesso O(1)]]

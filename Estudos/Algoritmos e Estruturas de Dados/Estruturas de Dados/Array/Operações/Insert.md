Adiciona um elemento:

No final:

Go
```
numbers := []int{10, 20, 30, 40, 50}

numbers = append(numbers, 60)
```

Typescript
```
const numbers = [10, 20, 30, 40, 50];

numbers.push(60);
```

Complexidade: O(1) amortizado. São arrays dinâmicos que vão ocasionalmente precisar crescer para capacitar o tamanho do array.

No inicio:

Go
```
numbers := []int{10, 20, 30, 40, 50}

numbers = append([]int{5}, numbers...)
```

Typescript
```
const numbers = [10, 20, 30, 40, 50]

numbers.unshift(5)
```

Complexidade: O(n), porque os elementos precisam ser deslocados.



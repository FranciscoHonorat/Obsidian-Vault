Percorrer os elementos de um array 

Exemplo:

Go
```
numbers := []int{10, 20, 30, 40, 50}

for i := 0; i < len(numbers); i++ {
	fmt.Println(numbers[i])
}
```

Typescript
```
const numbers = [10, 20, 30, 40, 50];

for (let i = 0; i < numbers.length; i++){
	console.log(numbers[i]);
}
```

Complexidade: O(n)

A ideia principal é que Traversal percorre os elementos do array para ler, processar ou verificar cada elemento.

Sua complexidade é O(n), pois todo ato de percorrer um array significa fazer a leitura de todos os elementos.
Binary Search é um algoritmo de busca que encontra um elemento em um coleção ordenada, eliminando metade dos elementos candidatos a cada passo.

É importante separar: 

Binary Search é um algoritmo/tecnica de busca, não uma operação básica do Array.

A ideia é considere um array ordenado:

```
[10, 20, 30, 40, 50, 60, 70]
  0   1   2   3   4   5   6
```

Queremos encontrar 60: ```target = 60``` e em vez de percorrer todo array para encontrar o valor, vamos seguir a ideia de uma buscar linear, olhando para um valor próximo daquele que queremos. Então logo podemos usar o meio do array que é 40 que é um valor menor que 60 e com isso descartamos todos os numeros a esquerda e ficamos com ```[50, 60, 70]``` e então olhamos para o meio novamente que é 60 o nosso alvo.

A implementação tradicional usa três variáveis: left, mid e right.
Inicialmente
```
[10, 20, 30, 40, 50, 60, 70]
 ↑              ↑              ↑
left           mid            right
 0               3              6
```
Calculamos:
```
mid = left + (right - left) / 2
```
Depois comparamos:
```
array[mid] com target
```

Com isso existe três possibilidades:

```
//Encontrou:
array[mid] == target // retornamos o índice

//Se o target está à direita

array[mid] < target

left = mid + 1

//Se o target está à esquerda
array[mid] > target

right = mid - 1
```

Implementação em Go
```
func binarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1
	
	for <= left right {
		mid := left + (right-left)/2
		
		if numbers[mid] == target {
			return mid
		}
		
		if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	
	index := binarySearch(numbers, 60)
	
	fmt.Println(index) // 5
}

```

Implementação em Typescript
```
function binarySearch(numbers: numbers[], target: number): number {
	let left = 0
	let right = numbers.lenght -1;
	
	while (left <= right) {
		const mid = left + Math.floor((right - left) / 2);
		
		if (numbers[mid] === target) {
			return mid;
		}
		
		if (numbers[mid] < target) {
			left = mid + 1;
		} else {
			right = mid - 1;
		}
	}
	
	return -1;
}

//Uso

const numbers = [10, 20, 30, 40, 50, 60, 70]

const index = binarySearch(numbers, 60);

console.log(index) // 5
```

Por que é O(log n)?

Essa é a parte mais importante da complexidade

Imagine um Array com: 1.000.000 elementos, a cada comparação elimina aproximadamente metade:
```
1.000.000
    ↓
500.000
    ↓
250.000
    ↓
125.000
    ↓
62.500
    ↓
...
    ↓
1
```
Quantas vezes conseguimos dividir por 2?
```
//Aproximadamente:
log₂(1.000.000) ≈ 20
```
Então, mesmo com 1 milhão de elementos, precisamos de aproximadamente 20 comparações no pior caso.

Por isso:

Binary Search -> O(log n)

Enquanto:

Linear Search -> O(n)

Por que o Array precisa estar ordenado?

Porque o algoritmo toma decisões baseado na ordenação, ou seja se temos: ```[10, 20, 30, 40, 50, 60, 70]```  e ```array[mid] = 40 target = 60``` sabemos que 40 < 60, logo podemos afirmar que tudo à esquerda de 40 também é menor que 60. Então podemos descartar a metade esquerda. Mas se o array fosse ```[50, 10, 70, 20, 60, 30, 40] ``` essa conclusão não seria válida.

Fórmula do mid

mid = (left + right) / 2

Forma mais robusta

mid = left + (right - left) / 2 

Em Go:
```
mid := left + (right-left)/2
```

Em Typescript
```
const mid = left + Math.floor((right - left) / 2);
```
A segunda forma evita **overflow de inteiro** em linguagens onde isso pode ser um problema quando `left + right` é muito grande.

**Binary Search não serve apenas para "encontrar um número".**

Esse padrão aparece em muitos problemas mais avançados, por exemplo:

- encontrar a primeira ocorrência;
- encontrar a última ocorrência;
- encontrar a posição de inserção;
- procurar um limite (`lower bound` / `upper bound`);
- buscar a resposta mínima/máxima possível;
- **Binary Search on Answer**.

Então o conceito fundamental que vale guardar agora é:

> **Binary Search reduz repetidamente pela metade um espaço de busca ordenado, usando a comparação com o elemento do meio para decidir qual metade pode ser descartada. Sua complexidade é O(log n).**

E ele se conecta diretamente ao que você acabou de estudar:

```
Array
  ↓
Indexação
  ↓
Acesso por índice O(1)
  ↓
left / mid / right
  ↓
Binary Search
  ↓
O(log n)
```

Essa conexão entre **indexação + acesso O(1) + `left/mid/right`** é a base para começar a resolver problemas de Binary Search.

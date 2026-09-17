Qual sua função?

é um algoritmo clássico para encontrar a maior soma de um subarray contíguo.

O que isso que dizer?

Significa que ele é especialmente importante para resolver problemas de array e se utiliza uma ideia de programação dinâmica.

Considere o problema:

```[-2, 1, -3, 4, -1, 2, 1, -5, 4]```

Queremos encontrar o subarray contíguo cuja soma seja a maior possível

A resposta é:

```
[4, -1, 2, 1]
```

Pois

```
4 + (-1) + 2 + 1 = 6
```

Observe que o subarray precisa ser contíguo, então não podemos simplesmente escolher:

%%Um subarray contíguo é uma sequeência de elementos consecutivos extraídos de um array principa, mantendo a ordem original sem pular nenhuma posição.
Contíguo significa sem interrupção, os elementos do subarry devem estar lado a lado na memória e na ordem do array original, não permitindo pegar elementos separados ou pular índices.%%

```
4 + 2 + 1 + 4
```

Pois esses elementos não formam um intervalo contíguo

A ideia central é a seguinte pergunta que Kadane faz a cada elemento é:

	 É melhor continuar o subarray atual ou começa um novo subarray aqui

Imagine:
```
[-2, 1, -3, 4, -1, 2, 1]
```
Quando chegamos em 4, temos uma decisão:

subarray anterior + 4  ou começa um novo subarray com 4

Kadane calcula:
```
currentSum = max(
	numero_atual,
	currentSum + numero_atual
)
```
E mantém também: maxSum que representa a maior soma encontrada aé aquele momento.

Implementação em GO

```
func maxSubArray(numbers []int) int {
	currentSum := numbers[0]
	maxSum := numbers[0]
	
	for i := i; i < len(numbers); i++ {
		currentSum = max(numbers[i], currentSum+numbers[i])
		maxSum = max(maxSum, currentSum)
	}
	
	return maxSum
}

func main() {
	numbers := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	
	result := maxSubArray(numbers)
	
	fmt.Println(result) // 6
}
```

Uma forma sem depender de max:

```
func maxSubArray(numbers []int) int {
    currentSum := numbers[0]
    maxSum := numbers[0]

    for i := 1; i < len(numbers); i++ {
        if currentSum+numbers[i] < numbers[i] {
            currentSum = numbers[i]
        } else {
            currentSum += numbers[i]
        }

        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return maxSum
}
```

Implementação em TS
```
function maxSubArray(numbers: number[]): number {
	let currentSum = numbers[0];
	let maxSum = numbers[0];
	
	for (let i = 1; i < numbers.lenght; i++) {
		currentSum = Math.maax(numbers[i], currentSum + numbers[i]);
		maxSum = Math.max(maxSum, CurrentSum);
	}
	
	return maxSum;
}

//Uso
const numbers = [-2, 1, -3, 4, -1, 2, 1, -5, 4];

const result = maxSubArray(numbers);

console.log(result); // 6
```

A complexidade:

O algoritmo percorre o Array uma única vez:
```
[-2, 1, -3, 4, -1, 2, 1, -5, 4]
 ↑
 percorremos uma vez 
```
Portanto:
Tempo  → O(n)
Espaço → O(1)

Isso é excelente. Uma abordagem ingênua poderia verificar todos os subarrays, o que teria um complexidade muito maior e o Kadane reduz o problema para: uma passagem -> O(n)

Kadane's Algorithm encontra a maior soma possível de um subarray contíguo em O(n) e O(1) de espaço. Para cada elemento, decide entre continuar o subarray atual (`currentSum + element`) ou iniciar um novo subarray naquele elemento (`element`). A variável `maxSum` mantém a maior soma encontrada durante o percurso.

Kadane é um ótimo exemplo de como um simples `Traversal` de Array pode carregar um estado (`currentSum`) para resolver um problema que, à primeira vista, parece exigir testar vários subarrays.



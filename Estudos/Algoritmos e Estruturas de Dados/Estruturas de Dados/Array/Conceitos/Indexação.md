O que é indexação?

Indexação é a forma de acessar um elemento de um array usando sua posição.

Geralmente na programação o primeiro índice normalmente é 0

```
array: [10, 20, 30, 40, 50]
Índice: 0,  1,   2,  3,  4
```

A ideia fundamental é: O índice representa a posição do elemento dentro do array, começando em 0

Por que começa em 0?

Imagine um array armazenado na memória: 

Endereço inicial = 100

Se cada elemento ocupa 4 bytes:
```
Elemento   Endereço
10           1000
20           1004
30           1008
40           1012
50           1016
```
O computador pode calcular diretamente o endereço de um elemento:

endereço = endereço_inicial + indice x tamanho_do_elemento

por exemplo, para acessar o elemento ```array[3]```
```
1000 + 3 x 4

1012
```
Por isso o acesso é O(1).

O acesso por índice é uma das grandes vantagens de um array

```
int[] numbers = {10, 20, 30, 40, 50};
System.out.println(numers[3]);

//Resultado 40
```

O computador consegue ir direto ao elemento de índice 3, diferente de uma estrutura sequencial como uma linked list que normalmente seria necessário percorrer os nós até cehgar à posição desejada.

Se um array possui ```n``` elementos:
```
primeiro índice = 0
último índice = n - 1

//Por exemplo:

n = 5

índices: 
1, 2, 3, 4

//Logo:

primeiro = 0
último = 5 - 1 = 4
```

Essa fórmula é fundamental para algoritmos, podemos encontrar constantemente em:

loops, two pointers, binary search, slinding window, prefix sum, ordenção, busca.

Se o índice estiver fora dos limites vai gerar um erro, a regra geralmente é
```
0 <= índice < tamanho
```

Percorrendo um array é onde indexação começa a aparecer nos algoritmos:
```
int[] numbers = {10, 20, 30, 40, 50};

for (int i = 0; i < numbers.length; i++) {
    System.out.println(numbers[i]);
}
```

Observe que:

```
i = 0 → numbers[0]
i = 1 → numbers[1]
i = 2 → numbers[2]
i = 3 → numbers[3]
i = 4 → numbers[4]
```

Quando: 
```
i = 5
```
a condição: i < numbers.length ficou falsa. Por isso usamos i < length 

Indexação e complexidade:
|Operação|Complexidade|
|---|--:|
|Acessar `array[i]`|**O(1)**|
|Atualizar `array[i]`|**O(1)**|
|Percorrer array|**O(n)**|
|Buscar elemento não ordenado|**O(n)**|
|Inserir no final*|**O(1)** amortizado|
|Inserir no início/meio|**O(n)**|
|Remover do início/meio|**O(n)**|

Em arrays dinâmicos, considerando espaço disponível e eventualmente pode ocorrer redimensionamento. O ponto central é:

Array possui acesso aleatório (random access) por índice em O(1)

Resumo do conceito central:

Indexação é o mecanismo que permite acessa diretamente um elemento do Array através de seu índice. Em um array de tamanho n, os índices válidos vão de 0 até n - 1. O acesso por índice possui complexidade O(1), pois o endereço do elemento pode ser calculado diretamente a partir do índice.

#Conceitos
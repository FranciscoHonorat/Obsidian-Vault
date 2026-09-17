A diferença prinicpal é se o tamanho da estrutura pode mudar depois que ela é criada.

Um array estático possui tamanho definido na sua criação e esse tamanho não pode ser alterado, suas características são: tamanho fixo, memória alocada para aquele tamanho, acesso por indice (O(1)), não cresce automaticamente e não diminui automaticamente.

Exemplo:

```
int[] numbers = new int[5]

numbers.length //5

// você pode alterar os valores, mas não pode adicionar
```

Um array dinâmico é uma estrutura que utiliza um Array internamente, mas consegue aumentar sua capacidade  quando necessário. Ou seja mesmo que sua capacidade inicial não seja suficiente, a estrutura pode criar um Array maior e transferir os elementos. O mecanismo interno faz algo parecido com: 

1. Aloca um array maior
2. Copia os elementos antigos
3. Coloca um elemento novo
4. Descarta o array antigo

Exemplo:

```
ArrayList<Integer> numbers = new ArrayList<>();

//podemos fazer
numbers.add(10);
numbers.add(20);
numbers.add(30);
numbers.add(40);
numbers.add(50);
numbers.add(60);
```

Por qual motivo add é O(1)?

Um conceito importante para explicar isso é: Inserir no final de um array dinâmico é O(1) amortizado.

Ou seja na maioria das vezes é uma operação constante, mas ocasionalmente, o array fica cheio e para adicionar um novo elemento é preciso redimensionar e copiar n elementos custa O(n), então temos a maioria das inserções sendo O(1) e algumas sendo O(n). Ao analisar uma sequência grande de inserções, o custo médio da operação é: O(1) amortizado.

Outro conceito importante é capcidade vs tamanho, e isso para um array dinâmico é especialmente importante.

Imagine:

```
[10][20][30][ ][ ][ ]
```

então podemos ter:

size = 3  quantidade de elementos existente
capacity = 6 quantidade de elementos que o espaço atualmente alocado consegue comportar.

se adicionarmos um elemento o size vira 4 e capacity continua 6, mas ao estourar a capacidade o array precisa crescer.

size == capacity 

O array cresce.

Um detalhe importante um array dinâmico não significa que o array muda fisicamente de tamanho. Esse é um detalhe importante, o array continua tendo seu tamanho fixo. Quem é dinâmico é a estrutura que gerencia o array.

conceitualmente:
```
ArrayList
   ↓
Array interno
   ↓
[10][20][30][40][ ][ ]


// Quando fica cheio:

Array interno antigo
[10][20][30][40]

        ↓

Array interno novo
[10][20][30][40][ ][ ][ ][ ]
```
A estrutura substitui o array interno por outro maior.

por isso a definição mais precisa seria:

Array estático: array cujo tamanho é definido na criação e não pode ser alterado

Array dinâmico: estrutura que usa um array interno de tamanho fixo, mas pode realocá-lo para uma área maior quando necessário, oferecendo ao usuário a aparência de um array que cresce dinamicamente.

Array estático possui tamanho fixo após sua criação. Array dinâmico utiliza um array interno e, quando sua capacidade é atingida, realoca uma área maior e copia os elementos. Isso permite crescimento automático. O acesso por índice continua sendo O(1), enquanto a inserção no final de um array dinâmico é O(1) amortizado.

#Conceitos
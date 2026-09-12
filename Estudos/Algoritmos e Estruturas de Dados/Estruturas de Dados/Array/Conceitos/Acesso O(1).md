
O(1) significa tempo constante.

Isso que dizer que a quantidade de operações necessárias para realizar determinada operação não depende do tamanho do array.

Ou seja imagine que eu tenha milhões de elementos, se eu pedir o ```array[2]``` o computador não precisa percorrer os elementos, ele vai calcula diretamente a posição do índice 2.

Portanto:
```
n = 5
array[2] → O(1)

n = 1.000
array[2] → O(1)

n = 1.000.000
array[2] → O(1)
```
O n não aparece no custo da operação.

O que não é O(1)?

É importante para não confundir acesso com busca

imagine um conjunto de array no intervalo de 10 a 50 e eu quero o índice 3 e o resultado é 40 a complexidade vai ser O(1).

Agora imagine outra situação onde aogra eu pergunto: Onde está o número 40? E o computador precisar percorrer vários elementos  a complexidade vai ser O(n)

Logo essa diferença é fundamental: 

O(1) o computador já calcula e entrega o resultado imediatamente 
O(n) o computador vai percorrer todos os elementos até chegar no resultado

Acesso != Busca

Acesso por índice e atualização por índice são complexidade O(1)
Busca por valor e percorrer tudo é complexidade O(n)

e qual a importância disso em algoritmos

Imagine Two Pointers
```
[1, 3, 5, 7, 9, 11]
 ↑                 ↑
left              right

// Você consegue fazer:

array[left]
array[right]

em:

O(1)

então pode mover os ponteiros

left++
right--
```

e continuar acessando diretamente os elementos.

Uma observação importante:

O(1) não significa necessariamente que a operação vai levar 1 segundo ou exatamente uma operação, significa que o custo é limitado por uma constante, independentemente de n.

Por exemplo: 

O(1)

pode envolver algumas intruções internas. O importante é:

n aumenta
   ↓
custo da operação não cresce proporcionalmente

enquanto: O(n)

significa:

n aumenta
   ↓
custo aumenta proporcionalmente

O Acesso O(1) em um array significa que, dado um índice válido, podemos acessar diretamente o elemento sem percorrer os elementos anteriores. O custo do acesso não depende do tamanho do Array.


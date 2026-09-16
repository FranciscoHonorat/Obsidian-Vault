A recursão ocorre quando uma função chama a si mesma para resolver um problema. É uma técnica poderosa que pode simplificar a solução de problemas complexos, dividindo-os em subproblemas menores e mais gerenciáveis. No entanto, é importante ter cuidado ao usar recursão, pois chamadas recursivas excessivas podem levar a estouro de pilha e problemas de desempenho.

Conceitos fundamentais da recursão incluem:

- Caso base: A condição que encerra a recursão, evitando chamadas infinitas. É essencial definir um caso base claro para garantir que a função termine corretamente.
- Chamada recursiva: A invocação da própria função dentro de seu corpo, geralmente com argumentos modificados para aproximar-se do caso base.
- Pilha de chamadas: Cada chamada recursiva é adicionada à pilha de chamadas, e a função só retorna quando o caso base é atingido. Isso significa que a memória usada para armazenar as chamadas recursivas pode crescer rapidamente, especialmente em casos de recursão profunda.
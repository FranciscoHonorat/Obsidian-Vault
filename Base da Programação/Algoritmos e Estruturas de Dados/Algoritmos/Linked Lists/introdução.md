Os algoritmos para listas encadeadas (simples, duplas ou circulares) fundamentam-se no manejo rigoroso de ponteiros para alterar estruturas de dados sem a necessidade de realocação contínua de memória.

**Operações de Base**

* **Inserção e Remoção:** Realizadas em $O(1)$ no início ou $O(n)$ em posições arbitrárias (sendo $O(1)$ no fim caso exista ponteiro de cauda).
* **Busca Sequencial:** Acesso linear $O(n)$, já que a estrutura não permite acesso aleatório por índice.

**Técnicas de Dois Ponteiros (Fast & Slow)**

* **Detecção de Ciclos (Algoritmo de Floyd):** Identifica loops usando um ponteiro lento (1 passo) e um rápido (2 passos); o encontro deles confirma a existência do ciclo.
* **Encontrar o Meio da Lista:** Quando o ponteiro rápido atinge o fim, o ponteiro lento estará exatamente no nó central.
* **$k$-ésimo Elemento a partir do Fim:** Mantém-se um intervalo de $k$ nós entre dois ponteiros antes de movê-los simultaneamente.
* **Ponto de Interseção:** Localiza o nó onde duas listas se fundem ajustando o ponto de partida com base na diferença de comprimento entre elas.

**Reordenação e Manipulação**

* **Inversão (Iterativa ou Recursiva):** Inverte o sentido de cada ponteiro `next` em tempo $O(n)$ e espaço $O(1)$.
* **Merge Sort para Listas:** Algoritmo de ordenação ideal para esta estrutura ($O(n \log n)$), dividindo a lista ao meio e fundindo os nós ordenadamente sem custo de alocação de memória extra.
* **Verificação de Palíndromo:** Localiza o meio da lista, inverte a segunda metade e compara valor a valor com a primeira metade.
* **Particionamento por Pivô:** Reorganiza os nós criando sublistas para elementos menores e maiores que um determinado valor $x$.

**Algoritmos Avançados e Aplicações**

* **Cache LRU (Least Recently Used):** Combina uma Lista Duplamente Encadeada com uma Tabela Hash para realizar acessos e atualizações de prioridade em $O(1)$.
* **Fusão de $K$ Listas Ordenadas:** Utiliza uma fila de prioridade (Min-Heap) para intercalar $K$ listas ordenadas em $O(N \log K)$.
* **Soma de Números Representados por Listas:** Simula a adição aritmética nó por nó (dígito por dígito), manipulando o transporte (*carry*).
* **Achatar Lista Multinível:** Transforma uma estrutura hierárquica (com nós apontando para listas filhas) em uma única lista unidimensional via percurso em profundidade (DFS).

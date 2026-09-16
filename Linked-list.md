# Linked Lists

Uma lista encadeada (linked lists) é uma estrutura de dados linear composta por nós (nodes) espalhados na memória. Ou seja em vez de utilizar posições contiguas em um array, cada nó armazena o seu proprio valor é um ponteiro para o nó seguinte.

# Estrutura básica:

- Head: O primeiro nó da lista. Se for perdido, toda a lista se torna inacessível.
- Tail: O último nó da lista, cujo ponteiro aponta para null (ou none) indicando o término.
- Node: A unidade fundamental contendo [Dado | Ponteiro]

# Tipos principais:

- Simplesmente Encadeada (Singly Linked List): Navegação em sentido único. Cada nó aponta apenas para o próximo.
- Duplamente Encadeada (Doubly Linked List): Navegação bidirecional. Cada nó armazena referências para o nó seguinte e para o anterior.
- Circulo: O ponteiro do último nó apona de volta para o Head, formando um ciclo.

# Complexidade de Tempo (Big-O)

Operação        | Início (HEAD) | Fim (Tail)                |                 Posição Arbitrária
inserção            O(1)            O(1)*                                           O(n)
remoção             O(1)            O(1)(Duplo)/O(n)(Simples)                       O(n)
acessa/busca        O(1)            O(n)                                            O(n)

# Prós e Contras:

## Vantagens:
    1. Tamanho Dinâmico
    2. Inserções/Remoção no topo rápidas

## Desvantagens:
    1. Sem acesso aleatório
    2. Sobrecarga de memória
    3. Performance de cache

# Quando Utilizar

Ideal para cenários com inserções e remoções constantes no início/fim, como na implementação interna de Pilhas (Stacks), Filas (Queues), sistemas de desfazer/refazer (Undo/Redo) e histórico de páginas do navegador (Avançar/Voltar).
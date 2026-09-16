Uma fila (Queue) é uma estrutura de dados linear que segue o princípio FIFO (First In, First Out), ou seja, o primeiro elemento a ser inserido é o primeiro a ser removido. Ela é utilizada em diversas aplicações, como gerenciamento de tarefas, impressão de documentos, e simulação de processos.

A estrutura  ideal para processamento de tarefas em fila de espera, buffers de mensagens e algoritmos de busca em largura (BFS) em grafos. Funciona da seguinte maneira: elementos são adicionados no final da fila (enqueue) e removidos do início da fila (dequeue).

As principais operações de uma fila são:
1. **Enqueue**: Adiciona um elemento ao final da fila. O(1) - tempo constante.
2. **Dequeue**: Remove o elemento do início da fila. O(1)
3. **Peek/Front**: Retorna o elemento do início da fila sem removê-lo. O(1)
4. **IsEmpty**: Verifica se a fila está vazia. O(1)

Como funciona em Go?

Go não possui um tipo Queue nativo em sua biblioteca padrão. As duas formas idiomáticas de criar uma fila na linguagem dependem do contexto de execução: Slices ou Channels. A escolha entre essas duas abordagens depende do caso de uso específico, como concorrência ou simplicidade de implementação.

1. Implementação Tradicional com Slices e Generics:
```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    item := q.items[0]
    q.items[0] = zero // Clear the reference for garbage collection
    q.items = q.items[1:]
    return item, true 
}
```


2. Implementação Conrrente com Channels:
```go
type Queue[T any] struct {
    ch chan T
}

func NewQueue[T any](size int) *Queue[T] {
    return &Queue[T]{ch: make(chan T, size)}
}

//Fila concorrente com capacidade para 3 itens
func (q *Queue[T]) Enqueue(item T) {
    q.ch <- item
}

func (q *Queue[T]) Dequeue() (T, bool) {
    select {
    case item := <-q.ch:
        return item, true
    default:
        var zero T
        return zero, false
    }
}
```

Abordagem em slice:

Vantagens: Tamanho dinâmico, baixo custo de aloção inicial.
Desvantagens: Requer manipulação manual de ponteiros para evitar retenção de memória no GC.
Uso Ideal: Algoritmos de grafo (BFS), utilitários locais e códigos sequneciais.

Abordagem com channels:

Vantagens: Suporte nativo à concorrência, thread-safe , fácil de usar em goroutines sem mutexes.
Desvantagens: Requer definição de capacidade fixa no buffer e bloqueia quando a fila está cheia ou vazia.
Uso Ideal: Worker pools, processamento de eventos assincronos e pipelines de dados concorrentes.


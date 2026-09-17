As Operações de base de uma lista encadeada incluem inserção, remoção e busca de elementos. A complexidade dessas operações depende da posição do elemento na lista e do tipo de lista utilizada (simplesmente encadeada, duplamente encadeada ou circular).

Resumo das Operações de Base:

| Operação | Complexidade de Tempo (Big-O) |
|----------|-------------------------------|
| Inserção no Início (Head) | O(1) |
| Inserção no Fim (Tail) | O(1)* |
| Inserção em Posição Arbitrária | O(n) |
| Remoção no Início (Head) | O(1) |
| Remoção no Fim (Tail) | O(1) (Duplo) / O(n) (Simples) |
| Remoção em Posição Arbitrária | O(n) |
|------------------------------------------|


1. Exemplo em Go: SIstema de Histórico de Ações (Undo History)

Em um sistema real de auditoria ou histórico de ações, como em editores de texto ou navegadores, eventos recentes são adicionados no topo da lista (Head) em O(1), buscando e removendo eventos antigos pode ser feito em O(n) se necessário. A lista encadeada permite que o histórico cresça dinamicamente sem a necessidade de realocação de memória, como seria o caso com arrays.
```go
package main

import "fmt"

type Node struct {
    Action string
    Next  *Node
}

type HistoryList struct {
    Head *Node
}

func (h *HistoryList) Push(action string) {
    newNode := &Node{Action: action, Next: h.Head}
    h.Head = newNode
}

func (h *HistoryList) Contains(action string) bool {
    current := h.Head
    for current != nil {
        if current.Action == action {
            return true
        }
        current = current.Next
    }
    return false
}

func (h *HistoryList) Delete(action string) bool {
    if h.Head == nil {
        return false
    }

    if h.Head.Action == action {
        h.Head = h.Head.Next
        return true
    }

    current := h.Head
    for current.Next != nil {
        if current.Next.Action == action {
            current.Next = current.Next.Next
            return true
        }
        current = current.Next
    }
    return false
}

func (h *HistoryList) Print() {
    current := h.Head
    for current != nil {
        fmt.Println(current.Action)
        current = current.Next
    }
}

func main() {
    history := &HistoryList{}

    history.Push("Open File")
    history.Push("Edit Text")
    history.Push("Save File")

    fmt.Println("Current History:")
    history.Print()

    fmt.Println("\nContains 'Edit Text'? ", history.Contains("Edit Text"))

    history.Delete("Edit Text")
    fmt.Println("\nHistory after deleting 'Edit Text':")
    history.Print()
}
```
2. Exemplo: Em aplicativos de streaming de música, listas encadeadas garatem a gestão de filas de reprodução, permitindo adicionar faixas ao final em tempo constante e remover músicas canceladas pelo usuário sem reindexar o restante dos elementos, mantendo a performance e a integridade da fila de reprodução.

Fila de Reprodução de Músicas

-  Entrada na fila (Enqueue): Inserção no final em O(1) usando um ponteiro direto para o último nó (Tail).
-  Tocar próxima (Dequeue): Remoção do início da fila em O(1) atualizando o ponteiro principal(Head).
-  Cancelar Música (Remove): Busca e desconexão de uma faixa específica no meio da fila em O(n), mantendo a integridade da lista encadeada.

```go
package main

import "fmt"

type Track struct {
    Title string
    Artist string
    Next  *Track
}

type PlaylistQueue struct {
    Head *Track
    Tail *Track
}

func (p *PlaylistQueue) Enqueue(title, artist string) {
    newTrack := &Track{Title: title, Artist: artist}

    if p.Head == nil {
        p.Head = newTrack
        p.Tail = newTrack
        return
    }

    p.Tail.Next = newTrack
    p.Tail = newTrack
}

func (p *PlaylistQueue) PlayNext() *Track {
    if p.Head == nil {
        return nil
    }

    currentTrack := p.Head
    p.Head = p.Head.Next

    if p.Head == nil {
        p.Tail = nil
    }

    currentTrack.Next = nil // Clear the reference for garbage collection
    return currentTrack
}

func (p *PlaylistQueue) CancelTrack(title string) bool {
    if p.Head == nil {
        return false
    }

    if p.Head.Title == title {
        p.PlayNext()
        return true
    }

    current := p.Head
    for current.Next != nil {
        if current.Next.Title == title {
            if current.Next == p.Tail {
                p.Tail = current
            }
            current.Next = current.Next.Next
            return true
        }
        current = current.Next
    }
    return false
}

func (p *PlaylistQueue) Display() {
    if p.Head == nil {
        fmt.Println("Playlist is empty.")
        return
    }

    current := p.Head
    for current != nil {
        fmt.Printf("Title: %s, Artist: %s\n", current.Title, current.Artist)
        current = current.Next
    }
    fmt.Println("Fim")
}

func main() {
    playlist := &PlaylistQueue{}

    playlist.Enqueue("Song A", "Artist 1")
    playlist.Enqueue("Song B", "Artist 2")
    playlist.Enqueue("Song C", "Artist 3")

    fmt.Println("Current Playlist:")
    playlist.Display()

    fmt.Println("\nPlaying next track:")
    nextTrack := playlist.PlayNext()
    if nextTrack != nil {
        fmt.Printf("Now playing: %s by %s\n", nextTrack.Title, nextTrack.Artist)
    }

    fmt.Println("\nPlaylist after playing next track:")
    playlist.Display()

    fmt.Println("\nCanceling 'Song B':")
    if playlist.CancelTrack("Song B") {
        fmt.Println("'Song B' has been canceled.")
    } else {
        fmt.Println("'Song B' not found in the playlist.")
    }

    fmt.Println("\nPlaylist after canceling 'Song B':")
    playlist.Display()
}
```

3. Exemplo: Em sistemas de mensagens instantâneas, listas encadeadas podem ser usadas para gerenciar a fila de mensagens recebidas, permitindo que novas mensagens sejam adicionadas rapidamente e mensagens antigas sejam removidas conforme necessário.

Fila de Mensagens Instantâneas: Em filas de mensagens instantâneas, cada mensagem recebida é adicionada ao final da lista encadeada (Enqueue) em O(1), enquanto mensagens antigas podem ser removidas do início da lista (Dequeue) também em O(1). A busca por mensagens específicas pode ser feita em O(n) caso seja necessário localizar e remover uma mensagem específica.

```go
package main

import "fmt"

type Message struct {
	ID     string
	Sender string
	Text   string
	Next   *Message
}

type MessageQueue struct {
	Head *Message
	Tail *Message
}

// Enqueue: Adiciona nova mensagem recebida ao final da fila em O(1)
func (q *MessageQueue) Enqueue(id, sender, text string) {
	newMessage := &Message{ID: id, Sender: sender, Text: text}

	if q.Head == nil {
		q.Head = newMessage
		q.Tail = newMessage
		return
	}

	q.Tail.Next = newMessage
	q.Tail = newMessage
}

// Dequeue: Remove e retorna a mensagem mais antiga para exibição/entrega em O(1)
func (q *MessageQueue) Dequeue() *Message {
	if q.Head == nil {
		return nil
	}

	msg := q.Head
	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	msg.Next = nil
	return msg
}

// Localiza e remove uma mensagem revogada por ID em O(n)
func (q *MessageQueue) DeleteByID(id string) bool {
	if q.Head == nil {
		return false
	}

	// Caso o item a apagar seja a mensagem no topo da fila
	if q.Head.ID == id {
		q.Dequeue()
		return true
	}

	current := q.Head
	for current.Next != nil {
		if current.Next.ID == id {
			// Ajusta a cauda caso a mensagem a ser removida seja a última
			if current.Next == q.Tail {
				q.Tail = current
			}
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

func (q *MessageQueue) Display() {
	if q.Head == nil {
		fmt.Println("Nenhuma mensagem pendente.")
		return
	}

	current := q.Head
	for current != nil {
		fmt.Printf("[%s | %s: \"%s\"] -> ", current.ID, current.Sender, current.Text)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	mq := &MessageQueue{}

	// Chegada de novas mensagens (Enqueue O(1))
	mq.Enqueue("msg_101", "Alice", "Oi, você viu o relatório?")
	mq.Enqueue("msg_102", "Bob", "Vou olhar agora.")
	mq.Enqueue("msg_103", "Alice", "Enviei o arquivo errado, desconsidere!")

	fmt.Println("Fila de mensagens recebidas:")
	mq.Display()

	// Remetente apaga a mensagem 'msg_103' antes de ser entregue (Exclusão O(n))
	mq.DeleteByID("msg_103")
	fmt.Println("\nApós a mensagem msg_103 ser apagada pelo remetente:")
	mq.Display()

	// Processamento e exibição da mensagem na tela (Dequeue O(1))
	processedMsg := mq.Dequeue()
	fmt.Printf("\nMensagem entregue na tela: %s - \"%s\"\n", processedMsg.Sender, processedMsg.Text)

	fmt.Println("\nFila de mensagens restante:")
	mq.Display()
}
```
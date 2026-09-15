# Prática Intensiva: Grafos em Go

## Como este programa funciona

Sete blocos progressivos. Cada bloco tem tarefas com **deliverable** (o que você escreve), **critério de aceite** (como você prova que está certo, sozinho) e **pergunta de fechamento** (o que você precisa saber explicar).

Regras do ciclo:

1. Você implementa. Eu não dou a solução antes da sua tentativa.
2. Você valida contra o oráculo antes de submeter. Código que não passa no oráculo não vai para revisão.
3. Eu reviso: correção, complexidade, Go idiomático, e faço perguntas sobre decisões.
4. Bloco só é concluído quando você responde a pergunta de fechamento sem consultar material.

Não pule blocos. Cada um depende da estrutura do anterior.

---

## Estrutura do repositório

```
grafos-lab/
  go.mod                    // module grafos-lab
  internal/
    graph/                  // BLOCO 0: CSR, builder, interner
    traverse/               // BLOCO 1: bfs, dfs, classify, toposort
    connect/                // BLOCO 2: cc, scc, bridges, articulation
    paths/                  // BLOCO 3: dijkstra, bellmanford, mst, semiring
    implicit/               // BLOCO 4: busca em espaço de estados
    flow/                   // BLOCO 5: dinic, split, matching
    perf/                   // BLOCO 6: benchmarks, renumeração, triângulos
    gen/                    // gerador de grafos aleatórios (fornecido)
    oracle/                 // implementações de referência lentas
```

---

## A metodologia do oráculo

Esta é a parte que torna a prática *intensiva*: você não depende de mim para saber se acertou.

O padrão é sempre o mesmo. Você escreve o algoritmo rápido e esperto. Você também escreve (ou recebe) uma implementação **burra, lenta e obviamente correta**. Então roda as duas em centenas de grafos aleatórios pequenos e compara.

Um Dijkstra com bug sutil em decrease-key passa em 10 testes feitos à mão. Não passa em 500 grafos aleatórios comparados contra Floyd-Warshall.

Regras para o oráculo ser útil:

- Grafos **pequenos** (n entre 1 e 12). O bug aparece igual e o oráculo O(n³) roda instantâneo.
- Sementes **determinísticas** (`rand.NewSource(int64(seed))`). Falhou na seed 137? Você reproduz.
- Inclua os degenerados: n=1, grafo vazio, grafo completo, grafo desconexo, self-loop.
- Quando falhar, **reduza**: diminua n até o menor grafo que ainda falha, e imprima.

---

## Scaffold fornecido

Copie estes dois pacotes. São infraestrutura, não conteúdo de aprendizado.

### `internal/gen/gen.go`

```go
// Package gen produz grafos aleatórios reprodutíveis para testes.
package gen

import "math/rand"

// Directed gera um dígrafo G(n, p): cada par ordenado (u,v), u != v,
// recebe aresta com probabilidade p. Sem self-loops, sem paralelas.
func Directed(r *rand.Rand, n int, p float64) [][]int32 {
	adj := make([][]int32, n)
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if u != v && r.Float64() < p {
				adj[u] = append(adj[u], int32(v))
			}
		}
	}
	return adj
}

// Undirected gera um grafo simples não-dirigido G(n, p).
// Cada aresta aparece nas listas das duas pontas.
func Undirected(r *rand.Rand, n int, p float64) [][]int32 {
	adj := make([][]int32, n)
	for u := 0; u < n; u++ {
		for v := u + 1; v < n; v++ {
			if r.Float64() < p {
				adj[u] = append(adj[u], int32(v))
				adj[v] = append(adj[v], int32(u))
			}
		}
	}
	return adj
}

// DAG gera um grafo dirigido acíclico orientando toda aresta do
// índice menor para o maior. Útil para testar ordenação topológica.
func DAG(r *rand.Rand, n int, p float64) [][]int32 {
	adj := make([][]int32, n)
	for u := 0; u < n; u++ {
		for v := u + 1; v < n; v++ {
			if r.Float64() < p {
				adj[u] = append(adj[u], int32(v))
			}
		}
	}
	return adj
}

// Tree gera uma árvore aleatória com n vértices conectando cada
// vértice v > 0 a um pai uniformemente sorteado em [0, v).
func Tree(r *rand.Rand, n int) [][]int32 {
	adj := make([][]int32, n)
	for v := 1; v < n; v++ {
		p := int32(r.Intn(v))
		adj[v] = append(adj[v], p)
		adj[p] = append(adj[p], int32(v))
	}
	return adj
}

// WeightEdge é uma aresta com peso, para os geradores ponderados.
type WeightEdge struct {
	To     int32
	Weight int64
}

// DirectedWeighted gera um dígrafo com pesos em [lo, hi].
// Use lo negativo para exercitar Bellman-Ford.
func DirectedWeighted(r *rand.Rand, n int, p float64, lo, hi int64) [][]WeightEdge {
	adj := make([][]WeightEdge, n)
	span := hi - lo + 1
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if u != v && r.Float64() < p {
				w := lo + r.Int63n(span)
				adj[u] = append(adj[u], WeightEdge{To: int32(v), Weight: w})
			}
		}
	}
	return adj
}
```

### `internal/oracle/oracle.go`

```go
// Package oracle contém implementações de referência: lentas,
// diretas e obviamente corretas. Servem para validar as rápidas.
package oracle

import "grafos-lab/internal/gen"

// Inf representa ausência de caminho. Escolhido grande o bastante
// para comparações e pequeno o bastante para não estourar na soma.
const Inf int64 = 1 << 40

// FloydUnweighted devolve dist[u][v] em número de arestas.
// Referência para BFS. Custo O(n^3).
func FloydUnweighted(adj [][]int32) [][]int64 {
	n := len(adj)
	d := make([][]int64, n)
	for i := range d {
		d[i] = make([]int64, n)
		for j := range d[i] {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = Inf
			}
		}
	}
	for u := range adj {
		for _, v := range adj[u] {
			if 1 < d[u][v] {
				d[u][v] = 1
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return d
}

// FloydWeighted é a referência para Dijkstra e Bellman-Ford.
// Só é válida se não houver ciclo negativo.
func FloydWeighted(adj [][]gen.WeightEdge) [][]int64 {
	n := len(adj)
	d := make([][]int64, n)
	for i := range d {
		d[i] = make([]int64, n)
		for j := range d[i] {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = Inf
			}
		}
	}
	for u := range adj {
		for _, e := range adj[u] {
			if e.Weight < d[u][e.To] {
				d[u][e.To] = e.Weight
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if d[i][k] >= Inf {
				continue
			}
			for j := 0; j < n; j++ {
				if d[k][j] >= Inf {
					continue
				}
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return d
}

// HasNegativeCycle detecta ciclo negativo alcançável, olhando a
// diagonal depois de Floyd. Referência para Bellman-Ford.
func HasNegativeCycle(adj [][]gen.WeightEdge) bool {
	d := FloydWeighted(adj)
	for i := range d {
		if d[i][i] < 0 {
			return true
		}
	}
	return false
}
```

### Padrão de teste

```go
func TestBFSContraOraculo(t *testing.T) {
	for seed := int64(0); seed < 500; seed++ {
		r := rand.New(rand.NewSource(seed))
		n := 1 + r.Intn(12)
		p := r.Float64()
		adj := gen.Directed(r, n, p)

		want := oracle.FloydUnweighted(adj)

		for s := 0; s < n; s++ {
			got := traverse.BFS(adj, int32(s))
			for v := 0; v < n; v++ {
				exp := want[s][v]
				if exp >= oracle.Inf {
					exp = -1 // convenção do seu BFS para inalcançável
				}
				if int64(got[v]) != exp {
					t.Fatalf("seed=%d n=%d src=%d v=%d: got %d, want %d",
						seed, n, s, v, got[v], exp)
				}
			}
		}
	}
}
```

Quando quebrar, o `t.Fatalf` te dá a seed. Reduza `n` até o mínimo que falha e imprima o grafo.

---

# BLOCO 0 — Fundação

Objetivo: parar de usar `map[int][]int` para sempre.

### T0.1 — Interner

**Deliverable:** `graph.Vertices` que traduz `string` para `int32` denso e de volta.

**Critério de aceite:** IDs consecutivos a partir de 0, idempotente (mesmo nome devolve mesmo ID), `Name(ID(s)) == s` para 10 mil strings aleatórias.

### T0.2 — Builder e CSR

**Deliverable:**

```go
type Builder struct { /* acumula arestas */ }
func (b *Builder) AddEdge(u, v int32)
func (b *Builder) Build(n int) *CSR

type CSR struct {
    offset []int32 // len n+1
    target []int32 // len m
}
func (g *CSR) Neighbors(v int32) []int32
func (g *CSR) NumVertices() int
func (g *CSR) NumEdges() int
```

**Restrição obrigatória:** o `Build` **não pode usar `sort.Slice`**. Use counting sort pela origem, em duas passadas: conte graus, faça prefix sum para gerar `offset`, preencha `target`. Custo O(V+E), não O(E log E).

**Critério de aceite:** contra 500 grafos aleatórios, o conjunto de vizinhos devolvido pelo CSR deve ser igual ao do `[][]int32` original, para todo vértice. Inclua n=0, n=1, grafo sem arestas.

**Armadilha:** o prefix sum é destrutivo se você não tomar cuidado. Você vai precisar de um cursor auxiliar ou restaurar o `offset` depois. Descubra por quê na prática.

### T0.3 — Estatísticas de grau

**Deliverable:** função que devolve mínimo, máximo, média, mediana e p99 do grau.

**Exercício de percepção:** gere `Undirected(r, 100000, 0.0001)` e imprima as estatísticas. Depois construa um grafo por anexação preferencial (cada vértice novo liga a um existente escolhido com probabilidade proporcional ao grau) e imprima de novo. Compare max/mediana nos dois.

**Pergunta de fechamento do bloco:**
Por que `offset` tem tamanho `V+1` e não `V`? O que quebra se for `V`?

---

# BLOCO 1 — Travessia

### T1.1 — BFS

**Deliverable:** `BFS(g, src) (dist, parent []int32)` e `Path(parent, dst) []int32`.

**Restrições:** fila como slice com índice `head`, sem `queue = queue[1:]`. `dist` inicializado com -1 para inalcançável.

**Critério de aceite:** contra `oracle.FloydUnweighted`, 500 grafos.

### T1.2 — DFS iterativa

**Deliverable:** DFS **sem recursão**, produzindo ordem de pré-ordem e pós-ordem.

**Por que importa:** a goroutine principal do Go cresce a pilha dinamicamente, mas uma DFS recursiva em grafo de milhões de vértices ainda é um risco real de estouro e é mais lenta. Saber converter recursão em pilha explícita é habilidade base.

**Armadilha central:** a pós-ordem exige que você saiba *quando voltou* de um filho. A pilha precisa guardar o estado da iteração sobre os vizinhos, não só o vértice. Pense em guardar `(vértice, índice do próximo vizinho)`.

**Critério de aceite:** a pré-ordem e pós-ordem devem bater exatamente com uma versão recursiva que você também escreve, em 500 grafos, percorrendo vizinhos na mesma ordem.

### T1.3 — Classificação de arestas

**Deliverable:** classifique cada aresta como tree, back, forward ou cross durante a DFS.

**Critério de aceite:** três invariantes verificadas automaticamente em grafos aleatórios:
1. Toda aresta recebe exatamente uma classificação.
2. O número de tree edges é igual a `V` menos o número de árvores da floresta DFS.
3. **Em grafo não-dirigido, forward e cross nunca aparecem.** Rode em 500 grafos não-dirigidos e falhe o teste se aparecerem.

A invariante 3 é a mais valiosa: ela é um teorema, e seu código verificando-a é a prova empírica de que você o entendeu.

### T1.4 — Detecção de ciclo com três cores

**Deliverable:** `HasCycle(g) bool` para dígrafo, mais `FindCycle(g) []int32` devolvendo um ciclo concreto.

**Restrição:** três estados (branco, cinza, preto), não um booleano.

**Critério de aceite:** em `gen.DAG` deve sempre devolver false. Em `gen.Directed` compare com um oráculo: existe ciclo se e somente se `FloydUnweighted(adj)[i][i] < Inf` para algum `i` usando alcançabilidade em 1+ passos. E quando devolver um ciclo, **valide que ele é de fato um ciclo**: consecutivos ligados por aresta, último liga no primeiro.

**Teste que separa quem entendeu:** construa `0→1, 0→2, 1→2`. Um detector ingênuo baseado em "já visitei" acusa ciclo. O de três cores não. Coloque isso como caso fixo de teste.

### T1.5 — Ordenação topológica

**Deliverable:** as duas versões, Kahn (in-degree + fila) e DFS (pós-ordem invertida). Ambas devem sinalizar ciclo.

**Critério de aceite:** validador genérico, não comparação de listas. A saída é válida se for permutação de `[0,n)` e, para toda aresta `u→v`, `pos[u] < pos[v]`. Existem muitas ordens válidas, então comparar contra uma resposta fixa é errado.

**Pergunta de fechamento do bloco:**
Por que a detecção de ciclo em grafo **não-dirigido** não precisa de três cores? Ligue sua resposta à invariante 3 do T1.3.

---

# BLOCO 2 — Conectividade

### T2.1 — Componentes conexas, duas vias

**Deliverable:** por BFS e por DSU (union-find com union by rank e path compression).

**Critério de aceite:** as duas devem produzir a mesma partição (rótulos podem diferir, a partição não). Escreva um comparador de partições.

### T2.2 — SCC de Tarjan, iterativa

**Deliverable:** Tarjan em uma passada, **sem recursão**.

Esta é a tarefa mais difícil do bloco. Converter Tarjan para iterativo exige manter na pilha o estado de iteração e o ponto de retorno, porque há trabalho a fazer *depois* da chamada recursiva (o `low[v] = min(low[v], low[u])`).

**Critério de aceite:** oráculo por força bruta. `u` e `v` estão na mesma SCC se e somente se `u` alcança `v` e `v` alcança `u`. Compute alcançabilidade com `FloydUnweighted` e compare as partições. 500 grafos.

### T2.3 — Condensação

**Deliverable:** construir o grafo quociente das SCCs, sem arestas paralelas.

**Critério de aceite:** rode seu `HasCycle` do T1.4 sobre a condensação. **Deve ser sempre falso.** Se algum grafo aleatório produzir ciclo, ou sua SCC está errada ou sua condensação está. Esse teste sozinho pega quase todo bug de SCC.

### T2.4 — Pontes e pontos de articulação

**Deliverable:** as duas em uma DFS com `disc`/`low`.

**Critério de aceite:** força bruta. Uma aresta é ponte se removê-la aumenta o número de componentes. Um vértice é ponto de corte se removê-lo (e suas arestas) aumenta o número de componentes no grafo restante. Compare, 500 grafos.

**Caso fixo obrigatório:** dois triângulos compartilhando um vértice. Esperado: zero pontes, um ponto de articulação. Se seu código acusar ponte, a condição `>` virou `>=` em algum lugar.

**Pergunta de fechamento do bloco:**
Por que a condição de ponte usa `low[u] > disc[v]` e a de ponto de articulação usa `>=`? E por que a raiz da DFS precisa de tratamento separado?

---

# BLOCO 3 — Caminhos e pesos

### T3.1 — Dijkstra

**Deliverable:** Dijkstra com `container/heap` e **lazy deletion**.

Go não tem decrease-key. O padrão correto é empurrar a entrada duplicada e descartar na hora de consumir, comparando a distância registrada com a distância corrente. Mesmo padrão do bucket queue do T6.4.

**Critério de aceite:** `FloydWeighted` com pesos não-negativos, 500 grafos.

**Experimento obrigatório:** rode seu Dijkstra em um grafo com uma aresta de peso negativo e **mostre onde ele erra**. Não basta saber que falha, você precisa ver o caminho errado que ele produz e explicar em que momento a decisão gulosa fechou cedo demais.

### T3.2 — Bellman-Ford

**Deliverable:** Bellman-Ford com detecção de ciclo negativo e **extração do ciclo**.

**Critério de aceite:** contra `oracle.HasNegativeCycle` em grafos com pesos em [-5, 10]. Quando detectar, valide que o ciclo extraído tem soma de pesos negativa de fato.

**Detalhe que quase todos erram:** o vértice onde a n-ésima iteração relaxou não está necessariamente *no* ciclo. Ele está alcançável a partir do ciclo. Você precisa andar `n` passos pelos pais antes de começar a coletar.

### T3.3 — Generalização por semianel

**Deliverable:** refatore Dijkstra para receber as operações como parâmetro, e instancie três problemas com **o mesmo corpo de algoritmo**:

| Instância | combinar | estender | neutro |
|---|---|---|---|
| Caminho mínimo | `min` | `+` | `+Inf` / `0` |
| Caminho mais largo | `max` | `min` | `-Inf` / `+Inf` |
| Alcançabilidade | `OR` | `AND` | `false` / `true` |

**Critério de aceite:** caminho mínimo continua batendo com Floyd. Caminho mais largo bate com uma variante de Floyd onde você troca as operações (escreva o oráculo trocando duas linhas, isso é metade da lição).

**Esta é a tarefa de maior retorno conceitual do programa.** Quando você vir o mesmo código resolver três problemas diferentes só trocando duas funções, a categoria "algoritmos de caminho" colapsa em uma coisa só na sua cabeça.

### T3.4 — Árvore geradora mínima

**Deliverable:** Kruskal (com o DSU do T2.1) e Prim (com o heap do T3.1).

**Critério de aceite:** os dois devem produzir o **mesmo peso total**. As árvores podem diferir quando há pesos empatados, o peso não pode. Valide também que a saída é árvore geradora: `n-1` arestas, acíclica, conexa.

**Pergunta de fechamento do bloco:**
Dijkstra e Prim têm estrutura quase idêntica. Qual é exatamente a linha que difere, e por que essa diferença muda o problema resolvido?

---

# BLOCO 4 — Grafos implícitos

Aqui não existe `adj`. O grafo nasce durante a busca.

### T4.1 — 8-puzzle

**Deliverable:** `State [9]byte`, função `neighbors`, BFS até o objetivo.

**Por que array e não slice:** array é comparável e serve de chave de mapa em Go. Slice não. Isso não é detalhe, é o que torna `map[State]int` possível.

**Critério de aceite:** gere estados embaralhando a partir do objetivo com `k` movimentos aleatórios. A distância encontrada deve ser `<= k` sempre, e `== k` frequentemente. Se alguma vez der `> k`, seu BFS está errado.

**Adicional:** detecte estados insolúveis por paridade de inversões, e confirme que sua busca não encontra solução para eles.

### T4.2 — BFS bidirecional

**Deliverable:** busca a partir das duas pontas, encontrando no meio.

**Critério de aceite:** mesma distância da BFS simples em 200 casos aleatórios, **e** conte os estados expandidos nas duas versões. Reporte a razão. Se não houver redução substancial, sua condição de encontro está errada.

**Armadilha:** o ponto de encontro não devolve a resposta correta ingenuamente. Pense em qual momento exato você pode parar sem perder um caminho mais curto.

### T4.3 — Estado enriquecido

**Deliverable:** labirinto com portas e chaves. Você só passa por uma porta se já pegou a chave correspondente.

**Deliverable central:** o vértice deixa de ser `(linha, coluna)` e passa a ser `(linha, coluna, bitmask de chaves)`.

**Critério de aceite:** construa um labirinto onde a solução exige pegar uma chave e **voltar por onde já passou**. Uma busca que usa `(linha, coluna)` como estado falha nesse caso. Confirme que a sua resolve.

Esta tarefa existe para uma lição só: **o que conta como vértice é decisão sua, e enriquecer o estado é como problemas insolúveis viram solúveis.**

### T4.4 — A*

**Deliverable:** A* com distância Manhattan no 8-puzzle.

**Critério de aceite:** mesma distância do BFS (a heurística é admissível, então o ótimo é preservado), com contagem de nós expandidos muito menor. Reporte a razão em uma tabela de 50 instâncias.

**Experimento:** multiplique a heurística por 2. Ela deixa de ser admissível. Meça quanto mais rápido fica e quantas vezes a resposta deixa de ser ótima. Esse é o trade-off que fundamenta busca heurística inteira.

**Pergunta de fechamento do bloco:**
Por que `map[State]int` é inevitável aqui, mas seria um erro grave no BLOCO 1?

---

# BLOCO 5 — Fluxo

### T5.1 — Aresta com identidade

**Deliverable:** arestas em slice único, adjacência guardando índices, pares residuais em índices `2k` e `2k+1`, gêmea acessada por `i ^ 1`.

**Critério de aceite:** teste unitário de que `addEdge` seguido de mutação por `i` e `i^1` preserva a soma de capacidades do par.

### T5.2 — Dinic

**Deliverable:** Dinic completo, com BFS de níveis e DFS de fluxo bloqueante com ponteiro de iteração por vértice.

**O ponteiro de iteração** (`iter []int32`, retomando de onde parou em vez de reiniciar a varredura de vizinhos) é o que dá a complexidade prometida. Sem ele você tem um Ford-Fulkerson disfarçado.

**Critério de aceite:** em grafos com `n <= 8`, enumere **todos** os `2^n` subconjuntos contendo a origem e não o destino, calcule a capacidade de cada corte, tome o mínimo. O fluxo máximo deve ser exatamente igual. Isto é o teorema max-flow min-cut virando teste automatizado, e é a validação mais bonita do programa inteiro.

### T5.3 — Vertex splitting

**Deliverable:** transformação que converte capacidade por vértice em capacidade por aresta, dividindo `v` em `v_in` e `v_out`.

**Aplicação:** com todas as capacidades de vértice iguais a 1, o fluxo máximo passa a ser o número de caminhos vértice-disjuntos.

**Critério de aceite:** teorema de Menger como teste. O número de caminhos vértice-disjuntos entre `s` e `t` deve ser igual ao menor número de vértices (fora `s` e `t`) cuja remoção desconecta `s` de `t`. Calcule o segundo por força bruta em grafos pequenos, testando todos os subconjuntos.

### T5.4 — Emparelhamento bipartido e König

**Deliverable:** emparelhamento máximo via fluxo, mais extração da cobertura mínima por vértices a partir do corte mínimo.

**Critério de aceite:** teorema de König como teste. `|emparelhamento máximo| == |cobertura mínima por vértices|` em todo grafo bipartido gerado. Valide também que a cobertura extraída de fato cobre todas as arestas.

**Pergunta de fechamento do bloco:**
Por que o truque `i ^ 1` funciona, e o que quebra se você inserir uma aresta sem sua residual?

---

# BLOCO 6 — Performance

Agora você tem algoritmos corretos. Vamos medir.

### T6.1 — CSR contra map

**Deliverable:** benchmark de BFS sobre `map[int32][]int32`, sobre `[][]int32` e sobre CSR.

**Metodologia:** grafo fixo de ~1M vértices e ~10M arestas, `go test -bench`, `-benchmem`. Rode cada um no mínimo 5 vezes e reporte a mediana, não uma medição única.

**Critério de aceite:** um relatório com números e uma explicação. A pergunta a responder não é "qual é mais rápido", é **por quanto e por quê**. Cite alocações, indireção e localidade.

### T6.2 — Renumeração

**Deliverable:** renumere os vértices em ordem de visitação BFS e reconstrua o CSR. Meça a mesma BFS antes e depois.

**Critério de aceite:** o algoritmo é idêntico, o resultado lógico é idêntico, e o tempo muda. Explique por quê, em termos de linhas de cache.

Esta tarefa existe para cravar uma ideia: **a ordem dos vértices é um parâmetro de otimização, não um detalhe arbitrário.**

### T6.3 — Contagem de triângulos

**Deliverable:** contagem orientando cada aresta do vértice de menor grau para o de maior, com `int32` como desempate.

**Critério de aceite:** contra força bruta O(n³) em grafos pequenos. Depois meça em grafo grande contra a versão ingênua e reporte a razão.

**Verificação algébrica adicional:** em um grafo pequeno, compute `traço(A³)/6` com multiplicação de matrizes e confirme que bate. Você acabou de validar um algoritmo combinatório contra álgebra linear.

### T6.4 — Degeneração

**Deliverable:** ordenação por degeneração com bucket queue e lazy deletion, em O(V+E).

**Critério de aceite:** contra uma força bruta que remove repetidamente o vértice de menor grau varrendo tudo, O(V²). Mesma degeneração, e a propriedade verificada: na ordem produzida, todo vértice tem no máximo `d` vizinhos posteriores a ele.

**Pergunta de fechamento do bloco:**
O lazy deletion apareceu no Dijkstra (T3.1) e aqui (T6.4). Qual é o problema comum que ele resolve nos dois casos?

---

# Projeto de consolidação

Depois do BLOCO 6, escolha um. Ambos são material de portfólio.

**Opção A: motor de consulta de grafo.** Ingestão de um dataset real (SNAP, DBLP, Wikipedia links) com interner, construção de CSR, e uma CLI que responde: caminho mais curto entre dois nomes, PageRank top-k, componentes, contagem de triângulos. Com benchmarks documentados e um README explicando as decisões de representação.

**Opção B: resolvedor de espaço de estados genérico.** Uma interface `Problem` com `Start`, `IsGoal`, `Neighbors`, `Heuristic`, e implementações de BFS, BFS bidirecional, Dijkstra, A* e IDA* que funcionam para qualquer `Problem`. Instancie em três domínios (8-puzzle, labirinto com chaves, transformação de palavras) e publique a tabela comparativa de nós expandidos.

---

# Protocolo de revisão

Ao submeter uma tarefa, mande:

1. O código completo do pacote (arquivos inteiros, não trechos)
2. A saída dos testes contra o oráculo
3. Uma frase sobre a decisão de design que você menos tem certeza

Eu revisto correção, complexidade e Go idiomático, aponto o que está errado sem corrigir direto, e pergunto sobre as escolhas. Se travar por mais de uma hora numa tarefa, mande o que tem e onde parou, e eu dou uma pista dimensionada, não a solução.
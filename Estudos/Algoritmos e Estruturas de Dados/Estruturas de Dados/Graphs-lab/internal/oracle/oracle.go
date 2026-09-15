package oracle

import (
	"graphs-lab/internal/gen"
)

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
func FloydWeighted(adj [][]gen.Weight) [][]int64 {
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
func HasNegativeCycle(adj [][]gen.Weight) bool {
	d := FloydWeighted(adj)
	for i := range d {
		if d[i][i] < 0 {
			return true
		}
	}
	return false
}

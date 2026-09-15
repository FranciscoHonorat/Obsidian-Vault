package gen

import (
	"math/rand"
)

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

func Dag(r *rand.Rand, n int, p float64) [][]int32 {
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

func Tree(r *rand.Rand, n int) [][]int32 {
	adj := make([][]int32, n)
	for v := 1; v < n; v++ {
		p := int32(r.Intn(v))
		adj[v] = append(adj[v], p)
		adj[p] = append(adj[p], int32(v))
	}
	return adj
}

type Weight struct {
	To     int32
	Weight int64
}

func DirectedWeighted(r *rand.Rand, n int, p float64, lo, hi int64) [][]Weight {
	adj := make([][]Weight, n)
	span := hi - lo + 1
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if u != v && r.Float64() < p {
				w := lo + r.Int63n(span)
				adj[u] = append(adj[u], Weight{To: int32(v), Weight: w})
			}
		}
	}
	return adj
}

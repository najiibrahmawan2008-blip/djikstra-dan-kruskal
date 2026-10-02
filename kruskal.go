package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	U      int
	V      int
	Weight int
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rootX := d.Find(x)
	rootY := d.Find(y)
	if rootX == rootY {
		return false
	}
	if d.rank[rootX] < d.rank[rootY] {
		d.parent[rootX] = rootY
	} else if d.rank[rootX] > d.rank[rootY] {
		d.parent[rootY] = rootX
	} else {
		d.parent[rootY] = rootX
		d.rank[rootX]++
	}
	return true
}

func Kruskal(edges []Edge, numVertices int) ([]Edge, int) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	dsu := NewDSU(numVertices)
	mst := []Edge{}
	totalWeight := 0

	for _, edge := range edges {
		if dsu.Union(edge.U, edge.V) {
			mst = append(mst, edge)
			totalWeight += edge.Weight
			if len(mst) == numVertices-1 {
				break
			}
		}
	}

	return mst, totalWeight
}

func main() {
	edges := []Edge{
		{0, 1, 4},
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 5},
		{2, 3, 8},
		{2, 1, 1},
		{2, 3, 9},
		{3, 4, 2},
	}

	numVertices := 5

	mst, totalWeight := Kruskal(edges, numVertices)

	fmt.Println("Minimum Spanning Tree (Kruskal):")
	for _, edge := range mst {
		fmt.Printf("%d - %d (weight: %d)\n", edge.U, edge.V, edge.Weight)
	}
	fmt.Printf("Total weight: %d\n", totalWeight)
}
package main

import (b
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Item struct {
	node     int
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func dijkstra(graph [][]Edge, start int) ([]int, []int) {
	n := len(graph)
	dist := make([]int, n)
	prev := make([]int, n)

	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}

	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.node

		if item.distance > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v := edge.To
			weight := edge.Weight

			alt := dist[u] + weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(pq, &Item{node: v, distance: alt})
			}
		}
	}

	return dist, prev
}

func reconstructPath(prev []int, target int) []int {
	path := []int{}
	for target != -1 {
		path = append([]int{target}, path...)
		target = prev[target]
	}
	return path
}

func main() {
	graph := [][]Edge{
		{{1, 4}, {2, 2}},
		{{2, 1}, {3, 5}},
		{{3, 8}, {1, 1}, {3, 9}},
		{{4, 2}},
		{},
	}

	start := 0
	dist, prev := dijkstra(graph, start)

	for i := 0; i < len(graph); i++ {
		if dist[i] == math.MaxInt32 {
			fmt.Printf("Vertex %d is unreachable from %d\n", i, start)
			continue
		}
		path := reconstructPath(prev, i)
		fmt.Printf("Jarak dari %d ke %d: %d | Path: ", start, i, dist[i])
		for j, node := range path {
			if j > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(node)
		}
		fmt.Println()
	}
}

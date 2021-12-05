package main

import (
	"fmt"
)

func main() {
	var vertexCount, edgeCount int
	fmt.Scanln(&vertexCount, &edgeCount)

	g := newGraph(vertexCount)

	var firstVertex, secondVertex int
	for i := 0; i < edgeCount; i++ {
		fmt.Scan(&firstVertex, &secondVertex)
		g.AddEdge(firstVertex, secondVertex)
	}

	fmt.Println(g)
}

func newGraph(vertexCount int) graph {
	graph := make(graph, vertexCount)
	for i := range graph {
		graph[i] = make([]bool, vertexCount)
	}
	return graph
}

type graph [][]bool

func (g *graph) AddEdge(firstVertex, secondVertex int) {
	(*g)[firstVertex-1][secondVertex-1] = true
	(*g)[secondVertex-1][firstVertex-1] = true
}

// 인접 행렬 그래프-무향

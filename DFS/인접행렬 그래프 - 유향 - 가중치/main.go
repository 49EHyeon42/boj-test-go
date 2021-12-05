package main

import (
	"fmt"
)

func main() {
	var vertexCount, edgeCount int
	fmt.Scanln(&vertexCount, &edgeCount)

	g := newGraph(vertexCount)

	var firstVertex, secondVertex, weight int
	for i := 0; i < edgeCount; i++ {
		fmt.Scan(&firstVertex, &secondVertex, &weight)
		g.AddEdge(firstVertex, secondVertex, weight)
	}

	fmt.Println(g)
}

func newGraph(vertexCount int) graph {
	graph := make(graph, vertexCount)
	for i := range graph {
		graph[i] = make([]int, vertexCount)
	}
	return graph
}

type graph [][]int

func (g *graph) AddEdge(firstVertex, secondVertex, weight int) {
	// 가중치 관련 코드(거리, 중복 등)는 여기에 작성
	(*g)[firstVertex-1][secondVertex-1] = weight
}

// 인접 행렬 그래프-유향-가중치

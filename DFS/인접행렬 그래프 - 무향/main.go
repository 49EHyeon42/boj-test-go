package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  [][]bool
)

func main() {
	var vertexCount, edgeCount int
	fmt.Fscanln(reader, &vertexCount, &edgeCount)

	graph = make([][]bool, vertexCount+1)
	for i := range graph {
		graph[i] = make([]bool, vertexCount+1)
	}

	var firstVertex, secondVertex int
	for i := 0; i < edgeCount; i++ {
		fmt.Fscan(reader, &firstVertex, &secondVertex)
		graph[firstVertex][secondVertex] = true
		graph[secondVertex][firstVertex] = true
	}

	fmt.Fprintln(writer, graph)

	writer.Flush()
}

/*
인접행렬-유향 그래프
*/

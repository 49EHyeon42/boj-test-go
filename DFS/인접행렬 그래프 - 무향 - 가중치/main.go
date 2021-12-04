package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  [][]int
)

func main() {
	var vertexCount, edgeCount int
	fmt.Fscanln(reader, &vertexCount, &edgeCount)

	graph = make([][]int, vertexCount+1)
	for i := range graph {
		graph[i] = make([]int, vertexCount+1)
	}

	var firstVertex, secondVertex, weight int
	for i := 0; i < edgeCount; i++ {
		fmt.Fscan(reader, &firstVertex, &secondVertex)
		graph[firstVertex][secondVertex] = weight
		graph[secondVertex][firstVertex] = weight
	}

	fmt.Fprintln(writer, graph)

	writer.Flush()
}

/*
인접행렬-유향 그래프
*/

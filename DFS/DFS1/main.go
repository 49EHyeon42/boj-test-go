package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	graph   [][]int
	visited []bool
)

func main() {
	var vertexCount, edgeCount, startValue int
	fmt.Fscanln(reader, &vertexCount, &edgeCount, &startValue)

	graph, visited = make([][]int, vertexCount+1), make([]bool, vertexCount+1)

	var first, second int
	for i := 0; i < edgeCount; i++ {
		fmt.Fscanln(reader, &first, &second)
		graph[first] = append(graph[first], second)
	}

	DFS(startValue)

	writer.Flush()
}

func DFS(value int) {
	visited[value] = true
	fmt.Fprintf(writer, "%d ", value)
	for i := 0; i < len(graph[value]); i++ {
		temp := graph[value][i] // 인접 정점 확인
		if !visited[temp] {
			DFS(temp)
		}
	}
}

/*
인접리스트-무방향-DFS 그래프
*/

package main

import "fmt"

var visited []bool

func main() {
	var edgeCount int
	fmt.Scanln(&edgeCount)

	g := newGraph()

	visited = make([]bool, 4+1)

	var firstVertex, secondVertex int
	for i := 0; i < edgeCount; i++ {
		fmt.Scan(&firstVertex, &secondVertex)
		g.AddEdge(firstVertex, secondVertex)
	}

	g.Print()
}

func newNode(vertex int, link *node) node {
	return node{vertex: vertex, link: link}
}

type node struct {
	vertex int
	link   *node
}

func newGraph() graph {
	return graph{}
}

type graph []node

func (g graph) isExist(vertex int) (int, bool) {
	for i := 0; i < len(g); i++ {
		if g[i].vertex == vertex {
			return i, true
		}
	}
	return -1, false
}

func (g *graph) AddEdge(firstVertex, secondVertex int) {
	_, isExist := g.isExist(firstVertex)
	if !isExist {
		*g = append(*g, newNode(firstVertex, nil))
	}

	index, _ := g.isExist(firstVertex)

	addNode := newNode(secondVertex, (*g)[index].link)
	(*g)[index].link = &addNode
}

func (g graph) Print() {
	for i := 0; i < len(g); i++ {
		link := g[i].link
		fmt.Printf("정점 %d 인집 리스트 ", i)
		for link != nil {
			fmt.Printf("-> %d ", link.vertex)
			link = link.link
		}
		fmt.Println()
	}
}

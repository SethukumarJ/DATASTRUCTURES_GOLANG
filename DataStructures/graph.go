package main

import "fmt"

type Graph struct {
	Map map[int][]int
}

func (g *Graph) AddVertex(vertex int) {
	g.Map[vertex] = []int{}
}

func (g *Graph) insert(vertex int, edge int, bidirectional bool) {
	_, exists := g.Map[vertex]
	_, exists2 := g.Map[edge]
	if !exists {
		g.AddVertex(vertex)
	}
	if !exists2 {
		g.AddVertex(edge)
	}
	g.Map[vertex] = append(g.Map[vertex], edge)
	if bidirectional {
		g.Map[edge] = append(g.Map[edge], vertex)
	}
}

func (g *Graph) display() {

	fmt.Println(g.Map)
}

func main() {
	g := &Graph{Map: make(map[int][]int)}

	g.insert(3, 5, true)
	g.insert(3, 4, true)
	g.insert(5, 6, false)

	g.display()
	// Output: map[1:[2 2] 2:[1 3] 3:[2 4]]
}

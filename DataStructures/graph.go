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
	_,exists1 := g.Map[edge]

	if !exists {

		g.AddVertex(vertex)
	}
	if !exists1 {
		g.AddVertex(edge)
	}

	g.Map[vertex] = append(g.Map[vertex], edge)

	if bidirectional {
		g.Map[edge] = append(g.Map[edge], vertex)
	}
}
func (g *Graph) BFS(start int) []int {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true
	result := []int{}

	for len(queue) >0 {

			vertex := queue[0]
			result = append(result, vertex)
			queue = queue[1:]
			for _, v := range g.Map[vertex] {

				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}	
	}

	for key := range g.Map {
		
		if !visited[key] {
			g.BFS(key)
		}
	}
	return result
}


func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{start}


	var DFSUTIL func(vertex int)
	DFSUTIL =  func (vertex int) {

		visited[vertex]=true
		for _,v := range g.Map[vertex] {

			if !visited[v] {

				DFSUTIL(v)
				result = append(result, v)
			}
		}

	}

	DFSUTIL(start)
	for key := range g.Map{
		if !visited[key]{
			g.DFS(key)
		}
	}
	return result
}


func (g *Graph) display() {

	fmt.Println(g.Map)
}

func main() {
	g := &Graph{Map: make(map[int][]int)}

	g.insert(3, 5, true)
	g.insert(3, 4, true)
	g.insert(5, 6, false)
	jello := g.BFS(3)
	jllo := g.DFS(3)
	fmt.Print("hello",jello)
	fmt.Print("jllo",jllo)
	g.display()
	
}


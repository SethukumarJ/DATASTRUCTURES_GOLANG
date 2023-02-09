package main

import "fmt"

type Graph struct {
    AdjMatrix [][]int
    Vertices  []string
}

func (g *Graph) AddVertex(v string) {
    g.Vertices = append(g.Vertices, v)
    for i := range g.AdjMatrix {
        g.AdjMatrix[i] = append(g.AdjMatrix[i], 0)
    }
    g.AdjMatrix = append(g.AdjMatrix, make([]int, len(g.Vertices)))
}

func (g *Graph) AddEdge(v1, v2 string) {
    i1 := g.getIndex(v1)
    i2 := g.getIndex(v2)
    g.AdjMatrix[i1][i2] = 1
    g.AdjMatrix[i2][i1] = 1
}

func (g *Graph) getIndex(v string) int {
    for i, vtx := range g.Vertices {
        if vtx == v {
            return i
        }
    }
    return -1
}

func main() {
    g := Graph{
        AdjMatrix: [][]int{},
        Vertices:  []string{},
    }

    g.AddVertex("A")
    g.AddVertex("B")
    g.AddVertex("C")
    g.AddVertex("D")

    g.AddEdge("A", "B")
    g.AddEdge("B", "C")
    g.AddEdge("C", "D")
    g.AddEdge("D", "A")

    fmt.Println(g.AdjMatrix)
    fmt.Println(g.Vertices)
}

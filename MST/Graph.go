package main

import "fmt"

// Vertex
type Vertex struct {
	edges []*Vertex
}

func (v Vertex) addEdge(vToAdd Vertex) Vertex {
	edges := v.edges
	edges = append(edges, &vToAdd)
	return Vertex{edges}
}

func (v Vertex) delEdge(vToDel Vertex) Vertex {
	edges := v.edges
	nEdges := len(edges)
	toDel := -1
	for i := 0; i < nEdges; i++ {
		toDel = i
	}
	if toDel > -1 {
		edges = append(edges[:toDel], edges[toDel+1:]...)
	}
	return Vertex{edges}
}

type Muter interface {
	addEdge(vToAdd Vertex) Vertex
	delEdge(vToDel Vertex) Vertex
}

type Graph struct {
	vertices []*Vertex
}

func main() {
	v1 := Vertex{}
	v2 := Vertex{}
	v1 = v1.addEdge(v2)

	graph := Graph{[]*Vertex{&v1, &v2}}

	fmt.Println(graph)
	fmt.Println(v1)
}

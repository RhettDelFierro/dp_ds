package ds

import "testing"

func NewGraph() *Graph {
	return &Graph{
		make(map[int]*GraphNode),
	}
}

func TestInsert(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex(5)
	graph.AddVertex(2)
	graph.AddVertex(6)
	graph.AddVertex(7)
	graph.AddVertex(8)
	graph.AddEdge(2, 5)
	graph.AddEdge(6, 7)
	graph.AddEdge(6, 8)
	graph.AddEdge(7, 5)
	graph.AddEdge(7, 5)
	if graph.GetEdge(7, 5).Weight != 2 {
		t.Errorf("Graph.AddEdge and Graph.GetEdge do not work together. Wanted: '%d', got: '%d'", 2, graph.GetEdge(7, 5).Weight)
	}
}

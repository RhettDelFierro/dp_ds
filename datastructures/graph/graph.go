package ds

type Edge struct {
	Weight int
}

type GraphNode struct {
	Val   int
	Edges map[int]*Edge
}

type Graph struct {
	Vertices map[int]*GraphNode
}

func (g *Graph) AddVertex(val int) {
	if g.Vertices == nil {
		g.Vertices = make(map[int]*GraphNode)
	}
	if g.Vertices[val] == nil {
		g.Vertices[val] = &GraphNode{Val: val}
	}
}

func (g *Graph) RemoveVertex(val int) {
	// delete the vertex:
	if g.Vertices[val] != nil {
		delete(g.Vertices, val)
	}

	// delete what is pointing to the vertex:
	for key, val := range g.Vertices {
		if g.Vertices[key].Edges[val.Val] != nil {
			delete(g.Vertices[key].Edges, val.Val)
		}
	}
}

func (g *Graph) GetVertex(val int) *GraphNode {
	return g.Vertices[val]
}

func (g *Graph) AddEdge(start, end int) {
	beginningVertex := g.Vertices[start]
	endVertex := g.Vertices[end]

	if beginningVertex != nil && endVertex != nil {
		if beginningVertex.Edges == nil {
			beginningVertex.Edges = make(map[int]*Edge)
		}

		if beginningVertex.Edges[end] != nil {
			beginningVertex.Edges[end].Weight++
		} else {
			beginningVertex.Edges[end] = &Edge{Weight: 1}
		}
	}
}

func (g *Graph) RemoveEdge(start, end int) {
	beginningVertex := g.Vertices[start]
	endVertex := g.Vertices[end]

	if beginningVertex != nil && endVertex != nil {
		if beginningVertex.Edges[end] != nil {
			delete(beginningVertex.Edges, end)
		}
	}
}

func (g *Graph) GetEdge(start, end int) *Edge {
	edge := g.Vertices[start].Edges[end]

	if edge == nil {
		return nil
	}

	return g.Vertices[start].Edges[end]
}

func (g *Graph) Neighbors(val int) map[int]*Edge {
	vertex := g.Vertices[val]

	if vertex == nil {
		return nil
	}

	return vertex.Edges
}

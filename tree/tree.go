package tree

import "errors"

var ErrVertexIdDoesNotExist = errors.New("vertex with such Id does not exist")
var ErrNoVertex = errors.New("tree does not contain any vertex")

// a directed graph, it is called a tree but it could contain cycles
type Generator[VertexId comparable, EdgeId comparable] interface {
	Tree() (*Vertex[VertexId, EdgeId], error) // pick any vertex and return the whole tree
	Successors(id VertexId) (*Vertex[VertexId, EdgeId], error)
	Predecessors(id VertexId) (*Vertex[VertexId, EdgeId], error)
}

type Vertex[VertexId comparable, EdgeId comparable] struct {
	Id    VertexId
	Edges []Edge[VertexId, EdgeId]
}

type Edge[VertexId comparable, EdgeId comparable] struct {
	Id   EdgeId
	From VertexId
	To   *Vertex[VertexId, EdgeId]
}

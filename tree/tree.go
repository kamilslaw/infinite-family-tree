package tree

import "errors"

var ErrVertexIdDoesNotExist = errors.New("vertex with such Id does not exist")

type Generator[VertexId comparable, EdgeId comparable] interface {
	Root() (*Vertex[VertexId, EdgeId], error) // pick any vertex and return the whole tree
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

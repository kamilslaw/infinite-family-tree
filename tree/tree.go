package tree

type Visitor[Node any] interface {
	Root() *Tree[Node]
	Successors(n Node) *Tree[Node]
	Predecessors(n Node) *Tree[Node]
}

// Tree struct represents a directed graph, there can be relations in one direction only
// (Successors() will return relations with a different direction than the Predecessors())
type Tree[Node any] struct {
	N      Node
	Others []Relation[Node]
}

type Relation[Node any] struct {
	N     *Tree[Node any]
	Name string
}
package family

/*
	The tree.Generator interface implementation for the Family struct
	The people with no relationships won't be included
*/

import (
	"github.com/kamilslaw/infinite-family-tree/tree"
	"github.com/kamilslaw/infinite-family-tree/utils"
)

var _ tree.Generator[PersonId, RelationshipId] = (*Family)(nil)

type vertex = tree.Vertex[PersonId, RelationshipId]
type edge = tree.Edge[PersonId, RelationshipId]

func (f *Family) Tree() (*vertex, error) {
	if len(f.people) == 0 {
		return nil, tree.ErrNoVertex
	}
	return f.Successors(utils.PickAnyKey(f.people))
}

func (f *Family) Successors(id PersonId) (*vertex, error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}

	vertexes, edges := f.generateGraphElements()

	for _, p := range f.people {
		for _, r := range directSuccessors(p.Id, f) {
			vertexes[r.From].Edges = append(vertexes[r.From].Edges, *edges[r.Id])
		}
	}

	tree := vertexes[id]
	return tree, nil
}

func (f *Family) Predecessors(id PersonId) (*vertex, error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}

	vertexes, edges := f.generateGraphElements()

	for _, p := range f.people {
		for _, r := range directPredecessors(p.Id, f) {
			vertexes[r.From].Edges = append(vertexes[r.From].Edges, *edges[r.Id])
		}
	}

	tree := vertexes[id]
	return tree, nil
}

func (f *Family) generateGraphElements() (map[PersonId]*vertex, map[RelationshipId]*edge) {
	vertexes := map[PersonId]*vertex{}
	for _, p := range f.people {
		vertexes[p.Id] = &vertex{Id: p.Id, Edges: make([]edge, 0)}
	}

	edges := map[RelationshipId]*edge{}
	for _, r := range f.relationships {
		edges[r.Id] = &edge{Id: r.Id, From: r.From, To: vertexes[r.To]}
	}

	return vertexes, edges
}

func directSuccessors(id PersonId, f *Family) []*Relationship {
	from := f.relationshipsFromPerson[id]
	r := make([]*Relationship, len(from))
	copy(r, from)

	to := f.relationshipsToPerson[id]
	for _, t := range to {
		if !t.OneSided() {
			r = append(r, t)
		}
	}

	return r
}

func directPredecessors(id PersonId, f *Family) []*Relationship {
	to := f.relationshipsToPerson[id]
	r := make([]*Relationship, len(to))
	copy(r, to)

	from := f.relationshipsFromPerson[id]
	for _, f := range from {
		if !f.OneSided() {
			r = append(r, f)
		}
	}

	return r
}

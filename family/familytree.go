package family

/*
	The tree.Generator interface implementation for the Family struct
	The people with no relationships won't be included
*/

import "github.com/kamilslaw/infinite-family-tree/tree"

var _ tree.Generator[PersonId, RelationshipId] = (*Family)(nil)

type vertex = tree.Vertex[PersonId, RelationshipId]
type edge = tree.Edge[PersonId, RelationshipId]

func (f *Family) Tree() (*vertex, error) {
	//TODO implement me
	panic("implement me")
}

func (f *Family) Successors(id PersonId) (*vertex, error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}

	t := vertex{Id: id, Edges: make([]edge, 0)}
	v := map[PersonId]*vertex{id: &t}

	return &t, nil
}

func (f *Family) Predecessors(id PersonId) (*vertex, error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}
	//TODO implement me
	panic("implement me")
}

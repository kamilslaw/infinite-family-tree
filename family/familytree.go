package family

/*
	The tree.Generator interface implementation for the Family struct
	The people with no relationships won't be included
*/

import "github.com/kamilslaw/infinite-family-tree/tree"

var _ tree.Generator[PersonId, RelationshipId] = (*Family)(nil)

func (f *Family) Root() (*tree.Vertex[PersonId, RelationshipId], error) {
	//TODO implement me
	panic("implement me")
}

func (f *Family) Successors(id PersonId) (*tree.Vertex[PersonId, RelationshipId], error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}
	//TODO implement me
	panic("implement me")
}

func (f *Family) Predecessors(id PersonId) (*tree.Vertex[PersonId, RelationshipId], error) {
	if _, ok := f.people[id]; !ok {
		return nil, tree.ErrVertexIdDoesNotExist
	}
	//TODO implement me
	panic("implement me")
}

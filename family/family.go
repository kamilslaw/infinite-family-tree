package family

import "errors"
import "github.com/kamilslaw/infinite-family-tree/utils"

// TODO: more efficient way of storing the relationships - for now it should be fine (Unless your family has tens of thousands of members)

// TODO: add implementation of interface to support the tree/graph visualisation/tracking (so the tree is generic and not family-tree dependent)

type Family struct {
	people    map[PersonId]*Person
	relations map[RelationshipId]*Relationship
}

var (
	ErrPersonIdExists       = errors.New("person with such Id already exists")
	ErrPersonIdDoesNotExist = errors.New("person with such Id does not exist")
)

func (f *Family) AddPerson(p *Person) error {
	if _, ok := f.people[p.Id]; ok {
		return ErrPersonIdExists
	}

	f.people[p.Id] = p
	return nil
}

func (f *Family) AddRelationship(r *Relationship) error {
	if _, ok := f.people[r.Person]; !ok {
		return ErrPersonIdDoesNotExist
	}
	if _, ok := f.people[r.Of]; !ok {
		return ErrPersonIdDoesNotExist
	}
	if err := r.CheckIfAllowed(utils.MapValues(f.relations)); err != nil {
		return err
	}

	f.relations[r.Id] = r
	return nil
}

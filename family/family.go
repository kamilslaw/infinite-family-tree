package family

import (
	"errors"

	"github.com/kamilslaw/infinite-family-tree/utils"
)

type Family struct {
	people                  map[PersonId]*Person
	relationships           map[RelationshipId]*Relationship
	relationshipsFromPerson map[PersonId][]*Relationship
	relationshipsToPerson   map[PersonId][]*Relationship
}

var (
	ErrPersonIdExists       = errors.New("person with such Id already exists")
	ErrPersonIdDoesNotExist = errors.New("person with such Id does not exist")
)

func Recreate(people []*Person, relationships []*Relationship) *Family {
	f := Family{}
	for _, p := range people {
		f.people[p.Id] = p
		f.relationshipsFromPerson[p.Id] = make([]*Relationship, 0)
		f.relationshipsToPerson[p.Id] = make([]*Relationship, 0)
	}
	for _, r := range relationships {
		f.relationships[r.Id] = r
		f.relationshipsFromPerson[r.From] = append(f.relationshipsFromPerson[r.From], r)
		f.relationshipsToPerson[r.To] = append(f.relationshipsToPerson[r.To], r)
	}
	return &f
}

func (f *Family) RawData() ([]*Person, []*Relationship) {
	return utils.MapValues(f.people), utils.MapValues(f.relationships)
}

func (f *Family) AddPerson(p *Person) error {
	if _, ok := f.people[p.Id]; ok {
		return ErrPersonIdExists
	}

	f.people[p.Id] = p
	return nil
}

func (f *Family) AddRelationship(r *Relationship) error {
	if _, ok := f.people[r.From]; !ok {
		return ErrPersonIdDoesNotExist
	}
	if _, ok := f.people[r.To]; !ok {
		return ErrPersonIdDoesNotExist
	}
	if err := r.CheckIfAllowed(utils.MapValues(f.relationships)); err != nil {
		return err
	}

	f.relationships[r.Id] = r
	return nil
}

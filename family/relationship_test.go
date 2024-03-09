package family

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//func TestNewRelationship(t *testing.T) {
//	personId := PersonId(uuid.New())
//	is := Partner
//	of := personId
//
//	if _, err := NewRelationship(personId, is, of, time.Time{}, time.Time{}); err == nil || !errors.Is(err, ErrSelfRelationship) {
//		t.Errorf("got '%v', want '%v'", err, ErrSelfRelationship)
//	}
//}

//func TestNewRelationship(t *testing.T) {
//	person1 := PersonId
//}

func TestRelationship_Equal(t *testing.T) {
	personId1 := PersonId(uuid.New())
	personId2 := PersonId(uuid.New())
	get := func() *Relationship {
		return withKind(withOf(withPerson(getRelationship(), personId1), personId2), Spouse)
	}

	tests := []struct {
		name   string
		r      *Relationship
		other  *Relationship
		result bool
	}{
		{
			name:   "different relationship",
			r:      get(),
			other:  getRelationship(),
			result: false,
		},
		{
			name:   "different kind",
			r:      get(),
			other:  withKind(get(), Partner),
			result: false,
		},
		{
			name:   "different start time",
			r:      get(),
			other:  withStartedOn(get(), time.Now()),
			result: false,
		},
		{
			name:   "same",
			r:      get(),
			other:  get(),
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, tt.r.Equal(tt.other))
		})
	}
}

func TestRelationship_PeopleEqual(t *testing.T) {
	personId1 := PersonId(uuid.New())
	personId2 := PersonId(uuid.New())

	tests := []struct {
		name   string
		r      *Relationship
		other  *Relationship
		result bool
	}{
		{
			name:   "no common person",
			r:      getRelationship(),
			other:  getRelationship(),
			result: false,
		},
		{
			name:   "one common person",
			r:      withPerson(getRelationship(), personId1),
			other:  withPerson(getRelationship(), personId1),
			result: false,
		},
		{
			name:   "one common person on a different side of relationship",
			r:      withPerson(getRelationship(), personId1),
			other:  withOf(getRelationship(), personId1),
			result: false,
		},
		{
			name:   "two common persons",
			r:      withOf(withPerson(getRelationship(), personId1), personId2),
			other:  withOf(withPerson(getRelationship(), personId1), personId2),
			result: true,
		},
		{
			name:   "two common persons, but switched",
			r:      withOf(withPerson(getRelationship(), personId1), personId2),
			other:  withOf(withPerson(getRelationship(), personId2), personId1),
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, tt.r.PeopleEqual(tt.other))
		})
	}
}

func getRelationship() *Relationship {
	person := PersonId(uuid.New())
	is := Child
	of := PersonId(uuid.New())
	if r, e := NewRelationship(person, is, of, time.Time{}, time.Time{}); e != nil {
		panic(e)
	} else {
		return r
	}
}

func withPerson(r *Relationship, person PersonId) *Relationship {
	r.Person = person
	return r
}

func withOf(r *Relationship, of PersonId) *Relationship {
	r.Of = of
	return r
}

func withKind(r *Relationship, is RelationshipKind) *Relationship {
	r.Is = is
	return r
}

func withStartedOn(r *Relationship, startedOn time.Time) *Relationship {
	r.StartedOn = startedOn
	return r
}

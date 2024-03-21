package family

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var personId1 = uuid.New()
var personId2 = uuid.New()

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
	tests := []struct {
		name   string
		r      *Relationship
		other  *Relationship
		result bool
	}{
		{
			name:   "different relationship",
			r:      getRelationshipInstance(),
			other:  getRandomRelationship(),
			result: false,
		},
		{
			name:   "different kind",
			r:      getRelationshipInstance(),
			other:  withKind(getRelationshipInstance(), Partner),
			result: false,
		},
		{
			name:   "different start time",
			r:      getRelationshipInstance(),
			other:  withStartedOn(getRelationshipInstance(), time.Now()),
			result: false,
		},
		{
			name:   "same",
			r:      getRelationshipInstance(),
			other:  getRelationshipInstance(),
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
			r:      getRandomRelationship(),
			other:  getRandomRelationship(),
			result: false,
		},
		{
			name:   "one common person",
			r:      withPerson(getRandomRelationship(), personId1),
			other:  withPerson(getRandomRelationship(), personId1),
			result: false,
		},
		{
			name:   "one common person on a different side of relationship",
			r:      withPerson(getRandomRelationship(), personId1),
			other:  withOf(getRandomRelationship(), personId1),
			result: false,
		},
		{
			name:   "two common persons",
			r:      withOf(withPerson(getRandomRelationship(), personId1), personId2),
			other:  withOf(withPerson(getRandomRelationship(), personId1), personId2),
			result: true,
		},
		{
			name:   "two common persons, but switched",
			r:      withOf(withPerson(getRelationshipInstance(), personId1), personId2),
			other:  withOf(withPerson(getRelationshipInstance(), personId2), personId1),
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, tt.r.PeopleEqual(tt.other))
		})
	}
}

func TestRelationship_CheckIfAllowed_RelationshipCannotBeDuplicated(t *testing.T) {
	err := getRelationshipInstance().CheckIfAllowed([]*Relationship{withKind(getRelationshipInstance(), Partner)})
	assert.Nil(t, err)

	err = getRelationshipInstance().CheckIfAllowed([]*Relationship{getRelationshipInstance()})
	assert.ErrorIs(t, err, ErrRelationshipExists)
}

func TestRelationship_CheckIfAllowed_RelationshipIdCannotBeDuplicated(t *testing.T) {
	relationship := getRelationshipInstance()
	err := relationship.CheckIfAllowed([]*Relationship{withId(getRelationshipInstance(), relationship.Id)})
	assert.ErrorIs(t, err, ErrRelationshipIdExists)
}

func TestRelationship_CheckIfAllowed_ChildRelationCannotCoexistWithOtherRelationship(t *testing.T) {
	relation := getRelationshipInstance()
	relationWithSamePeopleButChild := withKind(getRelationshipInstance(), Child)

	err := relationWithSamePeopleButChild.CheckIfAllowed([]*Relationship{relation})
	assert.ErrorIs(t, err, ErrChildRelationshipCannotCoexist)

	err = relation.CheckIfAllowed([]*Relationship{relationWithSamePeopleButChild})
	assert.ErrorIs(t, err, ErrChildRelationshipCannotCoexist)

	relationWithSamePeopleButChildReverted := withOf(withPerson(relationWithSamePeopleButChild, relation.To), relation.From)
	err = relation.CheckIfAllowed([]*Relationship{relationWithSamePeopleButChildReverted})
	assert.ErrorIs(t, err, ErrChildRelationshipCannotCoexist)
}

func getRelationshipInstance() *Relationship {
	personId1Copy := make([]byte, len(personId1))
	copy(personId1Copy, personId1[:])
	personId2Copy := make([]byte, len(personId2))
	copy(personId2Copy, personId2[:])
	return withOf(withPerson(getRandomRelationship(), PersonId(personId1Copy)), PersonId(personId2Copy))
}

func getRandomRelationship() *Relationship {
	person := PersonId(uuid.New())
	of := PersonId(uuid.New())
	if r, e := NewRelationship(RelationshipId(uuid.New()), person, Friend, of, time.Time{}, time.Time{}); e != nil {
		panic(e)
	} else {
		return r
	}
}

func withId(r *Relationship, id RelationshipId) *Relationship {
	r.Id = id
	return r
}

func withPerson(r *Relationship, person PersonId) *Relationship {
	r.From = person
	return r
}

func withOf(r *Relationship, of PersonId) *Relationship {
	r.To = of
	return r
}

func withKind(r *Relationship, is RelationshipKind) *Relationship {
	r.Kind = is
	return r
}

func withStartedOn(r *Relationship, startedOn time.Time) *Relationship {
	r.StartedOn = startedOn
	return r
}

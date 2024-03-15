package family

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type RelationshipKind string

const (
	Partner RelationshipKind = "partner"
	Spouse  RelationshipKind = "spouse"
	Child   RelationshipKind = "child"
	Friend  RelationshipKind = "friend"
)

var (
	ErrSelfRelationship    = errors.New("relationship has to be between two different people")
	ErrUnknownRelationship = errors.New("unknown relationshi" +
		"p type")
	ErrRelationshipEndedBeforeStarted = errors.New("relationship could not be ended before started")
	ErrRelationshipCouldNotBeEnded    = errors.New("such relationship kind could not be ended")

	ErrRelationshipIdExists           = errors.New("relationship with such Id already exists")
	ErrRelationshipExists             = errors.New("relationship already exists")
	ErrChildRelationshipCannotCoexist = errors.New("the child relationship cannot coexist with other types of relationship")
)

type RelationshipId uuid.UUID

type Relationship struct {
	Id        RelationshipId
	Person    PersonId
	Is        RelationshipKind
	Of        PersonId
	StartedOn time.Time
	EndedOn   time.Time
}

func (r *Relationship) Equal(other *Relationship) bool {
	return r.Person == other.Person &&
		r.Is == other.Is &&
		r.Of == other.Of &&
		r.StartedOn.Equal(other.StartedOn) &&
		r.EndedOn.Equal(other.EndedOn)
}

func (r *Relationship) PeopleEqual(other *Relationship) bool {
	return (r.Person == other.Person && r.Of == other.Of) ||
		(r.Person == other.Of && r.Of == other.Person)
}

func NewRelationship(id RelationshipId, person PersonId, is RelationshipKind,
	of PersonId, startedOn time.Time, endedOn time.Time) (*Relationship, error) {

	if person == of {
		return &Relationship{}, ErrSelfRelationship
	}

	if is != Partner && is != Spouse && is != Child && is != Friend {
		return &Relationship{}, ErrUnknownRelationship
	}

	if !startedOn.IsZero() && !endedOn.IsZero() && endedOn.Before(startedOn) {
		return &Relationship{}, ErrRelationshipEndedBeforeStarted
	}

	if !endedOn.IsZero() && is != Partner && is != Spouse && is != Friend {
		return &Relationship{}, ErrRelationshipCouldNotBeEnded
	}

	return &Relationship{Id: id, Person: person, Is: is, Of: of, StartedOn: startedOn}, nil
}

func (r *Relationship) CheckIfAllowed(others []*Relationship) error {
	for _, other := range others {
		if other.Id == r.Id {
			return ErrRelationshipIdExists
		}
		if other.Equal(r) {
			return ErrRelationshipExists
		}

		anyChildRelationship := r.Is == Child || other.Is == Child
		if anyChildRelationship && r.PeopleEqual(other) {
			return ErrChildRelationshipCannotCoexist
		}
	}

	// todo: implement - check if exists in the family tree already (and other checks)
	return nil
}

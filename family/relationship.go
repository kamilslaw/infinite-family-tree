package family

import (
	"errors"
	"time"
)

type RelationshipKind string

const (
	Partner RelationshipKind = "partner"
	Spouse  RelationshipKind = "spouse"
	Child   RelationshipKind = "child"
	Friend  RelationshipKind = "friend"
)

var (
	ErrSelfRelationship               = errors.New("relationship has to be between two different people")
	ErrUnknownRelationship            = errors.New("unknown relationship type")
	ErrRelationshipEndedBeforeStarted = errors.New("relationship could not be ended before started")
	ErrRelationshipCouldNotBeEnded    = errors.New("such relationship kind could not be ended")

	ErrRelationshipExists             = errors.New("relationship already exists")
	ErrChildRelationshipCannotCoexist = errors.New("the child relationship cannot coexist with other types of relationship")
)

type Relationship struct {
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

func NewRelationship(person PersonId, is RelationshipKind, of PersonId,
	startedOn time.Time, endedOn time.Time) (*Relationship, error) {

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

	return &Relationship{Person: person, Is: is, Of: of, StartedOn: startedOn}, nil
}

func (r *Relationship) CheckIfAllowed(others []*Relationship) (bool, error) {
	for _, other := range others {
		if other.Equal(r) {
			return false, ErrRelationshipExists
		}

		if r.Is == Child &&
			(other.Person == r.Of || other.Of == r.Of) &&
			(other.Person == r.Person || other.Of == r.Person) {
			return false, ErrChildRelationshipCannotCoexist
		}
	}

	// todo: implement - check if exists in the family tree already
	return true, nil
}

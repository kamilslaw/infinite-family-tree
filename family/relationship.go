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
	From      PersonId
	Kind      RelationshipKind
	To        PersonId
	StartedOn time.Time
	EndedOn   time.Time
}

func (r *Relationship) OneSided() bool {
	return r.Kind == Child
}

func (r *Relationship) Equal(other *Relationship) bool {
	return r.From == other.From &&
		r.Kind == other.Kind &&
		r.To == other.To &&
		r.StartedOn.Equal(other.StartedOn) &&
		r.EndedOn.Equal(other.EndedOn)
}

func (r *Relationship) PeopleEqual(other *Relationship) bool {
	return (r.From == other.From && r.To == other.To) ||
		(r.From == other.To && r.To == other.From)
}

func NewRelationship(id RelationshipId, from PersonId, kind RelationshipKind,
	to PersonId, startedOn time.Time, endedOn time.Time) (*Relationship, error) {

	if from == to {
		return &Relationship{}, ErrSelfRelationship
	}

	if kind != Partner && kind != Spouse && kind != Child && kind != Friend {
		return &Relationship{}, ErrUnknownRelationship
	}

	if !startedOn.IsZero() && !endedOn.IsZero() && endedOn.Before(startedOn) {
		return &Relationship{}, ErrRelationshipEndedBeforeStarted
	}

	if !endedOn.IsZero() && kind != Partner && kind != Spouse && kind != Friend {
		return &Relationship{}, ErrRelationshipCouldNotBeEnded
	}

	return &Relationship{Id: id, From: from, Kind: kind, To: to, StartedOn: startedOn}, nil
}

func (r *Relationship) CheckIfAllowed(others []*Relationship) error {
	for _, other := range others {
		if other.Id == r.Id {
			return ErrRelationshipIdExists
		}
		if other.Equal(r) {
			return ErrRelationshipExists
		}

		anyChildRelationship := r.Kind == Child || other.Kind == Child
		if anyChildRelationship && r.PeopleEqual(other) {
			return ErrChildRelationshipCannotCoexist
		}
	}

	// todo: implement - check if exists in the family tree already (and other checks)
	return nil
}

package family

import (
	"errors"
	"time"
)

type RelationKind string

const (
	Partner RelationKind = "partner"
	Spouse  RelationKind = "spouse"
	Child   RelationKind = "child"
	Friend  RelationKind = "friend"
)

var (
	ErrSelfRelation               = errors.New("relation has to be between two different people")
	ErrUnknownRelation            = errors.New("unknown relation type")
	ErrRelationEndedBeforeStarted = errors.New("relation could not be ended before started")
	ErrRelationCouldNotBeEnded    = errors.New("such relation kind could not be ended")
)

type Relation struct {
	Person    PersonId
	Is        RelationKind
	Of        PersonId
	StartedOn time.Time
	EndedOn   time.Time
}

func NewRelation(person PersonId, is RelationKind, of PersonId,
	startedOn time.Time, endedOn time.Time) (Relation, error) {

	if person == of {
		return Relation{}, ErrSelfRelation
	}

	if is != Partner && is != Spouse && is != Child && is != Friend {
		return Relation{}, ErrUnknownRelation
	}

	if !startedOn.IsZero() && !endedOn.IsZero() && endedOn.Before(startedOn) {
		return Relation{}, ErrRelationEndedBeforeStarted
	}

	if !endedOn.IsZero() && is != Partner && is != Spouse && is != Friend {
		return Relation{}, ErrRelationCouldNotBeEnded
	}

	return Relation{Person: person, Is: is, Of: of, StartedOn: startedOn}, nil
}

func IsRelationAllowed(r *Relation, other []*Relation) bool {
	// todo: implement - for instance child can't be a parent of their parent
	return true
}

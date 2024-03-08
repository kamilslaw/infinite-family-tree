package family

import (
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewRelationship(t *testing.T) {
	personId := PersonId(uuid.New())
	is := Partner
	of := personId

	if _, err := NewRelationship(personId, is, of, time.Time{}, time.Time{}); err == nil || !errors.Is(err, ErrSelfRelationship) {
		t.Errorf("got '%v', want '%v'", err, ErrSelfRelationship)
	}
}

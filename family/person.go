package family

import (
	"time"

	"github.com/google/uuid"
)

type PersonId uuid.UUID

type Person struct {
	Id          PersonId
	FirstNames  []string
	SecondNames []string
	BornOn      time.Time
	DiedOn      time.Time
	Attributes  map[string]string
}

// ByNameMatcher - comparing names is culture dependent,
// e.g. when comparing Dutch people we might want to ignore the 'van' prefix;
// returns [0-1], where 0 means 100% match and 1 means no match at all
type ByNameMatcher func(l, r *Person) uint

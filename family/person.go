package family

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID     uuid.UUID
	Names  []string
	BornOn time.Time
	DiedOn time.Time
}

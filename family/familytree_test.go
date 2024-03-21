package family

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kamilslaw/infinite-family-tree/utils"
	"github.com/stretchr/testify/assert"
)

var A = PersonId(uuid.New())
var B = PersonId(uuid.New())
var C = PersonId(uuid.New())
var D = PersonId(uuid.New())
var E = PersonId(uuid.New())
var F = PersonId(uuid.New())

var f *Family = Recreate(nil, nil)

func init() {
	utils.PanicIfError(f.AddPerson(getPerson(A)))
	utils.PanicIfError(f.AddPerson(getPerson(B)))
	utils.PanicIfError(f.AddPerson(getPerson(C)))
	utils.PanicIfError(f.AddPerson(getPerson(D)))
	utils.PanicIfError(f.AddPerson(getPerson(E)))
	utils.PanicIfError(f.AddPerson(getPerson(F)))
	utils.PanicIfError(f.AddRelationship(getRelationship(A, Child, B)))
	utils.PanicIfError(f.AddRelationship(getRelationship(B, Spouse, C)))
	utils.PanicIfError(f.AddRelationship(getRelationship(B, Child, D)))
	utils.PanicIfError(f.AddRelationship(getRelationship(C, Child, D)))
	utils.PanicIfError(f.AddRelationship(getRelationship(C, Child, E)))
	utils.PanicIfError(f.AddRelationship(getRelationship(D, Partner, F)))
	utils.PanicIfError(f.AddRelationship(getRelationship(F, Friend, E)))
}

func TestFamilyTree_Successors(t *testing.T) {
	tree, err := f.Successors(A)

	assert.Nil(t, err)
	assert.NotNil(t, tree)
	assert.ElementsMatch()
}

func getPerson(id PersonId) *Person {
	return &Person{Id: id}
}

func getRelationship(from PersonId, kind RelationshipKind, to PersonId) *Relationship {
	r, _ := NewRelationship(RelationshipId(uuid.New()), from, kind, to, time.Time{}, time.Time{})
	return r
}

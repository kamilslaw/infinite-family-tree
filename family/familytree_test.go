package family

import (
	"fmt"
	"testing"
	"time"

	"github.com/kamilslaw/infinite-family-tree/tree"
	"github.com/kamilslaw/infinite-family-tree/utils"
	"github.com/stretchr/testify/assert"
)

var PA = PersonId([16]byte{0})
var PB = PersonId([16]byte{1})
var PC = PersonId([16]byte{2})
var PD = PersonId([16]byte{3})
var PE = PersonId([16]byte{4})
var PF = PersonId([16]byte{5})

var RA = RelationshipId([16]byte{0})
var RB = RelationshipId([16]byte{1})
var RC = RelationshipId([16]byte{2})
var RD = RelationshipId([16]byte{3})
var RE = RelationshipId([16]byte{4})
var RF = RelationshipId([16]byte{5})
var RG = RelationshipId([16]byte{6})

var f *Family = Recreate(nil, nil)

func init() {
	utils.PanicIfError(f.AddPerson(getPerson(PA)))
	utils.PanicIfError(f.AddPerson(getPerson(PB)))
	utils.PanicIfError(f.AddPerson(getPerson(PC)))
	utils.PanicIfError(f.AddPerson(getPerson(PD)))
	utils.PanicIfError(f.AddPerson(getPerson(PE)))
	utils.PanicIfError(f.AddPerson(getPerson(PF)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RA, PA, Child, PB)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RB, PB, Spouse, PC)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RC, PB, Child, PD)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RD, PC, Child, PD)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RE, PC, Child, PE)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RF, PD, Partner, PF)))
	utils.PanicIfError(f.AddRelationship(getRelationship(RG, PF, Friend, PE)))
}

func TestFamilyTree_Successors(t *testing.T) {
	v, err := f.Successors(PA)

	assert.Nil(t, err)
	assert.NotNil(t, v)

	treeStr := tree.ToStr(v)

	expected := fmt.Sprintf(`%v (1)
  %v (2)
`, PA, PB)
	assert.Equal(t, expected, treeStr)

	// assert.Equal(t, PA, tree.Id)
	// assert.ElementsMatch(t, []RelationshipId{RA}, edgesToIds(tree.Edges))
}

func getPerson(id PersonId) *Person {
	return &Person{Id: id}
}

func getRelationship(id RelationshipId, from PersonId,
	kind RelationshipKind, to PersonId) *Relationship {
	r, _ := NewRelationship(id, from, kind, to, time.Time{}, time.Time{})
	return r
}

// func edgesToIds(edges []edge) []RelationshipId {
// 	ids := []RelationshipId{}
// 	for _, e := range edges {
// 		ids = append(ids, e.Id)
// 	}
// 	return ids
// }

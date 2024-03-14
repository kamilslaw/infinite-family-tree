package family

// TODO: more efficient way of storing the relationships - for now it should be fine (Unless your family has tens of thousands of members)

// TODO: add implementation of interface to support the tree/graph visualisation/tracking (so the tree is generic and not family-tree dependent)

type Family struct {
	people    map[PersonId]Person
	relations map[RelationshipId]Relationship
}

func (f *Family) AddPerson(p Person) error {
	f.people[p.Id] = p

	return nil
}

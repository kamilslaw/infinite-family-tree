package family

// TODO: more efficient way of storing the relationships - for now it should be fine (Unless your family has tens of thousands of members)

type Family struct {
	people    []Person
	relations []Relationship
}

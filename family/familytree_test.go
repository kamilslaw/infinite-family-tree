package family

import "github.com/google/uuid"

var f Family = Family{}

func init() {
	f.AddPerson(getPerson("a"))
	f.AddPerson(getPerson("b"))
	f.AddPerson(getPerson("c"))
	f.AddPerson(getPerson("d"))
}

func getPerson(name string) *Person {
	return &Person{Id: PersonId(uuid.New()), FirstNames: []string{name}}
}

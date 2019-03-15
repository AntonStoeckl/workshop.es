package domain

import "workshop.es/person/domain/value"

type Persons interface {
	Save(person Person) error
	Find(id value.ID) (Person, error)
}

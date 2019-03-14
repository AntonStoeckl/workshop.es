package value

import (
	"errors"
	"strings"
)

var (
	ErrNameInvalidInputForGivenName  = errors.New("name: invalid input given for givenname")
	ErrNameInvalidInputForFamilyName = errors.New("name: invalid input given for familyname")
)

type Name struct {
	givenName  string
	familyName string
}

func NewName(givenName, familyName string) (*Name, error) {
	newName := NewNameWithoutValidation(givenName, familyName)

	if err := newName.validate(); err != nil {
		return nil, err
	}

	return newName, nil
}

func NewNameWithoutValidation(givenName, familyName string) *Name {
	newName := &Name{
		givenName:  strings.TrimSpace(givenName),
		familyName: strings.TrimSpace(familyName),
	}

	return newName
}

func (name *Name) GivenName() string {
	return name.givenName
}

func (name *Name) FamilyName() string {
	return name.familyName
}

func (name *Name) validate() error {
	if name.givenName == "" {
		return ErrNameInvalidInputForGivenName
	}

	if name.familyName == "" {
		return ErrNameInvalidInputForFamilyName
	}

	return nil
}

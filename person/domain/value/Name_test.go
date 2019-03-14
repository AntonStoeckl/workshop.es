package value

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NameTestSuite struct {
	suite.Suite
}

func Test_NameTestSuite(t *testing.T) {
	tests := new(NameTestSuite)
	suite.Run(t, tests)
}

func (s *NameTestSuite) Test_Name_NewName() {
	// given
	givenName := "Franz"
	familyName := "Kafka"

	// when
	name, err := NewName(givenName, familyName)

	// then
	s.NoError(err, "it should create a Name from valid input")
	s.Equal(givenName, name.givenName, "id should expose expected GivenName")
	s.Equal(familyName, name.familyName, "it should expose expected FamilyName")
}

func (s *NameTestSuite) Test_Name_NewName_WithInvalidGivenName() {
	// given
	givenName := " "
	familyName := "Kafka"

	// when
	_, err := NewName(givenName, familyName)

	// then
	s.EqualError(err, ErrNameInvalidInputForGivenName.Error())
}

func (s *NameTestSuite) Test_Name_NewName_WithInvalidFamilyName() {
	// given
	givenName := "Franz"
	familyName := " "

	// when
	_, err := NewName(givenName, familyName)

	// then
	s.EqualError(err, ErrNameInvalidInputForFamilyName.Error())
}

package event

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/value"
)

type HomeAddressAddedTestSuite struct {
	suite.Suite
}

func Test_HomeAddressAddedTestSuite(t *testing.T) {
	tests := new(HomeAddressAddedTestSuite)
	suite.Run(t, tests)
}

func (s *HomeAddressAddedTestSuite) Test_HomeAddressWasAdded() {
	// given
	id := value.NewIdWithoutValidation("12345")
	countryCode := "DE"
	postalCode := "80803"
	city := "München"
	street := "Am Lehel"
	houseNumber := "18"
	homeAddress := value.NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber)

	// when
	homeAddressAdded := HomeAddressWasAdded(id, homeAddress)

	// then
	json, err := homeAddressAdded.ToJson()
	s.NoError(err, "it should marshall to json")
	s.NotEmpty(json, "it should expose proper json")
}

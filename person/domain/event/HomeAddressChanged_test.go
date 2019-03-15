package event

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/value"
)

type HomeAddressChangedTestSuite struct {
	suite.Suite
}

func Test_HomeAddressChangedTestSuite(t *testing.T) {
	tests := new(HomeAddressChangedTestSuite)
	suite.Run(t, tests)
}

func (s *HomeAddressChangedTestSuite) Test_HomeAddressWasChanged() {
	// given
	id := value.NewIdWithoutValidation("12345")
	countryCode := "DE"
	postalCode := "80803"
	city := "München"
	street := "Am Lehel"
	houseNumber := "18"
	homeAddress := value.NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber)

	// when
	homeAddressChanged := HomeAddressWasChanged(id, homeAddress)

	// then
	json, err := homeAddressChanged.ToJson()
	s.NoError(err, "it should marshall to json")
	s.NotEmpty(json, "it should expose proper json")
}

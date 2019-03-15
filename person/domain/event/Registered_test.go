package event

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/value"
)

type RegisteredTestSuite struct {
	suite.Suite
}

func Test_RegisteredTestSuite(t *testing.T) {
	tests := new(RegisteredTestSuite)
	suite.Run(t, tests)
}

func (s *RegisteredTestSuite) Test_HomeAddressWasAdded() {
	// given
	id := value.NewIdWithoutValidation("12345")
	givenName := "Franz"
	familyName := "Kafka"
	name := value.NewNameWithoutValidation(givenName, familyName)
	emailAddress := value.NewEmailAddressWithoutValidation("franz@kafka.de")

	// when
	registered := ItWasRegistered(id, name, emailAddress)

	// then
	json, err := registered.ToJson()
	s.NoError(err, "it should marshall to json")
	s.NotEmpty(json, "it should expose proper json")
}

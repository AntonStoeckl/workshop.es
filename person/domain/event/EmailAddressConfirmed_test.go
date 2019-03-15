package event

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/value"
)

type EmailAddressConfirmedTestSuite struct {
	suite.Suite
}

func Test_EmailAddressConfirmedTestSuite(t *testing.T) {
	tests := new(EmailAddressConfirmedTestSuite)
	suite.Run(t, tests)
}

func (s *EmailAddressConfirmedTestSuite) Test_EmailAddressWasConfirmed() {
	// given
	id := value.NewIdWithoutValidation("12345")

	// when
	emailAddressConfirmed := EmailAddressWasConfirmed(id)

	// then
	json, err := emailAddressConfirmed.ToJson()
	s.NoError(err, "it should marshall to json")
	s.NotEmpty(json, "it should expose proper json")
}

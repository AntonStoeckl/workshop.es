package value

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EmailAddressTestSuite struct {
	suite.Suite
}

func Test_EmailAddressTestSuite(t *testing.T) {
	tests := new(EmailAddressTestSuite)
	suite.Run(t, tests)
}

func (s *EmailAddressTestSuite) Test_EmailAddress_NewEmailAddress() {
	// given
	emailAddressValue := "franz@kafka.de"

	// when
	emailAddress, err := NewEmailAddress(emailAddressValue)

	// then
	s.NoError(err, "it should create an EmailAddress from valid input")
	s.Equal(emailAddressValue, emailAddress.value, "it should expose expected EmailAddress value")
}

func (s *EmailAddressTestSuite) Test_EmailAddress_NewEmailAddress_WithInvalidEmailAddress() {
	// given
	emailAddressValue := "franz@kafka"

	// when
	_, err := NewEmailAddress(emailAddressValue)

	// then
	s.EqualError(err, ErrEmailAddressInvalidInput.Error())
}

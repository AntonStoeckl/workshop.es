package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/value"
)

type PersonTestSuite struct {
	suite.Suite

	id           *value.ID
	name         *value.Name
	emailAddress *value.EmailAddress
	homeAddress  *value.Address
}

func (s *PersonTestSuite) SetupTest() {
	s.id = value.GenerateID()
	s.name = value.NewNameWithoutValidation("Franz", "Kafkae")
	s.emailAddress = value.NewEmailAddressWithoutValidation("franz@kafka.de")
	s.homeAddress = value.NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18")
}

func Test_PersonTestSuite(t *testing.T) {
	tests := new(PersonTestSuite)
	suite.Run(t, tests)
}

func (s *PersonTestSuite) Test_Person_Register() {
	// when
	recordedEvents := Register(s.id, s.name, s.emailAddress)

	// then
	s.Len(recordedEvents, 1, "it should record 1 DomainEvent")
	s.IsType(new(Registered), recordedEvents[0].Payload(), "it should record PersonRegistered")

	event := recordedEvents[0].Payload().(*Registered)
	s.Equal(event.ID, s.id.String(), "PersonRegistered should expose expected ID")
	s.Equal(event.GivenName, s.name.GivenName(), "PersonRegistered should expose expected GivenName")
	s.Equal(event.FamilyName, s.name.FamilyName(), "PersonRegistered should expose expected FamilyName")
	s.Equal(event.EmailAddress, s.emailAddress.String(), "PersonRegistered should expose expected EmailAddress")
}

func (s *PersonTestSuite) Test_Person_ConfirmEmailAddress_WhenItWasNotConfirmed() {
	// given
	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
	}

	// when
	recordedEvents := ConfirmEmailAddress(history)

	// then
	s.Len(recordedEvents, 1, "it should record 1 DomainEvent")
	s.IsType(new(EmailAddressConfirmed), recordedEvents[0].Payload(), "it should record PersonEmailAddressConfirmed")

	event := recordedEvents[0].Payload().(*EmailAddressConfirmed)
	s.Equal(event.ID, s.id.String(), "PersonEmailAddressConfirmed should expose expected ID")
}

func (s *PersonTestSuite) Test_Person_ConfirmEmailAddress_WhenItWasAlreadyConfirmed() {
	// given
	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
		EmailAddressWasConfirmed(s.id),
	}

	// when
	recordedEvents := ConfirmEmailAddress(history)

	// then
	s.Len(recordedEvents, 0, "it should NOT record DomainEvents")
}

func (s *PersonTestSuite) Test_Person_ChangeHomeAddress_WhenItWasEmpty() {
	// given
	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
	}

	// when
	recordedEvents := ChangeHomeAddress(history, s.homeAddress)

	// then
	s.Len(recordedEvents, 1, "it should record 1 DomainEvent")
	s.Equal("PersonHomeAddressAdded", recordedEvents[0].EventName(), "it should record PersonHomeAddressAdded")

	event, ok := recordedEvents[0].Payload().(*HomeAddressAdded)
	s.True(ok, "it should have a payload of type HomeAddressAdded")

	s.Equal(event.ID, s.id.String(), "PersonHomeAddressAdded  should expose expected ID")
	s.Equal(event.CountryCode, s.homeAddress.CountryCode(), "PersonHomeAddressAdded  should expose expected CountryCode")
	s.Equal(event.PostalCode, s.homeAddress.PostalCode(), "PersonHomeAddressAdded  should expose expected PostalCode")
	s.Equal(event.City, s.homeAddress.City(), "PersonHomeAddressAdded  should expose expected City")
	s.Equal(event.Street, s.homeAddress.Street(), "PersonHomeAddressAdded  should expose expected Street")
	s.Equal(event.HouseNumber, s.homeAddress.HouseNumber(), "PersonHomeAddressAdded  should expose expected HouseNumber")
}

func (s *PersonTestSuite) Test_Person_ChangeHomeAddress_WhenItWasDifferent() {
	// given
	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
		HomeAddressWasAdded(s.id, s.homeAddress),
	}

	// when
	differentHomeAddress := value.NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18b")
	recordedEvents := ChangeHomeAddress(history, differentHomeAddress)

	// then
	s.Require().Len(recordedEvents, 1, "it should record 1 DomainEvent")
	s.Require().Equal("PersonHomeAddressChanged", recordedEvents[0].EventName(), "it should record PersonHomeAddressChanged")

	event, ok := recordedEvents[0].Payload().(*HomeAddressChanged)
	s.Require().True(ok, "it should have a payload of type HomeAddressChanged")

	s.Equal(event.ID, s.id.String(), "PersonHomeAddressChanged  should expose expected ID")
	s.Equal(event.CountryCode, differentHomeAddress.CountryCode(), "PersonHomeAddressAdded  should expose expected CountryCode")
	s.Equal(event.PostalCode, differentHomeAddress.PostalCode(), "PersonHomeAddressAdded  should expose expected PostalCode")
	s.Equal(event.City, differentHomeAddress.City(), "PersonHomeAddressAdded  should expose expected City")
	s.Equal(event.Street, differentHomeAddress.Street(), "PersonHomeAddressAdded  should expose expected Street")
	s.Equal(event.HouseNumber, differentHomeAddress.HouseNumber(), "PersonHomeAddressAdded  should expose expected HouseNumber")
}

func (s *PersonTestSuite) Test_Person_ChangeHomeAddress_WhenItWasAddedEqual() {
	// given
	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
		HomeAddressWasAdded(s.id, s.homeAddress),
	}

	// when
	recordedEvents := ChangeHomeAddress(history, s.homeAddress)

	// then
	s.Require().Len(recordedEvents, 0, "it should NOT record DomainEvents")
}

func (s *PersonTestSuite) Test_Person_ChangeHomeAddress_WhenItWasChangedEqual() {
	// given
	differentHomeAddress := value.NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18b")

	history := []DomainEvent{
		ItWasRegistered(s.id, s.name, s.emailAddress),
		HomeAddressWasAdded(s.id, s.homeAddress),
		HomeAddressWasChanged(s.id, differentHomeAddress),
	}

	// when
	recordedEvents := ChangeHomeAddress(history, differentHomeAddress)

	// then
	s.Require().Len(recordedEvents, 0, "it should NOT record DomainEvents")
}

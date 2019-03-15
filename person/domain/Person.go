package domain

import (
	"workshop.es/person/domain/event"
	"workshop.es/person/domain/value"
	"workshop.es/shared"
)

type Person interface {
	ConfirmEmailAddress()
	ChangeHomeAddress(homeAddress *value.Address)
	RecordedEvents() []shared.DomainEvent
}

type person struct {
	id             *value.ID
	name           *value.Name
	emailAddress   *value.EmailAddress
	homeAddress    *value.Address
	recordedEvents []shared.DomainEvent
}

type RecordedEvents []shared.DomainEvent

func Reconstitute(history []shared.DomainEvent) *person {
	p := &person{}

	for _, domainEvent := range history {
		p.when(domainEvent)
	}

	return p
}

func Register(id *value.ID, name *value.Name, emailAddress *value.EmailAddress) *person {
	p := &person{}

	p.recordThat(event.ItWasRegistered(id, name, emailAddress))

	return p
}

func (p *person) ConfirmEmailAddress() {
	if !p.emailAddress.IsConfirmed() {
		p.recordThat(event.EmailAddressWasConfirmed(p.id))
	}
}

func (p *person) ChangeHomeAddress(homeAddress *value.Address) {
	if p.homeAddress == nil {
		p.recordThat(event.HomeAddressWasAdded(p.id, homeAddress))
	}

	if !p.homeAddress.Equals(homeAddress) {
		p.recordThat(event.HomeAddressWasChanged(p.id, homeAddress))
	}
}

func (p *person) recordThat(domainEvent shared.DomainEvent) {
	p.when(domainEvent)
	p.recordedEvents = append(p.recordedEvents, domainEvent)
}

func (p *person) when(domainEvent shared.DomainEvent) shared.DomainEvent {
	switch domainEvent.EventName() {
	case "PersonRegistered":
		p.whenItWasRegistered(domainEvent.Payload().(*event.Registered))
	case "PersonEmailAddressConfirmed":
		p.whenEmailAddressWasConfirmed()
	case "PersonHomeAddressAdded":
		p.whenHomeAddressWasAdded(domainEvent.Payload().(*event.HomeAddressAdded))
	case "PersonHomeAddressChanged":
		p.whenHomeAddressWasChanged(domainEvent.Payload().(*event.HomeAddressChanged))
	}

	return domainEvent
}

func (p *person) whenItWasRegistered(payload *event.Registered) {
	p.id = value.NewIdWithoutValidation(payload.ID)
	p.name = value.NewNameWithoutValidation(payload.GivenName, payload.FamilyName)
	p.emailAddress = value.NewEmailAddressWithoutValidation(payload.EmailAddress)
}

func (p *person) whenEmailAddressWasConfirmed() {
	p.emailAddress = p.emailAddress.Confirm()
}

func (p *person) whenHomeAddressWasAdded(payload *event.HomeAddressAdded) {
	p.homeAddress = value.NewAddressWithoutValidation(payload.CountryCode, payload.PostalCode, payload.City, payload.Street, payload.HouseNumber)
}

func (p *person) whenHomeAddressWasChanged(payload *event.HomeAddressChanged) {
	p.homeAddress = value.NewAddressWithoutValidation(payload.CountryCode, payload.PostalCode, payload.City, payload.Street, payload.HouseNumber)
}

func (p *person) RecordedEvents() []shared.DomainEvent {
	return p.recordedEvents
}

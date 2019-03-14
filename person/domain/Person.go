package domain

import "workshop.es/person/domain/value"

type person struct {
	id             *value.ID
	name           *value.Name
	emailAddress   *value.EmailAddress
	homeAddress    *value.Address
	recordedEvents []DomainEvent
}

type RecordedEvents []DomainEvent

func Register(id *value.ID, name *value.Name, emailAddress *value.EmailAddress) RecordedEvents {
	p := &person{}

	p.recordThat(ItWasRegistered(id, name, emailAddress))

	return p.recordedEvents
}

func ConfirmEmailAddress(history []DomainEvent) RecordedEvents {
	p := reconstitute(history)

	if !p.emailAddress.IsConfirmed() {
		p.recordThat(EmailAddressWasConfirmed(p.id))
	}

	return p.recordedEvents
}

func ChangeHomeAddress(history []DomainEvent, homeAddress *value.Address) RecordedEvents {
	p := reconstitute(history)

	switch true {
	case p.homeAddress == nil:
		p.recordThat(HomeAddressWasAdded(p.id, homeAddress))
	case !p.homeAddress.Equals(homeAddress):
		p.recordThat(HomeAddressWasChanged(p.id, homeAddress))
	}

	return p.recordedEvents
}

func reconstitute(history []DomainEvent) *person {
	p := &person{}

	for _, domainEvent := range history {
		p.when(domainEvent)
	}

	return p
}

func (p *person) recordThat(domainEvent DomainEvent) {
	p.when(domainEvent)
	p.recordedEvents = append(p.recordedEvents, domainEvent)
}

func (p *person) when(domainEvent DomainEvent) DomainEvent {
	switch domainEvent.EventName() {
	case "PersonRegistered":
		p.whenItWasRegistered(domainEvent.Payload().(*Registered))
	case "PersonEmailAddressConfirmed":
		p.whenEmailAddressWasConfirmed()
	case "PersonHomeAddressAdded":
		p.whenHomeAddressWasAdded(domainEvent.Payload().(*HomeAddressAdded))
	case "PersonHomeAddressChanged":
		p.whenHomeAddressWasChanged(domainEvent.Payload().(*HomeAddressChanged))
	}

	return domainEvent
}

func (p *person) whenItWasRegistered(payload *Registered) {
	p.id = value.NewIdWithoutValidation(payload.ID)
	p.name = value.NewNameWithoutValidation(payload.GivenName, payload.FamilyName)
	p.emailAddress = value.NewEmailAddressWithoutValidation(payload.EmailAddress)
}

func (p *person) whenEmailAddressWasConfirmed() {
	p.emailAddress = p.emailAddress.Confirm()
}

func (p *person) whenHomeAddressWasAdded(payload *HomeAddressAdded) {
	p.homeAddress = value.NewAddressWithoutValidation(payload.CountryCode, payload.PostalCode, payload.City, payload.Street, payload.HouseNumber)
}

func (p *person) whenHomeAddressWasChanged(payload *HomeAddressChanged) {
	p.homeAddress = value.NewAddressWithoutValidation(payload.CountryCode, payload.PostalCode, payload.City, payload.Street, payload.HouseNumber)
}

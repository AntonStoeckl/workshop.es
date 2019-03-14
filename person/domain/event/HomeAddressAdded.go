package event

import "workshop.es/person/domain/value"

type HomeAddressAdded struct {
	ID          string
	CountryCode string
	PostalCode  string
	City        string
	Street      string
	HouseNumber string
}

func HomeAddressWasAdded(id *value.ID, homeAddress *value.Address) *domainEvent {
	payload := &HomeAddressAdded{
		ID:          id.String(),
		CountryCode: homeAddress.CountryCode(),
		PostalCode:  homeAddress.PostalCode(),
		City:        homeAddress.City(),
		Street:      homeAddress.Street(),
		HouseNumber: homeAddress.HouseNumber(),
	}

	return NewDomainEventFromPayload(payload)
}

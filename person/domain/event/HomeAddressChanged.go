package event

import "workshop.es/person/domain/value"

type HomeAddressChanged struct {
	ID          string
	CountryCode string
	PostalCode  string
	City        string
	Street      string
	HouseNumber string
}

func HomeAddressWasChanged(id *value.ID, homeAddress *value.Address) *domainEvent {
	payload := &HomeAddressChanged{
		ID:          id.String(),
		CountryCode: homeAddress.CountryCode(),
		PostalCode:  homeAddress.PostalCode(),
		City:        homeAddress.City(),
		Street:      homeAddress.Street(),
		HouseNumber: homeAddress.HouseNumber(),
	}

	return NewDomainEventFromPayload(payload)
}

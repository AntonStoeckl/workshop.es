package event

import (
	"workshop.es/person/domain/value"
	"workshop.es/shared"
)

type HomeAddressChanged struct {
	ID          string `json:"id"`
	CountryCode string `json:"countryCode"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
}

func HomeAddressWasChanged(id *value.ID, homeAddress *value.Address) *shared.DomainEventData {
	payload := &HomeAddressChanged{
		ID:          id.String(),
		CountryCode: homeAddress.CountryCode(),
		PostalCode:  homeAddress.PostalCode(),
		City:        homeAddress.City(),
		Street:      homeAddress.Street(),
		HouseNumber: homeAddress.HouseNumber(),
	}

	return shared.NewDomainEventFromPayload(id, payload)
}

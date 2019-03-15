package event

import (
	"workshop.es/person/domain/value"
	"workshop.es/shared"
)

type Registered struct {
	ID           string `json:"id"`
	GivenName    string `json:"givenName"`
	FamilyName   string `json:"familyName"`
	EmailAddress string `json:"emailAddress"`
}

func ItWasRegistered(id *value.ID, name *value.Name, emailAddress *value.EmailAddress) *shared.DomainEventData {
	payload := &Registered{
		ID:           id.String(),
		GivenName:    name.GivenName(),
		FamilyName:   name.FamilyName(),
		EmailAddress: emailAddress.String(),
	}

	return shared.NewDomainEventFromPayload(id, payload)
}

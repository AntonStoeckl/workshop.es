package event

import "workshop.es/person/domain/value"

type Registered struct {
	ID           string
	GivenName    string
	FamilyName   string
	EmailAddress string
}

func ItWasRegistered(id *value.ID, name *value.Name, emailAddress *value.EmailAddress) *domainEvent {
	payload := &Registered{
		ID:           id.String(),
		GivenName:    name.GivenName(),
		FamilyName:   name.FamilyName(),
		EmailAddress: emailAddress.String(),
	}

	return NewDomainEventFromPayload(payload)
}

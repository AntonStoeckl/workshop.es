package event

import (
	"workshop.es/person/domain/value"
	"workshop.es/shared"
)

type EmailAddressConfirmed struct {
	ID string `json:"id"`
}

func EmailAddressWasConfirmed(id *value.ID) *shared.DomainEventData {
	payload := &EmailAddressConfirmed{
		ID: id.String(),
	}

	return shared.NewDomainEventFromPayload(id, payload)
}

package domain

import "workshop.es/person/domain/value"

type EmailAddressConfirmed struct {
	ID string
}

func EmailAddressWasConfirmed(id *value.ID) *domainEvent {
	payload := &EmailAddressConfirmed{
		ID: id.String(),
	}

	return NewDomainEventFromPayload(payload)
}

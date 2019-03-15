package event

import (
	"encoding/json"

	"workshop.es/shared"
)

type domainEventUnmarshaller struct {
}

func NewDomainEventUnmarshaller() *domainEventUnmarshaller {
	return &domainEventUnmarshaller{}
}

func (unmarshaller *domainEventUnmarshaller) UnmarshalFromJSON(eventJSON string, eventName string) (shared.DomainEvent, error) {
	var domainEvent shared.DomainEventData

	if err := json.Unmarshal([]byte(eventJSON), &domainEvent); err != nil {
		return nil, err
	}

	return &domainEvent, nil
}

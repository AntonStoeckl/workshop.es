package shared

type DomainEventUnmarshaller interface {
	UnmarshalFromJSON(eventJSON string, eventName string) (DomainEvent, error)
}

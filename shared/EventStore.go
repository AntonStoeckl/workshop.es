package shared

type EventStore interface {
	Append(recordedEvents []DomainEvent, currentVersion uint) error
	StreamFor(id StreamID) ([]DomainEvent, uint, error)
}

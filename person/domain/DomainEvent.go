package domain

import (
	"reflect"
	"strings"
	"time"
)

const DomainEventTimestampFormat = "2006-01-02T15:04:05.000Z07:00" // RFC3339Milli

type DomainEvent interface {
	EventName() string
	OccurredAt() string
	Payload() interface{}
}

type domainEvent struct {
	meta    *domainEventMeta
	payload interface{}
}

func NewDomainEventFromPayload(payload interface{}) *domainEvent {
	domainEvent := &domainEvent{
		meta:    NewDomainEventMeta(domainEventNameFromPayload(payload)),
		payload: payload,
	}

	return domainEvent
}

func domainEventNameFromPayload(payload interface{}) string {
	eventPayloadType := reflect.TypeOf(payload).String()
	eventPlayloadTypeParts := strings.Split(eventPayloadType, ".")
	domainEventName := "Person" + eventPlayloadTypeParts[len(eventPlayloadTypeParts)-1]

	return domainEventName
}

func (event *domainEvent) EventName() string {
	return event.meta.eventName
}

func (event *domainEvent) OccurredAt() string {
	return event.meta.occurredAt
}

func (event *domainEvent) Payload() interface{} {
	return event.payload
}

type domainEventMeta struct {
	eventName  string
	occurredAt string
}

func NewDomainEventMeta(name string) *domainEventMeta {
	return &domainEventMeta{
		eventName:  name,
		occurredAt: time.Now().In(time.UTC).Format(DomainEventTimestampFormat),
	}
}

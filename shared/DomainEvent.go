package shared

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type DomainEvent interface {
	StreamID() string
	EventName() string
	OccurredAt() string
	Payload() interface{}
	ToJson() (string, error)
}

const DomainEventTimestampFormat = "2006-01-02T15:04:05.000Z07:00" // RFC3339Milli

type DomainEventData struct {
	id          StreamID
	Meta        *DomainEventMeta `json:"meta"`
	PayloadData interface{}      `json:"payload"`
}

func NewDomainEventFromPayload(streamID StreamID, payload interface{}) *DomainEventData {
	domainEvent := &DomainEventData{
		id:          streamID,
		Meta:        NewDomainEventMeta(domainEventNameFromPayload(payload)),
		PayloadData: payload,
	}

	return domainEvent
}

func domainEventNameFromPayload(payload interface{}) string {
	eventPayloadType := reflect.TypeOf(payload).String()
	eventPlayloadTypeParts := strings.Split(eventPayloadType, ".")
	domainEventName := "Person" + eventPlayloadTypeParts[len(eventPlayloadTypeParts)-1]

	return domainEventName
}

func (event *DomainEventData) StreamID() string {
	return event.id.String()
}

func (event *DomainEventData) EventName() string {
	return event.Meta.EventName
}

func (event *DomainEventData) OccurredAt() string {
	return event.Meta.OccurredAt
}

func (event *DomainEventData) Payload() interface{} {
	return event.PayloadData
}

func (event *DomainEventData) ToJson() (string, error) {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

type DomainEventMeta struct {
	EventName  string `json:"eventName"`
	OccurredAt string `json:"occurredAt"`
}

func NewDomainEventMeta(name string) *DomainEventMeta {
	return &DomainEventMeta{
		EventName:  name,
		OccurredAt: time.Now().In(time.UTC).Format(DomainEventTimestampFormat),
	}
}

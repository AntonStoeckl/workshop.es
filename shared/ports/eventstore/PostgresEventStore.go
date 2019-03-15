package eventstore

import (
	"database/sql"
	"strings"

	"workshop.es/shared"
)

type postgresEventStore struct {
	db           *sql.DB
	tableName    string
	unmarshaller shared.DomainEventUnmarshaller
}

func NewPostgresEventStore(db *sql.DB, tableName string, unmarshaller shared.DomainEventUnmarshaller) *postgresEventStore {
	return &postgresEventStore{db: db, tableName: tableName, unmarshaller: unmarshaller}
}

func (store *postgresEventStore) Append(recordedEvents []shared.DomainEvent, currentVersion uint) error {
	queryTemplate := `INSERT INTO %name% (event_name, stream_id, payload, occurred_at, stream_version) VALUES ($1, $2, $3, $4, $5)`
	query := strings.Replace(queryTemplate, "%name%", store.tableName, 1)

	for _, event := range recordedEvents {
		eventJson, err := event.ToJson()
		if err != nil {
			return err
		}

		currentVersion++

		_, err = store.db.Exec(
			query,
			event.EventName(),
			event.StreamID(),
			[]byte(eventJson),
			event.OccurredAt(),
			currentVersion,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (store *postgresEventStore) StreamFor(id shared.StreamID) ([]shared.DomainEvent, uint, error) {
	queryTemplate := `SELECT event_name, payload, stream_version FROM %name% WHERE stream_id = $1 ORDER BY stream_version`
	query := strings.Replace(queryTemplate, "%name%", store.tableName, 1)

	eventRows, err := store.db.Query(query, id.String())
	if err != nil {
		return nil, 0, err
	}

	var stream []shared.DomainEvent
	var eventName string
	var payload string
	var streamVersion int
	var domainEvent shared.DomainEvent

	for eventRows.Next() {
		if err = eventRows.Scan(&eventName, &payload, &streamVersion); err == nil {
			if domainEvent, err = store.unmarshaller.UnmarshalFromJSON(payload, eventName); err != nil {
				return nil, 0, err
			}

			stream = append(stream, domainEvent)
		} else {
			return nil, 0, err
		}
	}

	return stream, uint(streamVersion), nil
}

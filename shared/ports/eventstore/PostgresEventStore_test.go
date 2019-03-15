package eventstore

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"workshop.es/person/domain/event"
	"workshop.es/person/domain/value"
	"workshop.es/shared"
)

type PostgresEventStoreTestSuite struct {
	suite.Suite
}

func Test_PostgresEventStoreTestSuite(t *testing.T) {
	tests := new(PostgresEventStoreTestSuite)
	suite.Run(t, tests)
}

func (s *PostgresEventStoreTestSuite) SetupTest() {

}

func (s *PostgresEventStoreTestSuite) Test_Append() {
	// given
	id := value.NewIdWithoutValidation("12345")
	givenName := "Franz"
	familyName := "Kafka"
	name := value.NewNameWithoutValidation(givenName, familyName)
	emailAddress := value.NewEmailAddressWithoutValidation("franz@kafka.de")
	countryCode := "DE"
	postalCode := "80803"
	city := "München"
	street := "Am Lehel"
	houseNumber1 := "18"
	houseNumber2 := "20"
	homeAddress1 := value.NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber1)
	homeAddress2 := value.NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber2)

	db, err := sql.Open("postgres", "postgresql://esworkshop:password123@localhost:5432/esworkshop?sslmode=disable")
	if err != nil {
		log.Fatalf("Database could not be opened: %v\n", err)
	}

	eventStore := NewPostgresEventStore(db, "eventstore", event.NewDomainEventUnmarshaller())

	recordedEvents := []shared.DomainEvent{
		event.ItWasRegistered(id, name, emailAddress),
		event.EmailAddressWasConfirmed(id),
		event.HomeAddressWasAdded(id, homeAddress1),
		event.HomeAddressWasChanged(id, homeAddress2),
	}

	// when
	err = eventStore.Append(recordedEvents, 0)
	s.NoError(err, "it should Append")

	stream, currentVersion, err := eventStore.StreamFor(id)
	s.NoError(err, "it should StreamFor")

	_ = stream
	_ = currentVersion
}

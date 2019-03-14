package value

import (
	"github.com/google/uuid"
)

type ID struct {
	id string
}

func GenerateID() *ID {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic("could not generate uuid: " + err.Error())
	}

	return &ID{id: uid.String()}
}

func NewIdWithoutValidation(id string) *ID {
	return &ID{id: id}
}

func (id *ID) String() string {
	return id.id
}

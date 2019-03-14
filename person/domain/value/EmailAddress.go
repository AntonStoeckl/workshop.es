package value

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmailAddressInvalidInput = errors.New("emailaddress: invalid input given")
	emailAddressRegExp          = regexp.MustCompile(`^[^\s]+@[^\s]+\.[\w]{2,}$`)
)

type EmailAddress struct {
	value       string
	isConfirmed bool
}

func NewEmailAddress(emailAddress string) (*EmailAddress, error) {
	newEmailAddress := NewEmailAddressWithoutValidation(emailAddress)

	if err := newEmailAddress.validate(); err != nil {
		return nil, err
	}

	return newEmailAddress, nil
}

func NewEmailAddressWithoutValidation(emailAddress string) *EmailAddress {
	newEmailAddress := &EmailAddress{
		value: strings.TrimSpace(emailAddress),
	}

	return newEmailAddress
}

func (emailAddress *EmailAddress) Confirm() *EmailAddress {
	return &EmailAddress{
		value:       emailAddress.value,
		isConfirmed: true,
	}
}

func (emailAddress *EmailAddress) String() string {
	return emailAddress.value
}

func (emailAddress *EmailAddress) IsConfirmed() bool {
	return emailAddress.isConfirmed
}

func (emailAddress *EmailAddress) validate() error {
	if matched := emailAddressRegExp.MatchString(emailAddress.value); matched != true {
		return ErrEmailAddressInvalidInput
	}

	return nil
}

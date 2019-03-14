package value

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrAddressInvalidInputForCountryCode = errors.New("address: invalid input given for countrycode")
	ErrAddressInvalidInputForPostalCode  = errors.New("address: invalid input given for postalcode")
	ErrAddressInvalidInputForCity        = errors.New("address: invalid input given for city")
	ErrAddressInvalidInputForStreet      = errors.New("address: invalid input given for street")
	ErrAddressInvalidInputForHouseNumber = errors.New("address: invalid input given for housenumber")
)

type Address struct {
	countryCode string
	postalCode  string
	city        string
	street      string
	houseNumber string
}

func NewAddress(countryCode, postalCode, city, street, houseNumber string) (*Address, error) {
	newAddress := NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber)

	if err := newAddress.validate(); err != nil {
		return nil, err
	}

	return newAddress, nil
}

func NewAddressWithoutValidation(countryCode, postalCode, city, street, houseNumber string) *Address {
	newAddress := &Address{
		countryCode: strings.TrimSpace(countryCode),
		postalCode:  strings.TrimSpace(postalCode),
		city:        strings.TrimSpace(city),
		street:      strings.TrimSpace(street),
		houseNumber: strings.TrimSpace(houseNumber),
	}

	return newAddress
}

func (address *Address) CountryCode() string {
	return address.countryCode
}

func (address *Address) PostalCode() string {
	return address.postalCode
}

func (address *Address) City() string {
	return address.city
}

func (address *Address) Street() string {
	return address.street
}

func (address *Address) HouseNumber() string {
	return address.houseNumber
}

func (address *Address) Equals(other *Address) bool {
	return reflect.DeepEqual(address, other)
}

func (address *Address) validate() error {
	if address.countryCode == "" {
		return ErrAddressInvalidInputForCountryCode
	}

	if address.postalCode == "" {
		return ErrAddressInvalidInputForPostalCode
	}

	if address.city == "" {
		return ErrAddressInvalidInputForCity
	}

	if address.street == "" {
		return ErrAddressInvalidInputForStreet
	}

	if address.houseNumber == "" {
		return ErrAddressInvalidInputForHouseNumber
	}

	return nil
}

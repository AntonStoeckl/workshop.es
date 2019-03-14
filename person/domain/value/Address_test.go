package value

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AddressTestSuite struct {
	suite.Suite
}

func Test_AddressTestSuite(t *testing.T) {
	tests := new(AddressTestSuite)
	suite.Run(t, tests)
}

func (s *AddressTestSuite) Test_Address_NewAddress() {
	// given
	countryCode := "DE"
	postalCode := "80803"
	city := "München"
	street := "Am Lehel"
	houseNumber := "18"

	// when
	address, err := NewAddress(countryCode, postalCode, city, street, houseNumber)

	// then
	s.NoError(err, "it should create an Address from valid input")
	s.Equal(countryCode, address.countryCode)
	s.Equal(postalCode, address.postalCode)
	s.Equal(city, address.city)
	s.Equal(street, address.street)
	s.Equal(houseNumber, address.houseNumber)
}

func (s *AddressTestSuite) Test_Address_NewAddress_WithInvalidCountryCode() {
	// given
	countryCode := " "
	postalCode := "80803"
	city := "München"
	street := "Am Lehel"
	houseNumber := "18"

	// when
	_, err := NewAddress(countryCode, postalCode, city, street, houseNumber)

	// then
	s.EqualError(err, ErrAddressInvalidInputForCountryCode.Error())
}

func (s *AddressTestSuite) Test_Address_NewAddress_Equals() {
	// given
	address := NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18")
	equalAddress := NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18")

	// when / then
	s.True(address.Equals(equalAddress), "Addresses should be equal")
}

func (s *AddressTestSuite) Test_Address_NewAddress_Not_Equals() {
	// given
	address := NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18a")
	equalAddress := NewAddressWithoutValidation("DE", "80803", "München", "Am Lehel", "18b")

	// when / then
	s.False(address.Equals(equalAddress), "Addresses should NOT be equal")
}

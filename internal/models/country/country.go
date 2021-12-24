package country

import (
	"strings"

	"github.com/jolinGalal/jumia/internal/models/country/factory"
	"github.com/jolinGalal/jumia/internal/models/country/types"
)

// CountryFacade ...
type CountryFacade struct {
	countryFactory
}

// New ...
func New(phone *string) CountryI {

	if strings.HasPrefix(*phone, types.CountryCodes[types.CountryLst[0]]) {
		return &CountryFacade{factory.NewCameroon(*phone)}
	}

	if strings.HasPrefix(*phone, types.CountryCodes[types.CountryLst[1]]) {
		return &CountryFacade{factory.NewEthiopia(*phone)}
	}

	if strings.HasPrefix(*phone, types.CountryCodes[types.CountryLst[2]]) {
		return &CountryFacade{factory.NewMorocco(*phone)}
	}

	if strings.HasPrefix(*phone, types.CountryCodes[types.CountryLst[3]]) {
		return &CountryFacade{factory.NewMozambique(*phone)}
	}

	if strings.HasPrefix(*phone, types.CountryCodes[types.CountryLst[4]]) {
		return &CountryFacade{factory.NewUganda(*phone)}
	}
	return &CountryFacade{}
}

// Name ...
func (c *CountryFacade) Name() string {
	if c.countryFactory != nil {
		return c.countryFactory.Name()
	}
	return ""
}

//State ....
func (c *CountryFacade) State() string {
	if c.countryFactory != nil {
		if c.countryFactory.State() {
			return types.Statelst[0]
		}
	}
	return types.Statelst[1]
}

// Code ...
func (c *CountryFacade) Code() string {
	if c.countryFactory != nil {
		return c.countryFactory.Code()
	}
	return ""
}

// CountryCodeByName ...
func (c *CountryFacade) CountryCodeByName(countryStr string) string {
	if code, ok := types.CountryCodes[countryStr]; !ok {
		return ""
	} else {
		return code
	}
}

// CountryLst ...
func (c *CountryFacade) GetCountries() []string {
	return types.CountryLst
}

// GetStates ...
func (c *CountryFacade) GetStates() []string {
	return types.Statelst
}

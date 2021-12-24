package country

// countryI ...
type countryFactory interface {
	Name() string
	State() bool
	Code() string
}

// CountryI ...
type CountryI interface {
	Name() string
	State() string
	Code() string
	CountryCodeByName(countryStr string) string
	GetCountries() []string
	GetStates() []string
}

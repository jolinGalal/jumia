package country

import "regexp"

// CountryI ...
type CountryI interface {
	Name() string
	State() bool
}

// New ...
func New(phone *string) CountryI {

	var re = regexp.MustCompile(`\(237\)\ [0-9]`)
	if re.MatchString(*phone) {
		return newCameroon(*phone)
	}
	re = regexp.MustCompile(`\(251\)\ [0-9]`)
	if re.MatchString(*phone) {
		return newEthiopia(*phone)
	}
	re = regexp.MustCompile(`\(212\)\ [0-9]`)
	if re.MatchString(*phone) {
		return newMorocco(*phone)
	}
	re = regexp.MustCompile(`\(258\)\ [0-9]`)
	if re.MatchString(*phone) {
		return newMozambique(*phone)
	}
	re = regexp.MustCompile(`\(256\)\ [0-9]`)
	if re.MatchString(*phone) {
		return newUganda(*phone)
	}
	return nil
}

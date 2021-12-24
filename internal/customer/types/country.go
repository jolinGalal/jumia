package types

import "github.com/jolinGalal/jumia/internal/models/country"

var phone = ""
var CountryModel = country.New(&phone)
var Countrylst = CountryModel.GetCountries()

package types

var Statelst = CountryModel.GetStates()

// State ...
var State = struct {
	Valid    string
	NotValid string
	All      string
}{
	Valid:    Statelst[0],
	NotValid: Statelst[1],
	All:      "all",
}

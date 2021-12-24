package validation

var phone = ""
var v = New(&phone)

// cameronRegularExpression ...
var cameronRegularExpression = func() {
	v.Code("(237)")
	v.Optional(' ')
	v.OneOf('2', '3', '6', '8')
	v.Digits(7, 8)
}

// EthiopiaRegularExpression ...
var EthiopiaRegularExpression = func() {
	v.Code("(251)")
	v.Optional(' ')
	v.OneOf('1', '2', '3', '4', '5', '9')
	v.Digits(8)
}

// MoroccoRegularExpression ...
var MoroccoRegularExpression = func() {
	v.Code("(212)")
	v.Optional(' ')
	v.OneOf('5', '6', '7', '8', '9')
	v.Digits(8)
}

// MozambiqueRegularExpression ...
var MozambiqueRegularExpression = func() {
	v.Code("(258)")
	v.Optional(' ')
	v.OneOf('2', '8')
	v.Digits(7, 8)
}

// UgandaRegularExpression ...
var UgandaRegularExpression = func() {
	v.Code("(256)")
	v.Optional(' ')
	v.Digits(9)
}

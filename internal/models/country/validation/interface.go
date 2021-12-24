package validation

// ValidationI ...
type ValidationI interface {
	Valid(regularExpression func()) bool
	Code(code string)
	OneOf(chars ...rune)
	Optional(chars ...rune)
	Digits(numberOfDigits ...int)
}

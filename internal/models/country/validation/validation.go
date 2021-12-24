package validation

import (
	"sort"
	"strconv"
	"strings"
)

// validation ...
type validation struct {
	phone *string
	index int
	valid bool
}

// New ...
func New(phone *string) ValidationI {
	return &validation{phone: phone, index: 0, valid: true}
}

// Valid ...
func (v *validation) Valid(regularExpression func()) bool {
	if v.valid {
		regularExpression()
	}
	if v.index != len(*v.phone) {
		v.valid = false
	}
	return v.valid
}

// Code ...
func (v *validation) Code(code string) {
	if v.valid {
		if strings.HasPrefix(*v.phone, code) {
			v.index = v.index + len(code)
			return
		}
	}
	v.valid = false
}

// OneOf ...
func (v *validation) OneOf(chars ...rune) {
	if v.valid {
		for _, char := range chars {
			if v.index >= len(*v.phone) {
				break
			}
			if string(char) == string((*v.phone)[v.index]) {
				v.index++
				return
			}
		}
		// incase the the index is not found
		// incase the index out of range
		v.valid = false
	}
}

// Optional ...
func (v *validation) Optional(chars ...rune) {
	if v.valid {
		for _, char := range chars {
			if v.index < len(*v.phone) {
				if string(char) == string((*v.phone)[v.index]) {
					v.index++
				}
			} else {
				break
			}
		}
	}

}

// Digits ...
func (v *validation) Digits(numberOfDigits ...int) {
	var currentIndex = v.index
	sort.Sort(sort.Reverse(sort.IntSlice(numberOfDigits)))
	if v.valid {
		for _, number := range numberOfDigits {
			v.index = currentIndex
			for count := 1; count <= number; count++ {
				if v.index >= len(*v.phone) {
					break
				}
				// not digit
				if _, err := strconv.Atoi(string((*v.phone)[v.index])); err != nil {
					break
				}

				v.index++
				// the number of digits
				if count == number {
					return
				}

			}
		}
		v.valid = false

	}

}

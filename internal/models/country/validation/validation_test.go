package validation

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

//checkValidation
func checkValidation(t *testing.T, re *regexp.Regexp, f func()) {
	var valid = re.MatchString(phone)
	v = New(&phone)
	var validFunc = v.Valid(f)
	assert.Equal(t, valid, validFunc)
}

// TestCameronValidation ...
func TestCameronValidation(t *testing.T) {

	// test cameron regular express
	var re = regexp.MustCompile(`\(237\)\ ?[2368]\d{7,8}$`)
	//case valid
	phone = "(237) 678000952"
	checkValidation(t, re, cameronRegularExpression)

	// case not valid
	phone = "(237) 6780009a2"
	checkValidation(t, re, cameronRegularExpression)

	// case long length
	phone = "(237) 67800092345"
	checkValidation(t, re, cameronRegularExpression)

	// case short length
	phone = "(237) 67800"
	checkValidation(t, re, cameronRegularExpression)

	// case empty
	phone = ""
	checkValidation(t, re, cameronRegularExpression)

}

// TestEthiopiaValidation ...
func TestEthiopiaValidation(t *testing.T) {
	// test Ethiopia regular express
	var re = regexp.MustCompile(`\(251\)\ ?[1-59]\d{8}$`)

	//case valid
	phone = "(251) 411168450"
	checkValidation(t, re, EthiopiaRegularExpression)

	// case not valid
	phone = "(251) 711168450"
	checkValidation(t, re, EthiopiaRegularExpression)

	// case long length
	phone = "(251) 71116845074829374"
	checkValidation(t, re, EthiopiaRegularExpression)

	// case short length
	phone = "(251) 7111"
	checkValidation(t, re, EthiopiaRegularExpression)

	// case empty
	phone = ""
	checkValidation(t, re, EthiopiaRegularExpression)

}

// TestMorocoValidation ...
func TestMorocoValidation(t *testing.T) {
	// test Ethiopia regular express
	var re = regexp.MustCompile(`\(212\)\ ?[5-9]\d{8}$`)

	//case valid
	phone = "(212) 698054317"
	checkValidation(t, re, MoroccoRegularExpression)

	// case not valid
	phone = "(212) 198054317"
	checkValidation(t, re, MoroccoRegularExpression)

	// case long length
	phone = "(212) 69805431738798734"
	checkValidation(t, re, MoroccoRegularExpression)

	// case short length
	phone = "(212) 698054"
	checkValidation(t, re, MoroccoRegularExpression)

	// case empty
	phone = ""
	checkValidation(t, re, MoroccoRegularExpression)

}

// TestMozambiqueValidation ...
func TestMozambiqueValidation(t *testing.T) {
	// test Ethiopia regular express
	var re = regexp.MustCompile(`\(258\)\ ?[28]\d{7,8}$`)

	//case valid
	phone = "(258) 84765150"
	checkValidation(t, re, MozambiqueRegularExpression)

	// case not valid
	phone = "(258) 747651504"
	checkValidation(t, re, MozambiqueRegularExpression)

	// case long length
	phone = "(258) 8476515042344"
	checkValidation(t, re, MozambiqueRegularExpression)

	// case short length
	phone = "(258) 84765"
	checkValidation(t, re, MozambiqueRegularExpression)

	// case empty
	phone = ""
	checkValidation(t, re, MozambiqueRegularExpression)

}

// TestEthiopiaValidation ...
func TestUgandaValidation(t *testing.T) {
	// test Ethiopia regular express
	var re = regexp.MustCompile(`\(256\)\ ?\d{9}$`)

	//case valid
	phone = "(256) 775069443"
	checkValidation(t, re, UgandaRegularExpression)

	// case not valid
	phone = "(256) 7750j9443"
	checkValidation(t, re, UgandaRegularExpression)

	// case long length
	phone = "(256) 775069443324"
	checkValidation(t, re, UgandaRegularExpression)

	// case short length
	phone = "(256) 7750694"
	checkValidation(t, re, UgandaRegularExpression)

	// case empty
	phone = ""
	checkValidation(t, re, UgandaRegularExpression)

}

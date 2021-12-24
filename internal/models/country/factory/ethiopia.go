package factory

import (
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country/types"
	"github.com/jolinGalal/jumia/internal/models/country/validation"
)

// ethiopia ...
type ethiopia struct {
	validObj validation.ValidationI
}

// ethiopiaLock ...
var ethiopiaLock = &sync.Mutex{}

// ethiopiaInstance ...
var ethiopiaInstance *ethiopia

// NewEthiopia ...
func NewEthiopia(phone string) *ethiopia {
	if ethiopiaInstance == nil {
		ethiopiaLock.Lock()
		defer ethiopiaLock.Unlock()
		if ethiopiaInstance == nil {
			ethiopiaInstance = &ethiopia{}
		}
	}
	ethiopiaInstance.validObj = validation.New(&phone)

	return ethiopiaInstance
}

// Name ...
func (e *ethiopia) Name() string {
	return types.CountryLst[1]
}

// State ...
func (e *ethiopia) State() bool {
	if e.validObj != nil {
		return e.validObj.Valid(func() {
			e.validObj.Code("(251)")
			e.validObj.Optional(' ')
			e.validObj.OneOf('1', '2', '3', '4', '5', '9')
			e.validObj.Digits(8)
		})
	}
	return false
}

// Code ...
func (e *ethiopia) Code() string {
	if code, ok := types.CountryCodes[types.CountryLst[1]]; ok {
		return code
	}
	return ""
}

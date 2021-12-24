package factory

import (
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country/types"
	"github.com/jolinGalal/jumia/internal/models/country/validation"
)

// cameroon ...
type cameroon struct {
	validObj validation.ValidationI
}

// cameroonLock ...
var cameroonLock = &sync.Mutex{}

// cameroonInstance ...
var cameroonInstance *cameroon

// NewCameroon ...
func NewCameroon(phone string) *cameroon {
	if cameroonInstance == nil {
		cameroonLock.Lock()
		defer cameroonLock.Unlock()
		if cameroonInstance == nil {
			cameroonInstance = &cameroon{}
		}
	}
	cameroonInstance.validObj = validation.New(&phone)
	return cameroonInstance
}

// Name ...
func (c *cameroon) Name() string {
	return types.CountryLst[0]
}

// State ...
func (c *cameroon) State() bool {
	return c.validObj.Valid(func() {
		c.validObj.Code("(237)")
		c.validObj.Optional(' ')
		c.validObj.OneOf('2', '3', '6', '8')
		c.validObj.Digits(7, 8)
	})
}

// Code ...
func (c *cameroon) Code() string {
	if code, ok := types.CountryCodes[types.CountryLst[0]]; ok {
		return code
	}
	return ""
}

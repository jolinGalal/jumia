package factory

import (
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country/types"
	"github.com/jolinGalal/jumia/internal/models/country/validation"
)

// uganda ...
type uganda struct {
	validObj validation.ValidationI
}

// ugandaLock ...
var ugandaLock = &sync.Mutex{}

// ugandaInstance ...
var ugandaInstance *uganda

// NewUganda ...
func NewUganda(phone string) *uganda {
	if ugandaInstance == nil {
		ugandaLock.Lock()
		defer ugandaLock.Unlock()
		if ugandaInstance == nil {
			ugandaInstance = &uganda{}
		}
	}
	ugandaInstance.validObj = validation.New(&phone)
	return ugandaInstance
}

//Name ...
func (u *uganda) Name() string {
	return types.CountryLst[4]
}

// State ...
func (u *uganda) State() bool {
	if u.validObj != nil {
		return u.validObj.Valid(func() {
			u.validObj.Code("(256)")
			u.validObj.Optional(' ')
			u.validObj.Digits(9)
		})
	}
	return false
}

// Code ...
func (u *uganda) Code() string {
	if code, ok := types.CountryCodes[types.CountryLst[4]]; ok {
		return code
	}
	return ""
}

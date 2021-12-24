package factory

import (
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country/types"
	"github.com/jolinGalal/jumia/internal/models/country/validation"
)

// mozambique ...
type mozambique struct {
	validObj validation.ValidationI
}

// mozambiqueLock ...
var mozambiqueLock = &sync.Mutex{}

// mozambiqueInstance ...
var mozambiqueInstance *mozambique

// NewMozambique ...
func NewMozambique(phone string) *mozambique {
	if mozambiqueInstance == nil {
		mozambiqueLock.Lock()
		defer mozambiqueLock.Unlock()
		if mozambiqueInstance == nil {
			mozambiqueInstance = &mozambique{}
		}
	}

	mozambiqueInstance.validObj = validation.New(&phone)
	return mozambiqueInstance
}

// Name ...
func (m *mozambique) Name() string {
	return types.CountryLst[3]
}

// State ...
func (m *mozambique) State() bool {
	if m.validObj != nil {
		return m.validObj.Valid(func() {
			m.validObj.Code("(258)")
			m.validObj.Optional(' ')
			m.validObj.OneOf('2', '8')
			m.validObj.Digits(7, 8)
		})
	}
	return false
}

// Code ...
func (m *mozambique) Code() string {
	if code, ok := types.CountryCodes[types.CountryLst[3]]; ok {
		return code
	}
	return ""
}

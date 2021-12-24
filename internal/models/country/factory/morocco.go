package factory

import (
	"sync"

	"github.com/jolinGalal/jumia/internal/models/country/types"
	"github.com/jolinGalal/jumia/internal/models/country/validation"
)

// morocco ...
type morocco struct {
	validObj validation.ValidationI
}

// moroccoLock ...
var moroccoLock = &sync.Mutex{}

// moroccoInstance ...
var moroccoInstance *morocco

// NewMorocco ...
func NewMorocco(phone string) *morocco {
	if moroccoInstance == nil {
		moroccoLock.Lock()
		defer moroccoLock.Unlock()
		if moroccoInstance == nil {
			moroccoInstance = &morocco{}
		}
	}
	moroccoInstance.validObj = validation.New(&phone)
	return moroccoInstance
}

// Name ...
func (m *morocco) Name() string {
	return types.CountryLst[2]
}

// State ...
func (m *morocco) State() bool {
	if m.validObj != nil {
		return m.validObj.Valid(func() {
			m.validObj.Code("(212)")
			m.validObj.Optional(' ')
			m.validObj.OneOf('5', '6', '7', '8', '9')
			m.validObj.Digits(8)
		})
	}
	return false
}

// Code ...
func (m *morocco) Code() string {
	if code, ok := types.CountryCodes[types.CountryLst[2]]; ok {
		return code
	}
	return ""
}

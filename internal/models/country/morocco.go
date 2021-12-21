package country

import (
	"regexp"
	"sync"
)

// morocco ...
type morocco struct {
	phone string
}

// moroccoLock ...
var moroccoLock = &sync.Mutex{}

// moroccoInstance ...
var moroccoInstance *morocco

// newMorocco ...
func newMorocco(phone string) CountryI {
	if moroccoInstance == nil {
		moroccoLock.Lock()
		defer moroccoLock.Unlock()
		if moroccoInstance == nil {
			moroccoInstance = &morocco{}
		}
	}
	moroccoInstance.phone = phone
	return moroccoInstance
}

// Name ...
func (m *morocco) Name() string {
	return "Morocco"
}

// State ...
func (m *morocco) State() bool {
	var re = regexp.MustCompile(`\(212\)\ ?[5-9]\d{8}$`)
	return re.MatchString(m.phone)
}

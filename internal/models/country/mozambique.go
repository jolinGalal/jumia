package country

import (
	"regexp"
	"sync"
)

// mozambique ...
type mozambique struct {
	phone string
}

// mozambiqueLock ...
var mozambiqueLock = &sync.Mutex{}

// mozambiqueInstance ...
var mozambiqueInstance *mozambique

// newMozambique ...
func newMozambique(phone string) CountryI {
	if mozambiqueInstance == nil {
		mozambiqueLock.Lock()
		defer mozambiqueLock.Unlock()
		if mozambiqueInstance == nil {
			mozambiqueInstance = &mozambique{}
		}
	}

	mozambiqueInstance.phone = phone
	return mozambiqueInstance
}

// Name ...
func (m *mozambique) Name() string {
	return "Mozambique"
}

// State ...
func (m *mozambique) State() bool {
	var re = regexp.MustCompile(`\(258\)\ ?[28]\d{7,8}$`)
	return re.MatchString(m.phone)
}

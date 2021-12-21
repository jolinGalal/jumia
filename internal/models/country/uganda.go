package country

import (
	"regexp"
	"sync"
)

// uganda ...
type uganda struct {
	phone string
}

// ugandaLock ...
var ugandaLock = &sync.Mutex{}

// ugandaInstance ...
var ugandaInstance *uganda

// newUganda ...
func newUganda(phone string) CountryI {
	if ugandaInstance == nil {
		ugandaLock.Lock()
		defer ugandaLock.Unlock()
		if ugandaInstance == nil {
			ugandaInstance = &uganda{}
		}
	}
	ugandaInstance.phone = phone
	return ugandaInstance
}

//Name ...
func (u *uganda) Name() string {
	return "Uganda"
}

// State ...
func (u *uganda) State() bool {
	var re = regexp.MustCompile(`\(256\)\ ?\d{9}$`)
	return re.MatchString(u.phone)
}

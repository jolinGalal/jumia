package country

import (
	"regexp"
	"sync"
)

// ethiopia ...
type ethiopia struct {
	phone string
}

// ethiopiaLock ...
var ethiopiaLock = &sync.Mutex{}

// ethiopiaInstance ...
var ethiopiaInstance *ethiopia

// newEthiopia ...
func newEthiopia(phone string) CountryI {
	if ethiopiaInstance == nil {
		ethiopiaLock.Lock()
		defer ethiopiaLock.Unlock()
		if ethiopiaInstance == nil {
			ethiopiaInstance = &ethiopia{}
		}
	}
	ethiopiaInstance.phone = phone
	return ethiopiaInstance
}

// Name ...
func (e *ethiopia) Name() string {
	return "Ethiopia"
}

// State ...
func (e *ethiopia) State() bool {
	var re = regexp.MustCompile(`\(251\)\ ?[1-59]\d{8}$`)
	return re.MatchString(e.phone)
}

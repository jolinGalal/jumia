package country

import (
	"regexp"
	"sync"
)

// cameroon ...
type cameroon struct {
	phone string
}

// cameroonLock ...
var cameroonLock = &sync.Mutex{}

// cameroonInstance ...
var cameroonInstance *cameroon

// newCameroon ...
func newCameroon(phone string) CountryI {
	if cameroonInstance == nil {
		cameroonLock.Lock()
		defer cameroonLock.Unlock()
		if cameroonInstance == nil {
			cameroonInstance = &cameroon{}
		}
	}
	cameroonInstance.phone = phone
	return cameroonInstance
}

// Name ...
func (c *cameroon) Name() string {
	return "Cameroon"
}

// State ...
func (c *cameroon) State() bool {
	var re = regexp.MustCompile(`\(237\)\ ?[2368]\d{7,8}$`)
	return re.MatchString(c.phone)
}

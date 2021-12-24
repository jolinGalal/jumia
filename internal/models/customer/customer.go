package customer

import (
	"sync"
)

// Customer ...
type Customer struct {
	ID    int
	Name  string
	Phone string
}

// lock ...
var lock = &sync.Mutex{}

// customerInstance ...
var customerInstance *Customer

// New ...
func New() CustomerI {
	if customerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if customerInstance == nil {
			customerInstance = &Customer{}
		}
	}

	return customerInstance
}

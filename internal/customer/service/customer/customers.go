package customer

import (
	"log"

	"github.com/jinzhu/gorm"
	customers "github.com/jolinGalal/jumia/internal/customer/gen/customers"
)

// customers service example implementation.
// The example methods log the requests and return zero values.
type customerssrvc struct {
	logger *log.Logger
	db     *gorm.DB
}

// NewCustomers returns the customers service implementation.
func NewCustomers(logger *log.Logger, db *gorm.DB) customers.Service {
	return &customerssrvc{logger: logger, db: db}
}

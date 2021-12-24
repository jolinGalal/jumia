package customer

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// customerServiceMock ...
type customerServiceMock struct {
	Mock   sqlmock.Sqlmock
	GormDB *gorm.DB
	SVC    customerssrvc
}

// obj of customer service
var obj customerServiceMock

//newCustomerService ...
func newCustomerService() {
	logger := log.New(os.Stderr, "[jumia] ", log.Ltime)
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	g, err := gorm.Open("sqlite3", db)
	if err != nil {
		log.Fatalln(err)
	}

	obj.Mock = mock
	obj.GormDB = g
	obj.SVC = customerssrvc{
		db:     obj.GormDB,
		logger: logger,
	}
}

func TestMain(m *testing.M) {

	// run the test suite setup
	setup()

	// run all the unit test in here
	code := m.Run()

	// run after the test suites get executed
	teardown()
	os.Exit(code)
}

//setup run before all the tests get Executed
func setup() {

}

//teardown run after all tests get executed
func teardown() {
}

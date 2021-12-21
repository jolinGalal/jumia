package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// DatabaseI database interface
type DatabaseI interface {
	Folder(dbFolder string) DatabaseI
	GetInstance() (DatabaseI, error)
	Select(str ...string) DatabaseI
	SelectDistinct(str ...string) DatabaseI
	LeftJoin(tableName, fOperator, sOperator string) DatabaseI
	RightJoin(tableName, fOperator, sOperator string) DatabaseI
	Join(tableName, fOperator, sOperator string) DatabaseI
	Where(condition string) DatabaseI
	From(table string) DatabaseI
	Paging(pageNumber, pageSize *int) DatabaseI
	SortAsc(str ...string) DatabaseI
	SortDesc(str ...string) DatabaseI
	Count() (*int64, error)
	Find(res interface{}) error
	Save(i interface{}) error
	Delete(table, id interface{}) error
	Execute(str string) error
}

// Builder ...
type Builder struct {
	dbFolder string
	DB       *gorm.DB
}

// New ...
func New() *Builder {
	return &Builder{}
}

// Folder ...
func (b *Builder) Folder(dbFolder string) DatabaseI {
	b.dbFolder = dbFolder
	return b
}

// Get builds the database instance
func (b *Builder) GetInstance() (DatabaseI, error) {
	if len(b.dbFolder) == 0 {
		return nil, fmt.Errorf("Invalid folder")
	}
	return b.conn()
}

//conn connect to database
func (b *Builder) conn() (DatabaseI, error) {
	db, err := gorm.Open("sqlite3", b.dbFolder)
	if err != nil {
		return nil, err
	}
	b.DB = db
	return b, nil
}

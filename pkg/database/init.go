package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseI interface {
	GetInstance() (*gorm.DB, error)
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
func (b *Builder) GetInstance() (*gorm.DB, error) {
	if len(b.dbFolder) == 0 {
		return nil, fmt.Errorf("Invalid folder")
	}
	return b.conn()
}

//conn connect to database
func (b *Builder) conn() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", b.dbFolder)
	if err != nil {
		return nil, err
	}
	return db, nil
}

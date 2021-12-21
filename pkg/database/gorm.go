package database

import (
	"strings"
)

// Select ...
func (b *Builder) Select(str ...string) DatabaseI {
	b.DB = b.DB.Select(strings.Join(str, ","))
	return b
}

// SelectDistinct ...
func (b *Builder) SelectDistinct(str ...string) DatabaseI {
	b.DB = b.DB.Select("DISTINCT " + strings.Join(str, ","))
	return b
}

// LeftJoin ...
func (b *Builder) LeftJoin(tableName, fOperator, sOperator string) DatabaseI {
	b.DB = b.DB.Joins("left join " + tableName + " on " + fOperator + " = " + sOperator)
	return b
}

// RightJoin ...
func (b *Builder) RightJoin(tableName, fOperator, sOperator string) DatabaseI {
	b.DB = b.DB.Joins("right join " + tableName + " on " + fOperator + " = " + sOperator)
	return b
}

// Join ...
func (b *Builder) Join(tableName, fOperator, sOperator string) DatabaseI {
	b.DB = b.DB.Joins("join " + tableName + " on " + fOperator + " = " + sOperator)
	return b
}

// Where ...
func (b *Builder) Where(condition string) DatabaseI {
	b.DB = b.DB.Where(condition)
	return b
}

// Model ...
func (b *Builder) From(table string) DatabaseI {
	b.DB = b.DB.Table(table)
	return b
}

// Paging ...
func (b *Builder) Paging(pageNumber, pageSize *int) DatabaseI {
	if pageSize != nil && pageNumber != nil {
		if *pageSize != -1 && *pageNumber != -1 {
			offset := (*pageNumber - 1) * *pageSize
			b.DB = b.DB.Limit(*pageSize).Offset(offset)
		}
	}
	return b
}

// SortAsc ...
func (b *Builder) SortAsc(str ...string) DatabaseI {
	b.DB = b.DB.Order(strings.Join(str, "asc,"))
	return b
}

// SortDesc ...
func (b *Builder) SortDesc(str ...string) DatabaseI {
	b.DB = b.DB.Order(strings.Join(str, "desc,"))
	return b
}

// Count ...
func (b *Builder) Count() (*int64, error) {
	totalCount := new(int64)
	err := b.DB.Count(totalCount).Error
	return totalCount, err
}

// Find ...
func (b *Builder) Find(res interface{}) error {
	return b.DB.Find(res).Error
}

// Save ...
func (b *Builder) Save(i interface{}) error {
	return b.DB.Save(i).Error
}

// Detele ...
func (b *Builder) Delete(table, id interface{}) error {
	return b.DB.Delete(table, id).Error
}

// Execute ...
func (b *Builder) Execute(str string) error {
	return b.DB.Exec(str).Error
}

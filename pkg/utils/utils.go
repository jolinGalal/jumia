package utils

import "github.com/jinzhu/gorm"

// Paging ...
func Paging(db *gorm.DB, pageNumber, pageSize *int) *gorm.DB {
	if pageSize != nil && pageNumber != nil {
		if *pageSize != -1 && *pageNumber != -1 {
			offset := (*pageNumber - 1) * *pageSize
			db = db.Limit(*pageSize).Offset(offset)
		}
	}
	return db
}

// Intervals function to return law and hight slice index
func Intervals(page_number, page_size *int, length int) (int, int) {
	law := (*page_number - 1) * *page_size
	if law > length {
		return -1, -1
	}
	high := *page_number * *page_size
	if high > length {
		high = length
	}
	return law, high

}

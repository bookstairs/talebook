package model

import (
	"time"
)

const TableNameBook = "books"

// Book mapped from table <books>
type Book struct {
	ID           int32     `gorm:"column:id;type:INTEGER" json:"id"`
	Title        string    `gorm:"column:title;type:TEXT" json:"title"`
	Sort         string    `gorm:"column:sort;type:TEXT" json:"sort"`
	Timestamp    time.Time `gorm:"column:timestamp;type:TIMESTAMP" json:"timestamp"`
	Pubdate      time.Time `gorm:"column:pubdate;type:TIMESTAMP" json:"pubdate"`
	SeriesIndex  float64   `gorm:"column:series_index;type:REAL" json:"series_index"`
	AuthorSort   string    `gorm:"column:author_sort;type:TEXT" json:"author_sort"`
	Isbn         string    `gorm:"column:isbn;type:TEXT" json:"isbn"`
	Lccn         string    `gorm:"column:lccn;type:TEXT" json:"lccn"`
	Path         string    `gorm:"column:path;type:TEXT" json:"path"`
	Flags        int32     `gorm:"column:flags;type:INTEGER" json:"flags"`
	UUID         string    `gorm:"column:uuid;type:TEXT" json:"uuid"`
	HasCover     string    `gorm:"column:has_cover;type:BOOL" json:"has_cover"`
	LastModified time.Time `gorm:"column:last_modified;type:TIMESTAMP" json:"last_modified"`
}

// TableName Book's table name
func (*Book) TableName() string {
	return TableNameBook
}

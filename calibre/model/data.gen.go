package model

const TableNameDatum = "data"

// Datum mapped from table <data>
type Datum struct {
	ID               int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book             int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Format           string `gorm:"column:format;type:TEXT" json:"format"`
	UncompressedSize int32  `gorm:"column:uncompressed_size;type:INTEGER" json:"uncompressed_size"`
	Name             string `gorm:"column:name;type:TEXT" json:"name"`
}

// TableName Datum's table name
func (*Datum) TableName() string {
	return TableNameDatum
}

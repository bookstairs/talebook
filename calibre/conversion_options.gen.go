package calibre

const TableNameConversionOption = "conversion_options"

// ConversionOption mapped from table <conversion_options>
type ConversionOption struct {
	ID     int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Format string `gorm:"column:format;type:TEXT" json:"format"`
	Book   int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Data   []byte `gorm:"column:data;type:BLOB" json:"data"`
}

// TableName ConversionOption's table name
func (*ConversionOption) TableName() string {
	return TableNameConversionOption
}

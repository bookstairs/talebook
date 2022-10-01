package calibre

const TableNameLastReadPosition = "last_read_positions"

// LastReadPosition mapped from table <last_read_positions>
type LastReadPosition struct {
	ID      int32   `gorm:"column:id;type:INTEGER" json:"id"`
	Book    int32   `gorm:"column:book;type:INTEGER" json:"book"`
	Format  string  `gorm:"column:format;type:TEXT" json:"format"`
	User    string  `gorm:"column:user;type:TEXT" json:"user"`
	Device  string  `gorm:"column:device;type:TEXT" json:"device"`
	Cfi     string  `gorm:"column:cfi;type:TEXT" json:"cfi"`
	Epoch   float64 `gorm:"column:epoch;type:REAL" json:"epoch"`
	PosFrac float64 `gorm:"column:pos_frac;type:REAL" json:"pos_frac"`
}

// TableName LastReadPosition's table name
func (*LastReadPosition) TableName() string {
	return TableNameLastReadPosition
}

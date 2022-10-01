package model

const TableNameRating = "ratings"

// Rating mapped from table <ratings>
type Rating struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Rating int32 `gorm:"column:rating;type:INTEGER" json:"rating"`
}

// TableName Rating's table name
func (*Rating) TableName() string {
	return TableNameRating
}

package models

type Coursestudent struct {
	StudentID int64 `gorm:"column:StudentID"`
	CourseID  int64 `gorm:"column:CourseID"`
}

func (Coursestudent) TableName() string {
	return "courseStudent"
}

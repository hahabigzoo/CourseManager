package models

type Course struct {
	CourseID   int64  `gorm:"column:CourseID"`
	CourseName string `gorm:"column:CourseName"`
	Capacity   int64  `gorm:"column:Cap"`
	TeacherID  int64  `gorm:"column:TeacherID"`
	RestCap    int64  `gorm:"column:RestCap"`
}

func (Course) TableName() string {
	return "course"
}

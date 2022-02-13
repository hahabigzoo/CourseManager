package models

import (
	"Course/configs"
	"fmt"
	"gorm.io/gorm"
)

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

/**
课程的增删改查，针对数据库的操作，与业务逻辑里的增删改查不一致
*/
func (course *Course) saveCourse() *gorm.DB {
	db := configs.DB
	err := db.Create(course)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

func (course *Course) delCourse() *gorm.DB {
	db := configs.DB
	err := db.Delete(course)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

func (course *Course) updateCourse(col string) *gorm.DB {
	db := configs.DB
	err := db.Model(&course).Select(col).Updates(course)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

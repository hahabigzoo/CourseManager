package models

import (
	"Course/configs"
	"Course/entity"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	UserID    int64           `gorm:"primaryKey;column:UserID"`
	UserName  string          `gorm:"column:UserName"`
	Password  string          `gorm:"column:Password"`
	UserState bool            `gorm:"column:UserState"`
	Nickname  string          `gorm:"column:Nickname"`
	UserType  entity.UserType `gorm:"column:UserType"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) saveUser() *gorm.DB {
	db := configs.DB
	err := db.Create(user)
	fmt.Println(user.UserID)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

func (user *User) delUser() *gorm.DB {
	db := configs.DB
	user.UserState = true
	err := db.Delete(user)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

func (user *User) updateUser(col string) *gorm.DB {
	db := configs.DB
	err := db.Model(&user).Select(col).Updates(user)
	if err != nil {
		fmt.Println(err.Error)
	}
	return err
}

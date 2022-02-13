package models

import (
	"Course/configs"
	"Course/entity"
	"gorm.io/gorm/utils/tests"
	"testing"
)

//用户的curd测试

func TestUser_TableName(t *testing.T) {
	configs.InitDB()
	user1 := &User{
		UserName:  "TestUser0",
		Password:  "123456",
		UserState: false,
		Nickname:  "TestUser0",
		UserType:  entity.Admin,
	}
	/**
	user2 := User{
		UserName:  "TestUser1",
		Password:  "123456",
		UserState: false,
		Nickname:  "TestUser1",
		UserType:  entity.Teacher,
	}
	*/
	//user2.saveUser()
	if err := user1.saveUser().Error; err != nil {
		t.Fatalf("failed to create user, got error %v", err)
	}

	var result User
	if err := configs.DB.Find(&result, user1.UserID).Error; err != nil {
		t.Fatalf("failed to query user, got error %v", err)
	}

	tests.AssertEqual(t, result, user1)
	user1.delUser()
	//user2.delUser()

}

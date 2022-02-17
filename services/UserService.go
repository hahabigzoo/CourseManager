package services

import (
	"Course/configs"
	"Course/entity"
	"Course/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

/*
	User这个Model的增删改查操作都放在这里
*/

func CreateUser(user models.User) (err error, userInter models.User) {
	err = configs.DB.Create(&user).Error
	return err, user
}

func GetUserById(id int64) (user models.User, err error) {
	err = configs.DB.First(&user, id).Error
	return
}

func GetUserByUserName(username string) (user models.User, err error) {
	err = configs.DB.Where("UserName = ?", username).First(&user).Error
	return
}

func GetUser(request entity.GetMemberRequest) (user models.User, err error) {
	// 因为是软删除，还未过滤软删除的记录
	// 直接使用数据库来筛选软删除

	id_str := request.UserID
	id, _ := strconv.ParseInt(id_str, 10, 64)
	user, err = GetUserById(id)
	return
}

func GetUserList(request entity.GetMemberListRequest) (userList []models.User, err error) {
	// 因为是软删除，还未过滤软删除的记录
	// 直接使用数据库来筛选软删除
	limit := request.Limit
	offset := request.Offset * limit
	err = configs.DB.Where(map[string]interface{}{"UserState": 0}).Offset(offset).Limit(limit).Find(&userList).Error
	return
}

func UpdateUserNicknameById(id int64, Nickname string) (err error) {
	var user models.User
	user.UserID = id
	err = configs.DB.Model(&user).Update("Nickname", Nickname).Error
	return
}

func UpdateUserStateById(id int64) (err error) {
	var user models.User
	user.UserID = id
	err = configs.DB.Model(&user).Update("UserState", 1).Error
	return
}

func DeleteUserById(id int64) (err error) {
	return UpdateUserStateById(id)
}

func DeleteUser(request entity.DeleteMemberRequest) (response entity.DeleteMemberResponse) {
	id_str := request.UserID
	id, err1 := strconv.ParseInt(id_str, 10, 64)
	if err1 != nil {
		response.Code = entity.ParamInvalid
	}
	user, err := GetUserById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Code = entity.UserNotExisted
		return
	}
	if err != nil {
		response.Code = entity.UnknownError
		fmt.Println(err)
		return
	}
	if user.UserState == true {
		response.Code = entity.UserHasDeleted
		return
	}
	if err2 := DeleteUserById(id); err2 != nil {
		response.Code = entity.UnknownError
		fmt.Println(err)
		return
	}
	response.Code = entity.OK
	return
}

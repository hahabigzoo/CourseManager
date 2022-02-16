package services

import (
	"Course/configs"
	"Course/entity"
	"Course/models"
)

/*
	User这个Model的增删改查操作都放在这里
*/

func CreateUser(user models.User) (err error, userInter models.User){
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

func GetUserList(request entity.GetMemberListRequest) (userList []models.User, err error) {
	// 因为是软删除，还未过滤软删除的记录
	// todo
	limit := request.Limit
	offset := request.Offset
	err = configs.DB.Limit(limit).Offset(offset).Find(&userList).Error
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

func DeleteUserById(id int64) (err error){
	return UpdateUserStateById(id)
}
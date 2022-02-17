package utils

import (
	"Course/entity"
	"strconv"
)

func CreateUserVerify(request entity.CreateMemberRequest) bool {
	// 用户昵称
	if len(request.Nickname) < 4 || len(request.Nickname) > 20 {
		return false
	}
	//for _, value := range request.Nickname {
	//	if !(('a' <= value && value <= 'z') || ('A' <= value && value <= 'Z')) {
	//		return false
	//	}
	//}
	// 用户名
	if len(request.Username) < 8 || len(request.Username) > 20 {
		return false
	}
	for _, value := range request.Username {
		if !(('a' <= value && value <= 'z') || ('A' <= value && value <= 'Z')) {
			return false
		}
	}
	// 密码
	if len(request.Password) < 8 || len(request.Password) > 20 {
		return false
	}
	num := false
	lowLetter := false
	upLetter := false
	for _, value := range request.Password {
		if 'a' <= value && value <= 'z' {
			lowLetter = true
		} else if 'A' <= value && value <= 'Z' {
			upLetter = true
		} else if '0' <= value && value <= '9' {
			num = true
		} else {
			return false
		}
	}
	if !(num && lowLetter && upLetter) {
		return false
	}
	// 用户类型
	if request.UserType != entity.Admin && request.UserType != entity.Student && request.UserType != entity.Teacher {
		return false
	}
	return true
}

func UpdateUserVerify(request entity.UpdateMemberRequest) bool {
	if _, err := strconv.ParseInt(request.UserID, 10, 64); err != nil {
		return false
	}
	// 用户昵称
	if len(request.Nickname) < 4 || len(request.Nickname) > 20 {
		return false
	}
	for _, value := range request.Nickname {
		if !(('a' <= value && value <= 'z') || ('A' <= value && value <= 'Z')) {
			return false
		}
	}
	return true
}

package controllers

import (
	"Course/entity"
	"Course/models"
	"Course/services"
	"Course/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var request entity.CreateMemberRequest
	var response entity.CreateMemberResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)

	// 校验参数是否合法
	if err := utils.CreateUserVerify(request); !err {
		response.Code = entity.ParamInvalid
		c.JSON(http.StatusOK, response)
		return
	}

	// 权限判断
	// 判断当前用户是不是管理员，只有管理员才有创建用户的权限
	// 通过登录模块实现的接口，g.GET("/auth/whoami")， 判断当前用户是不是管理员
	// todo

	// 检测用户名是否已存在
	if _, err := services.GetUserByUserName(request.Username); err == nil {
		//用户名已存在
		response.Code = entity.UserHasExisted
		c.JSON(http.StatusOK, response)
		return
	}

	// 构造需新建的成员对象
	member := &models.User{UserName: request.Username, Password: request.Password, Nickname: request.Nickname, UserType: request.UserType}
	//err, memberReturn := service.MemberService.CreateMember(*member)
	err, memberReturn := services.CreateUser(*member)
	if err != nil {
		response.Code = entity.UnknownError
		c.JSON(http.StatusOK, response)
	} else {
		response.Code = entity.OK
		UserIDStr := strconv.FormatInt(memberReturn.UserID,10)
		response.Data = struct{ UserID string }{UserID: UserIDStr}
		c.JSON(http.StatusOK, response)
	}
}

// 获取单个成员
func GetUser(c *gin.Context)  {

}

// 批量获取成员
func GetUserList(c *gin.Context) {

}

// 更新成员, 只允许更新昵称
func UpdateUser(c *gin.Context) {
	var request entity.UpdateMemberRequest
	var response entity.UpdateMemberResponse
	// 获取参数，UserId
	_ = c.ShouldBind(&request)
	//参数校验
	if err := utils.UpdateUserVerify(request); !err {
		response.Code = entity.ParamInvalid
		c.JSON(http.StatusOK, response)
		return
	}
	//string转int64
	UserID, err := strconv.ParseInt(request.UserID, 10, 64)
	if err != nil {
		// 非法UserID，不是int64数字字符串
		response.Code = entity.ParamInvalid
		c.JSON(http.StatusOK, response)
		return
	}

	member, err := services.GetUserById(UserID)
	if err != nil {
		response.Code = entity.UserNotExisted
		c.JSON(http.StatusOK, response)
		return
	}
	// 需更新的用户是否已删除
	if member.UserState == true {
		response.Code = entity.UserHasDeleted
		c.JSON(http.StatusOK, response)
		return
	}

	if err := services.UpdateUserNicknameById(UserID, request.Nickname);err != nil{
		response.Code = entity.UnknownError
		c.JSON(http.StatusOK, response)
	}else{
		response.Code = entity.OK
		c.JSON(http.StatusOK, response)
	}
}

// 删除成员, 成员删除后，该成员不能够被登录且不应该不可见
func DeleteUser(c *gin.Context) {

}

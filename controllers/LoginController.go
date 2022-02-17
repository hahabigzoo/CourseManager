package controllers

import (
	"Course/configs"
	"Course/entity"
	"Course/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	var request entity.LoginRequest
	var response entity.LoginResponse

	cookie, err := c.Cookie("camp-session")

	session := sessions.Default(c)

	//不带cookie同时也没有session
	if err == nil {
		if err2 := c.ShouldBind(&request); err2 != nil {
			response.Code = entity.ParamInvalid
		} else {
			JudgeUP(c, request, &response)
		}
	} else { //有cookie
		name := session.Get(cookie)
		if name == nil {
			if err2 := c.ShouldBind(&request); err2 != nil {
				response.Code = entity.ParamInvalid
			} else {
				JudgeUP(c, request, &response)
			}
		} else {
			response.Code = entity.OK
			user, _ := GetUserByUserName(cookie)
			response.Data.UserID = strconv.FormatInt(user.UserID, 10)
		}
	}
	c.JSON(http.StatusOK, response)
	return
}

func Logout(c *gin.Context) {
	var request entity.LogoutRequest
	var response entity.LogoutResponse
	//删除session
	_ = c.ShouldBind(&request)
	session := sessions.Default(c)
	session.Delete("camp-session")
	session.Save()
	response.Code = entity.OK
	c.JSON(http.StatusOK, response)
	return

}

func Whoami(c *gin.Context) {
	var request entity.WhoAmIRequest
	var response entity.WhoAmIResponse

	_ = c.ShouldBind(&request)
	session := sessions.Default(c)

	cookie, err := c.Cookie("camp-session")
	if err == nil { //判断有无cookie
		response.Code = entity.LoginRequired
		c.JSON(http.StatusOK, response)
		return
	} else { //判断是否登录，查找session是否存在
		name := session.Get(cookie)
		if name == nil {
			response.Code = entity.LoginRequired
			c.JSON(http.StatusOK, response)
			return
		} else { //存在session，已经登录，寻找用户，返回用户
			response.Code = entity.OK
			response.Data, _ = GetTMemberByUserName(cookie)
			c.JSON(http.StatusOK, response)
			return
		}
	}

}

//设置session值
func setsession(ctx *gin.Context, username string) {
	//初始化session对象
	session := sessions.Default(ctx)

	//设置session
	session.Set(username, username)
	session.Save()

	ctx.SetCookie("camp-session", username, 3600, "/", "localhost", false, true)
}

//用户名获取user
func GetUserByUserName(username string) (user models.User, err error) {
	err = configs.DB.Where("UserName = ?", username).First(&user).Error
	fmt.Println(user)
	return
}

//用户名获取TMember
func GetTMemberByUserName(username string) (Tmen entity.TMember, err error) {
	var user models.User
	err = configs.DB.Where("UserName = ?", username).First(&user).Error
	Tmen.UserType = user.UserType
	Tmen.Username = user.UserName
	Tmen.Nickname = user.Nickname
	Tmen.UserID = strconv.FormatInt(user.UserID, 10)
	return
}

//判断用户密码是否正确
func JudgeUP(c *gin.Context, request entity.LoginRequest, response *entity.LoginResponse) {
	user, uerr := GetUserByUserName(request.Username)
	if uerr != nil {
		response.Code = entity.UserNotExisted

	} else if request.Password != user.Password {
		response.Code = entity.WrongPassword

	} else if request.Password == user.Password {
		setsession(c, request.Username)
		response.Code = entity.OK
		response.Data.UserID = strconv.FormatInt(user.UserID, 10)
	}
}

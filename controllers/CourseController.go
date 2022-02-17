package controllers

import (
	"Course/entity"
	"Course/models"
	"Course/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCourse(c *gin.Context) {
	var request entity.CreateCourseRequest
	var response entity.CreateCourseResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)

	// 构造需新建的成员对象
	course := &models.Course{CourseName: request.Name, Capacity: request.Cap, RestCap: request.Cap}
	err, _ := services.CreateCourse(course)
	//err, memberReturn := service.MemberService.CreateMember(*member)
	if err != nil {
		response.Code = entity.UnknownError
		c.JSON(http.StatusOK, response)
	} else {
		response.Code = entity.OK
		response.Data = struct {
			CourseID string
		}{strconv.FormatInt(course.CourseID, 10)}
		c.JSON(http.StatusOK, response)
	}
}

func GetCourse(c *gin.Context) {
	var request entity.GetCourseRequest
	var response entity.GetCourseResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)

	// 查询课程信息对象
	response = services.GetCourse(request)
	c.JSON(http.StatusOK, response)
}

func BindCourse(c *gin.Context) {
	var request entity.BindCourseRequest
	var response entity.BindCourseResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)

	// 绑定课程信息对象
	response = services.BindCourseTeacher(request)
	c.JSON(http.StatusOK, response)
}

func UnbindCourse(c *gin.Context) {
	var request entity.UnbindCourseRequest
	var response entity.UnbindCourseResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)

	// 绑定课程信息对象
	response = services.UnbindCourseTeacher(request)
	c.JSON(http.StatusOK, response)
}

func GetTeacherCourse(c *gin.Context) {
	var request entity.GetTeacherCourseRequest
	var response entity.GetTeacherCourseResponse
	// 获取请求参数
	_ = c.ShouldBind(&request)
	// 绑定课程信息对象
	response = services.GetTeacherCourse(request)
	c.JSON(http.StatusOK, response)
}

package controllers

import (
	"Course/entity"
	"Course/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func ScheuleCoursecontroller(c *gin.Context) {
	//解析请求
	req := &entity.ScheduleCourseRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, entity.ScheduleCourseResponse{Code: entity.OK})
		return
	}
	/**
	测试代码
	var teacherCourseRelationShip map[string][]string
	teacherCourseRelationShip = req.TeacherCourseRelationShip
	fmt.Println("ScheuleCoursecontroller")
	fmt.Println(teacherCourseRelationShip)
	*/
	//调用服务并返回
	resp := services.ScheduleCourseService(req)
	c.JSON(200, resp)
}

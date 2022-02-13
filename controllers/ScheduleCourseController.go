package controllers

import (
	"Course/entity"
	"Course/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ScheuleCoursecontroller(c *gin.Context) {
	var teacherCourseRelationShip map[string][]string

	err := json.Unmarshal([]byte(c.Query("TeacherCourseRelationShip")), &teacherCourseRelationShip)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
	}
	req := entity.ScheduleCourseRequest{TeacherCourseRelationShip: teacherCourseRelationShip}
	resp := services.ScheduleCourseService(req)
	c.JSON(200, resp)
}

package controllers

import (
	"Course/serializer"
	"Course/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func ScheuleCoursecontroller(c *gin.Context) {
	var teacherCourseRelationShip map[string][]string

	err := json.Unmarshal([]byte(c.Query("TeacherCourseRelationShip")), &teacherCourseRelationShip)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
	}
	req := serializer.ScheduleCourseRequest{TeacherCourseRelationShip: teacherCourseRelationShip}
	resp := services.ScheduleCourseService(req)
	c.JSON(200, resp)
}

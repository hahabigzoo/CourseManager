package controllers

import (
	"Course/configs"
	"Course/entity"
	"Course/models"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var rw sync.RWMutex

func BookCourse(c *gin.Context) {
	var json entity.BookCourseRequest
	if err := c.ShouldBind(&json); err != nil {
		fmt.Println(json)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//第1步：查询学生ID是否存在
	studentID, _ := strconv.ParseInt(json.StudentID, 10, 64)
	if exist, _ := configs.Rdb.Exists(context.Background(), "courseNumber:"+json.StudentID).Result(); exist != 1 {
		var user models.User
		rw.RLock()
		if err := configs.DB.Where("UserID = ? AND UserType =?", studentID, entity.Student).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) || user.UserState {
			rw.RUnlock()
			c.JSON(http.StatusOK, gin.H{"Code": entity.StudentNotExisted})
			return
		}
		rw.RUnlock()
	}

	//第2步：查询课程ID是否存在
	courseID, _ := strconv.ParseInt(json.CourseID, 10, 64)
	if exist, _ := configs.Rdb.Exists(context.Background(), "course:"+json.StudentID).Result(); exist != 1 {
		var course models.Course
		rw.RLock()
		if configs.DB.First(&course, courseID); course.CourseID == 0 {
			rw.RUnlock()
			c.JSON(http.StatusOK, gin.H{"Code": entity.CourseNotExisted})
			return
		}
		rw.RUnlock()
		cour := make(map[string]interface{})
		cour["CourseName"] = course.CourseName
		cour["TeacherID"] = course.TeacherID
		cour["RestCap"] = course.RestCap
		configs.Rdb.HMSet(context.Background(), "course:"+json.CourseID, cour)
	}
	//查询重复选课
	if repeat, _ := configs.Rdb.SIsMember(context.Background(), "courseDetail:"+json.StudentID, json.CourseID).Result(); repeat {
		c.JSON(http.StatusOK, gin.H{"Code": entity.StudentHasCourse})
		return
	}

	//第4步：查询是否有课余量
	rest, _ := configs.Rdb.HGet(context.Background(), "course:"+json.CourseID, "RestCap").Result()
	t := fmt.Sprintf("%v", rest)
	restCap, _ := strconv.ParseInt(t, 10, 64)
	fmt.Println(restCap)
	if restCap == 0 {
		c.JSON(http.StatusOK, gin.H{"Code": entity.CourseNotAvailable})
		return
	}

	//第5步：抢课成功，更新redis和数据库
	configs.Rdb.HIncrBy(context.Background(), "course:"+json.CourseID, "restCap", -1)
	configs.Rdb.Incr(context.Background(), "courseNumber:"+json.StudentID)
	configs.Rdb.SAdd(context.Background(), "courseDetail:"+json.StudentID, json.CourseID)
	rw.Lock()
	configs.DB.Table("course").Where("CourseID = ?", courseID).UpdateColumn("RestCap", restCap-1)
	configs.DB.Table("courseStudent").Create(models.Coursestudent{studentID, courseID})
	rw.Unlock()
	c.JSON(http.StatusOK, gin.H{"Code": entity.OK})
}

func Course(c *gin.Context) {
	var json entity.GetStudentCourseRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//第1步：查询学生ID是否存在
	studentID, _ := strconv.ParseInt(json.StudentID, 10, 64)
	if exist, _ := configs.Rdb.Exists(context.Background(), "courseNumber:"+json.StudentID).Result(); exist != 1 {
		var user models.User
		if configs.DB.First(&user, json.StudentID); !user.UserState {
			c.JSON(http.StatusOK, gin.H{"Code": entity.StudentNotExisted})
			return
		}
	}

	flag := true //true表示进Redis查，false表示进mysql查
	if exist, _ := configs.Rdb.Exists(context.Background(), "courseDetail:"+json.StudentID).Result(); exist != 1 {
		flag = false
	} else {
		num1 := configs.Rdb.Get(context.Background(), "courseNumber:"+json.StudentID).Val()
		num2 := configs.Rdb.SCard(context.Background(), "courseDetail:"+json.StudentID).Val()
		if num3, _ := strconv.ParseInt(num1, 10, 64); num3 != num2 {
			flag = false
		}
	}

	courseListInfo := make([]entity.TCourse, 0)
	if !flag {
		configs.Rdb.Del(context.Background(), "courseDetail:"+json.StudentID)
		//第2步：根据学生ID查询他的课程
		var courseList []models.Coursestudent
		configs.DB.Table("courseStudent").Find(&courseList, "StudentID = ?", studentID)

		//第3步：把查询到的课程装配成courseListInfo
		for _, v := range courseList {
			course, _ := configs.Rdb.HGetAll(context.Background(), "course:"+strconv.FormatInt(v.CourseID, 10)).Result()
			courseListInfo = append(courseListInfo, entity.TCourse{strconv.FormatInt(v.CourseID, 10), course["CourseName"], course["TeacherID"]})
			configs.Rdb.SAdd(context.Background(), "courseDetail:"+json.StudentID, v.CourseID)
		}

		//第4步：存入redis的courseNumber:{StudentID}和courseDetail:{StudentID}，处理幂等问题
		configs.Rdb.Set(context.Background(), "courseNumber:"+json.StudentID, len(courseList), 0)

	} else {
		courseList, _ := configs.Rdb.SMembers(context.Background(), "courseDetail:"+json.StudentID).Result()
		for _, courseID := range courseList {
			course, _ := configs.Rdb.HGetAll(context.Background(), "course:"+courseID).Result()
			courseListInfo = append(courseListInfo, entity.TCourse{courseID, course["CourseName"], course["TeacherID"]})
		}
		fmt.Println(courseList)
	}

	//第5步：返回
	code := entity.StudentHasCourse
	if len(courseListInfo) == 0 {
		code = entity.StudentHasNoCourse
	}
	c.JSON(http.StatusOK, gin.H{"Code": code, "Data": courseListInfo})
}

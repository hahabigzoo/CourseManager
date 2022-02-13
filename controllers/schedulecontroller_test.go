package controllers

import (
	"Course/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUserController_QueryUserByName(t *testing.T) {
	// 初始化路由
	router := gin.Default()
	g := router.Group("/api/v1")
	g.POST("/course/schedule", ScheuleCoursecontroller)
	//mock data
	schedulecourse := entity.ScheduleCourseRequest{TeacherCourseRelationShip: map[string][]string{"sun": {
		"asd",
		"sdf",
		"dewewfg",
	},
		"sun2": {
			"asdqw",
			"sdewewqf",
			"dfweg",
		},
	}}

	jsonByte, _ := json.Marshal(schedulecourse)

	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", "http://127.0.0.1/api/v1/course/schedule", bytes.NewReader(jsonByte))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}

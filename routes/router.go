package routes

import (
	"Course/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create", controllers.CreateUser)
	g.GET("/member", controllers.GetUser)
	g.GET("/member/list", controllers.GetUserList)
	g.POST("/member/update", controllers.UpdateUser)
	g.POST("/member/delete", controllers.DeleteUser)

	// 登录

	g.POST("/auth/login", controllers.Login)
	g.POST("/auth/logout", controllers.Logout)
	g.GET("/auth/whoami", controllers.Whoami)

	// 排课
	g.POST("/course/create", controllers.CreateCourse)
	g.GET("/course/get", controllers.GetCourse)

	g.POST("/teacher/bind_course", controllers.BindCourse)
	g.POST("/teacher/unbind_course", controllers.UnbindCourse)
	g.GET("/teacher/get_course", controllers.GetTeacherCourse)
	g.POST("/course/schedule", controllers.ScheuleCoursecontroller)

	// 抢课
	g.POST("/student/book_course", controllers.BookCourse)
	g.GET("/student/course", controllers.Course)
	g.GET("/ping", controllers.Pong)

}

package main

import (
	"Course/configs"
	"Course/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()
	configs.InitClient()
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("camp-session", store))
	routes.RegisterRouter(router)
	router.Run(":80")
}

//进程号23838

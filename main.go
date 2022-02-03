package main

import (
	"Course/configs"
	"Course/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()
	MysqlDb := configs.GetDB()
	defer MysqlDb.Close()
	router := gin.Default()
	routes.RegisterRouter(router)
	router.Run(":80")
}

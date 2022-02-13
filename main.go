package main

import (
	"Course/configs"
	"Course/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()
	configs.InitClient()
	router := gin.Default()
	routes.RegisterRouter(router)
	router.Run(":80")
}

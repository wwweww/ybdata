package main

import (
	"YBData/controller"
	"YBData/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.POST("/search/class", controller.SearchClass)
	r.POST("/search/student", controller.SearchStudent)
	return r
}

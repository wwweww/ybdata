package main

import (
	"YBData/service"
	"github.com/gin-gonic/gin"
)

func main() {
	service.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":9578"))
}

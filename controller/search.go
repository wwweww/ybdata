package controller

import (
	"YBData/module"
	"YBData/service"
	"github.com/gin-gonic/gin"
)

type SCFromData struct {
	College string
}

type SSFromData struct {
	Class string
}

func SearchClass(c *gin.Context) {
	var majors []string
	var fromData SCFromData
	err := c.BindJSON(&fromData)

	if err != nil {
		c.JSON(403,
			gin.H{
				"code": 403,
				"msg":  "参数不正确",
			})
		c.Abort()
		return
	}

	db := service.GetDB()
	db.Model(&module.StData{}).
		Distinct().
		Where("college = ?", fromData.College).
		Pluck("classid", &majors)
	c.JSON(200,
		gin.H{
			"code": 200,
			"data": majors,
			"msg":  "查询成功",
		})
}

func SearchStudent(c *gin.Context) {
	var fromData SSFromData
	err := c.BindJSON(&fromData)
	if err != nil {
		c.JSON(403,
			gin.H{
				"code": 403,
				"msg":  "参数不正确",
			})
		c.Abort()
		return
	}

	db := service.GetDB()
	var students []module.StData
	db.Model(&module.StData{}).
		Where("classid = ?", fromData.Class).
		Find(&students)

	c.JSON(200,
		gin.H{
			"code": 200,
			"data": students,
			"msg":  "查询成功",
		})
}

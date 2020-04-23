package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main(){
	router:=gin.Default()
	router.Use(cors.Default())

	router.POST("/KeBiao/xuanke",func(c *gin.Context){//选课
		stunum:=c.PostForm("stunum")
		username:=c.PostForm("username")
		dacourse_num:=c.PostForm("dacourse_num")
		stunumm,_ :=strconv.Atoi(stunum)

		if err := CreateAnimals(stunumm,username,dacourse_num);err != nil{
			c.JSON(200,gin.H{"state":10001,"message":"选课失败","err":err.Error()})
		}

		c.JSON(200,gin.H{"state":10000,"message":"选课成功"})
	})

	_ = router.Run(":8080")
}
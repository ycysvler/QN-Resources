package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/QN-Resources/server/entity"
	"github.com/QN-Resources/server/sections"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)


func main(){

	fmt.Printf("qiniu:%v\n", sections.Qiniu)

	mac := qbox.NewMac(
		"z2uEqEqNRC1APAm4j78V5JMaPLCYeHdi6JBdPSMt",
		"mchL2Tk6jJt0Kp5Dz-K374MgrX1UtNyuBt7n9V5-")

	putPolicy := storage.PutPolicy{
		Scope: "test",
	}
	upToken := putPolicy.UploadToken(mac)

	fmt.Printf("upToken:%v\n", upToken)

	root := entity.Instance()

	root.AddTreeNode("name", "path", 0)

	fmt.Printf("root:%v\n",root)

	// 设置 gin 的模式（调试模式：DebugMode, 发行模式：ReleaseMode）
	gin.SetMode(gin.DebugMode)
	// 创建一个不包含中间件的路由器
	r := gin.Default()

	r.Static("/static", "./static")
	// static icon
	r.StaticFile("/favicon.ico", "./static/images/favicon.ico")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/nodes", func(c *gin.Context){
		c.JSON(200, root)
	})


	// 在8080 端口，启动http服务
	r.Run(":8090")
}
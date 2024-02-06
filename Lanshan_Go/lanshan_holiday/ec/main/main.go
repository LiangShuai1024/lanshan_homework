package main

import (
	"demo/ec/bootstrap"
	"demo/ec/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	//初始化日志
	global.Ec.Log = bootstrap.InitializeLog()
	global.Ec.Log.Info("log init success!")
	// 初始化数据库
	global.Ec.DB = bootstrap.InitializeDB()
	// 初始化验证器
	bootstrap.InitializeValidator() //bootstrap/validator

	r := gin.Default()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.Ec.DB != nil {
			db, _ := global.Ec.DB.DB()
			db.Close()
		}
	}()

	// 测试路由,成功返回pong
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.Ec.Redis = bootstrap.InitializeRedis()

	// 启动服务器
	//r.Run(":" + global.Ec.Config.Ec.Port)
	bootstrap.RunServer() //bootstrap/router
}

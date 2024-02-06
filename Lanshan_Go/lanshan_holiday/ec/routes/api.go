package routes

import (
	"demo/ec/controllers/app"
	"demo/ec/middleware"
	"demo/ec/services"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.POST("/auth/register", app.Register)

	router.POST("/auth/login", app.Login)

	//router.POST("/auth/update", app.Update)

	router.POST("/product/create", app.Create)


	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info) //controllers/app/auth
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/auth/update", app.Update)
	}
}

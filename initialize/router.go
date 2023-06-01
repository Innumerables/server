package initialize

import (
	"net/http"
	"server/global"
	"server/router"

	"server/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	// Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)    //注册功能api路由
		systemRouter.InitCasbinRouter(PrivateGroup) //权限相关的路由
		systemRouter.InitUserRouter(PrivateGroup)   //用户相关的路由

	}
	return Router
}

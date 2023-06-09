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
	exampleRouter := router.RouterGroupApp.Example
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
		systemRouter.InitApiRouter(PrivateGroup)                 //注册功能api路由
		systemRouter.InitCasbinRouter(PrivateGroup)              //权限相关的路由
		systemRouter.InitUserRouter(PrivateGroup)                //注册用户相关的路由
		systemRouter.InitAuthorityRouter(PrivateGroup)           //注册角色路由
		systemRouter.InitMenuRouter(PrivateGroup)                //菜单路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)       //字典管理
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) //字典详情管理
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)  //操作记录管理

		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) //文件上传，断点续传功能实现
		exampleRouter.InitCustomerRouter(PrivateGroup)
	}
	return Router
}

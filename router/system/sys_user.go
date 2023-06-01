package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)               //注册用户
		userRouter.POST("changePasword", baseApi.ChangePasword)           //更改密码
		userRouter.POST("setUserAuthotity", baseApi.SetUserAuthotity)     //设置用户权限，未发现其是否在用，目前设置权限使用的是用户权限组
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               //删除用户
		userRouter.PUT("setUserInfo", baseApi.SetuserInfo)                //设置用户信息，是否启用该用户
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}

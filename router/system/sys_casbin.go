package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (c *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
}

package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type DictionaryDetailApiRouter struct {
}

func (s *DictionaryDetailApiRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail")
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", sysDictionDetailApi.CreateSysDictionaryDetail)
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sysDictionDetailApi.DeleteSysDictionaryDetail)
		dictionaryDetailRouter.PUT("upadateSysdictionaryDetail", sysDictionDetailApi.UpadateSysdictionaryDetail)
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", sysDictionDetailApi.FindSysDictionaryDetail)
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", sysDictionDetailApi.GetSysDictionaryDetailList)
	}
}

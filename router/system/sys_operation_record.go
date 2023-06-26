package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct {
}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	operationApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", operationApi.CreateSysOperationRecord)
		operationRecordRouter.DELETE("deleteSysOperationRecord", operationApi.DeleteSysOperationRecord)
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", operationApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", operationApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", operationApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}

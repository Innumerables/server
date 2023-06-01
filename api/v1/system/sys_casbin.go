package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

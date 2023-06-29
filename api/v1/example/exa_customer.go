package example

import (
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/example"
	exampleRes "server/model/example/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct{}

func (ca *CustomerApi) CreateCustomer(c *gin.Context) {
	var customer example.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = customerService.CreateCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (ca *CustomerApi) UpdateCustomer(c *gin.Context) {
	var customer example.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.UpdateCustomer(&customer)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (ca *CustomerApi) DeleteCustomer(c *gin.Context) {
	var customer example.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.DeleteCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)

}

func (ca *CustomerApi) GetCustomer(c *gin.Context) {
	var customer example.Customer
	err := c.ShouldBindQuery(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := customerService.GetCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.CustomerResponse{Customer: data}, "获取失败", c)
}

func (ca *CustomerApi) GetCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := customerService.GetCustomerList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

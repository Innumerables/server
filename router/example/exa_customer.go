package example

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (c *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	customerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", customerApi.CreateCustomer)
		customerRouter.PUT("customer", customerApi.UpdateCustomer)
		customerRouter.DELETE("customer", customerApi.DeleteCustomer)
	}
	{
		customerRouterWithoutRecord.GET("customer", customerApi.GetCustomer)
		customerRouterWithoutRecord.GET("customerList", customerApi.GetCustomerList)
	}
}

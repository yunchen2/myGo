package customerRouter

import (
	"github.com/gin-gonic/gin"
	v1 "myGo/api/v1"
	middleware "myGo/operation"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customerRouter").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customerRouter")
	var exaCustomerApi = v1.ApiGroup.ApiGroup.CustomerApi
	{
		customerRouter.POST("customerRouter", api.CreateExaCustomer)              // 创建客户
		customerRouter.PUT("customerRouter", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customerRouter", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customerRouter", exaCustomerApi.GetExaCustomer)   // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}

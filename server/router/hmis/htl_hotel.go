package hmis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HtlHotelRouter struct {
}

// InitHtlHotelRouter 初始化 htlHotel表 路由信息
func (s *HtlHotelRouter) InitHtlHotelRouter(Router *gin.RouterGroup) {
	htlHotelRouter := Router.Group("htlHotel").Use(middleware.OperationRecord())
	htlHotelRouterWithoutRecord := Router.Group("htlHotel")
	var htlHotelApi = v1.ApiGroupApp.HmisApiGroup.HtlHotelApi
	{
		htlHotelRouter.POST("createHtlHotel", htlHotelApi.CreateHtlHotel)             // 新建htlHotel表
		htlHotelRouter.DELETE("deleteHtlHotel", htlHotelApi.DeleteHtlHotel)           // 删除htlHotel表
		htlHotelRouter.DELETE("deleteHtlHotelByIds", htlHotelApi.DeleteHtlHotelByIds) // 批量删除htlHotel表
		htlHotelRouter.PUT("updateHtlHotel", htlHotelApi.UpdateHtlHotel)              // 更新htlHotel表
	}
	{
		htlHotelRouterWithoutRecord.GET("findHtlHotel", htlHotelApi.FindHtlHotel)       // 根据ID获取htlHotel表
		htlHotelRouterWithoutRecord.GET("getHtlHotelList", htlHotelApi.GetHtlHotelList) // 获取htlHotel表列表
	}
}

package hmis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hmis"
	hmisReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmis/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HtlHotelApi struct {
}

var htlHotelService = service.ServiceGroupApp.HmisServiceGroup.HtlHotelService

// CreateHtlHotel 创建htlHotel表
// @Tags HtlHotel
// @Summary 创建htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmis.HtlHotel true "创建htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /htlHotel/createHtlHotel [post]
func (htlHotelApi *HtlHotelApi) CreateHtlHotel(c *gin.Context) {
	var htlHotel hmis.HtlHotel
	err := c.ShouldBindJSON(&htlHotel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := htlHotelService.CreateHtlHotel(&htlHotel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHtlHotel 删除htlHotel表
// @Tags HtlHotel
// @Summary 删除htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmis.HtlHotel true "删除htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /htlHotel/deleteHtlHotel [delete]
func (htlHotelApi *HtlHotelApi) DeleteHtlHotel(c *gin.Context) {
	ID := c.Query("ID")
	if err := htlHotelService.DeleteHtlHotel(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHtlHotelByIds 批量删除htlHotel表
// @Tags HtlHotel
// @Summary 批量删除htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /htlHotel/deleteHtlHotelByIds [delete]
func (htlHotelApi *HtlHotelApi) DeleteHtlHotelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := htlHotelService.DeleteHtlHotelByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHtlHotel 更新htlHotel表
// @Tags HtlHotel
// @Summary 更新htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmis.HtlHotel true "更新htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /htlHotel/updateHtlHotel [put]
func (htlHotelApi *HtlHotelApi) UpdateHtlHotel(c *gin.Context) {
	var htlHotel hmis.HtlHotel
	err := c.ShouldBindJSON(&htlHotel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := htlHotelService.UpdateHtlHotel(htlHotel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHtlHotel 用id查询htlHotel表
// @Tags HtlHotel
// @Summary 用id查询htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmis.HtlHotel true "用id查询htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /htlHotel/findHtlHotel [get]
func (htlHotelApi *HtlHotelApi) FindHtlHotel(c *gin.Context) {
	ID := c.Query("ID")
	if rehtlHotel, err := htlHotelService.GetHtlHotel(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehtlHotel": rehtlHotel}, c)
	}
}

// GetHtlHotelList 分页获取htlHotel表列表
// @Tags HtlHotel
// @Summary 分页获取htlHotel表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmisReq.HtlHotelSearch true "分页获取htlHotel表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htlHotel/getHtlHotelList [get]
func (htlHotelApi *HtlHotelApi) GetHtlHotelList(c *gin.Context) {
	var pageInfo hmisReq.HtlHotelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := htlHotelService.GetHtlHotelInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

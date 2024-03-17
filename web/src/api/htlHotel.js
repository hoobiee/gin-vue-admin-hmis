import service from '@/utils/request'

// @Tags HtlHotel
// @Summary 创建htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtlHotel true "创建htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /htlHotel/createHtlHotel [post]
export const createHtlHotel = (data) => {
  return service({
    url: '/htlHotel/createHtlHotel',
    method: 'post',
    data
  })
}

// @Tags HtlHotel
// @Summary 删除htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtlHotel true "删除htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /htlHotel/deleteHtlHotel [delete]
export const deleteHtlHotel = (params) => {
  return service({
    url: '/htlHotel/deleteHtlHotel',
    method: 'delete',
    params
  })
}

// @Tags HtlHotel
// @Summary 批量删除htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /htlHotel/deleteHtlHotel [delete]
export const deleteHtlHotelByIds = (params) => {
  return service({
    url: '/htlHotel/deleteHtlHotelByIds',
    method: 'delete',
    params
  })
}

// @Tags HtlHotel
// @Summary 更新htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtlHotel true "更新htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /htlHotel/updateHtlHotel [put]
export const updateHtlHotel = (data) => {
  return service({
    url: '/htlHotel/updateHtlHotel',
    method: 'put',
    data
  })
}

// @Tags HtlHotel
// @Summary 用id查询htlHotel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HtlHotel true "用id查询htlHotel表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /htlHotel/findHtlHotel [get]
export const findHtlHotel = (params) => {
  return service({
    url: '/htlHotel/findHtlHotel',
    method: 'get',
    params
  })
}

// @Tags HtlHotel
// @Summary 分页获取htlHotel表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取htlHotel表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htlHotel/getHtlHotelList [get]
export const getHtlHotelList = (params) => {
  return service({
    url: '/htlHotel/getHtlHotelList',
    method: 'get',
    params
  })
}

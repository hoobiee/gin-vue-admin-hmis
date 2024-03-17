package hmis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hmis"
	hmisReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmis/request"
)

type HtlHotelService struct {
}

// CreateHtlHotel 创建htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) CreateHtlHotel(htlHotel *hmis.HtlHotel) (err error) {
	err = global.MustGetGlobalDBByDBName("hmis").Create(htlHotel).Error
	return err
}

// DeleteHtlHotel 删除htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) DeleteHtlHotel(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("hmis").Delete(&hmis.HtlHotel{}, "id = ?", ID).Error
	return err
}

// DeleteHtlHotelByIds 批量删除htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) DeleteHtlHotelByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("hmis").Delete(&[]hmis.HtlHotel{}, "id in ?", IDs).Error
	return err
}

// UpdateHtlHotel 更新htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) UpdateHtlHotel(htlHotel hmis.HtlHotel) (err error) {
	err = global.MustGetGlobalDBByDBName("hmis").Save(&htlHotel).Error
	return err
}

// GetHtlHotel 根据ID获取htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) GetHtlHotel(ID string) (htlHotel hmis.HtlHotel, err error) {
	err = global.MustGetGlobalDBByDBName("hmis").Where("id = ?", ID).First(&htlHotel).Error
	return
}

// GetHtlHotelInfoList 分页获取htlHotel表记录
// Author [piexlmax](https://github.com/piexlmax)
func (htlHotelService *HtlHotelService) GetHtlHotelInfoList(info hmisReq.HtlHotelSearch) (list []hmis.HtlHotel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("hmis").Model(&hmis.HtlHotel{})
	var htlHotels []hmis.HtlHotel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&htlHotels).Error
	return htlHotels, total, err
}

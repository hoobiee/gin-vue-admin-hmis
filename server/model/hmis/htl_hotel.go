// 自动生成模板HtlHotel
package hmis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// htlHotel表 结构体  HtlHotel
type HtlHotel struct {
	global.GVA_MODEL
	Address          string     `json:"address" form:"address" gorm:"column:address;comment:酒店地址;size:255;"`                             //酒店地址
	AddressEn        string     `json:"addressEn" form:"addressEn" gorm:"column:address_en;comment:酒店英文地址;size:255;"`                    //酒店英文地址
	BaiduLat         string     `json:"baiduLat" form:"baiduLat" gorm:"column:baidu_lat;comment:百度维度;size:128;"`                         //百度维度
	BaiduLon         string     `json:"baiduLon" form:"baiduLon" gorm:"column:baidu_lon;comment:百度经度;size:128;"`                         //百度经度
	CanclePolicy     string     `json:"canclePolicy" form:"canclePolicy" gorm:"column:cancle_policy;comment:取消政策;size:255;"`             //取消政策
	CityName         string     `json:"cityName" form:"cityName" gorm:"column:city_name;comment:城市名称;size:255;"`                         //城市名称
	CreateBy         string     `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;size:255;"`                          //创建人
	CreateTime       *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                            //创建时间
	CreditCards      string     `json:"creditCards" form:"creditCards" gorm:"column:credit_cards;comment:支持的信用卡;size:255;"`              //支持的信用卡
	DecorateTime     *time.Time `json:"decorateTime" form:"decorateTime" gorm:"column:decorate_time;comment:装修时间;"`                      //装修时间
	DestinationLabel *int       `json:"destinationLabel" form:"destinationLabel" gorm:"column:destination_label;comment:目的地标签;size:10;"` //目的地标签
	Facilities       string     `json:"facilities" form:"facilities" gorm:"column:facilities;comment:酒店设施列表;size:255;"`                  //酒店设施列表
	GaodeLat         string     `json:"gaodeLat" form:"gaodeLat" gorm:"column:gaode_lat;comment:高德经度;size:128;"`                         //高德经度
	GaodeLon         string     `json:"gaodeLon" form:"gaodeLon" gorm:"column:gaode_lon;comment:高德维度;size:128;"`                         //高德维度
	GoogleLat        string     `json:"googleLat" form:"googleLat" gorm:"column:google_lat;comment:谷歌维度;size:128;"`                      //谷歌维度
	GoogleLon        string     `json:"googleLon" form:"googleLon" gorm:"column:google_lon;comment:谷歌经度;size:128;"`                      //谷歌经度
	HotelId          *int       `json:"hotelId" form:"hotelId" gorm:"primarykey;column:hotel_id;comment:酒店编号;size:10;"`                  //酒店编号
	HotelName        string     `json:"hotelName" form:"hotelName" gorm:"column:hotel_name;comment:酒店名称;size:255;"`                      //酒店名称
	HotelNameEn      string     `json:"hotelNameEn" form:"hotelNameEn" gorm:"column:hotel_name_en;comment:英文名;size:255;"`                //英文名
	HotelPhone       string     `json:"hotelPhone" form:"hotelPhone" gorm:"column:hotel_phone;comment:电话;size:128;"`                     //电话
	Introduction     string     `json:"introduction" form:"introduction" gorm:"column:introduction;comment:酒店介绍;size:2048;"`             //酒店介绍
	OpeningTime      *time.Time `json:"openingTime" form:"openingTime" gorm:"column:opening_time;comment:开业时间;"`                         //开业时间
	Star             *int       `json:"star" form:"star" gorm:"column:star;comment:星级(1/2/3/4/5);size:10;"`                              //星级(1/2/3/4/5)
	UpdateBy         string     `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改人;size:255;"`                          //修改人
	UpdateTime       *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`                            //修改时间
}

// TableName htlHotel表 HtlHotel自定义表名 htl_hotel
func (HtlHotel) TableName() string {
	return "htl_hotel"
}

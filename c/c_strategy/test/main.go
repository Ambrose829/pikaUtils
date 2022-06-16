package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
  @Author: pika
  @Date: 2022/4/14 12:08 下午
  @Description:
*/

//func main() {
//
//	var cf *impl.CityFilter
//
//	var bf *impl.BizFilter
//
//	cf = impl.NewCityFilter(*filter.NewBaseFilter(filter.EQ, "北京"))
//
//	bf = impl.NewBizFilter(*filter.NewBaseFilter(filter.UEQ, "快车"))
//
//	var fc impl.FilterChan
//	fc.AddFilter(cf)
//	fc.AddFilter(bf)
//	fc.SetTarget("城市业、务线过滤器链")
//	s := vo.Strategy{
//		CityList: "北京",
//		BizTypes: "快车",
//	}
//	pass := fc.Execute(s)
//	fmt.Println(pass)
//
//	//mapp := make(map[int][]int, 1)
//	//
//	//mapp[1] = []int{1,2, 3}
//	//mapp[1] = nil
//	//_, ok := mapp[1]
//	//
//	//fmt.Println(ok)
//
//
//
//
//
//}

func main() {
	r := ResourceState{
		Id:               0,
		Status:           0,
		TemplateType:     0,
		SourceId:         0,
		Location:         "",
		Business:         0,
		Name:             "",
		CreateUserName:   "",
		ModifiedUserName: "",
		GmtCreate:        time.Time{},
		GmtModified:      time.Time{},
	}
	m := MaterialR{
		Id:               0,
		Status:           0,
		MaterialId:       0,
		ResourceStateId:  0,
		MaterialType:     0,
		BuiltInType:      0,
		TemplateType:     0,
		HighlightType:    0,
		HighlightInfo:    "",
		Name:             "",
		Title:            "",
		CoverUrl:         "",
		BfClickImgUrl:    "",
		AfClickImgUrl:    "",
		Ext:              "",
		CreateUserName:   "",
		ModifiedUserName: "",
		GmtCreate:        time.Time{},
		GmtModified:      time.Time{},
	}

	res := []ResInfo{{
		ResourceStateData: r,
		MaterialData:      m,
	},
		{
			ResourceStateData: r,
			MaterialData:      m,
		}}

	b, _ := json.Marshal(res)
	fmt.Println(string(b))

}

type ResInfo struct {
	ResourceStateData ResourceState `json:"resourceState"`
	MaterialData      MaterialR     `json:"materialR"`
}

type ResourceState struct {
	Id               int64     `orm:"id" json:"id"`                               // 主键id
	Status           int       `orm:"status" json:"status"`                       // 状态：-1:删除,1:生效,2:失效
	TemplateType     int       `orm:"template_type" json:"templateType"`          // 模板类型,1:底导、2:首页1
	SourceId         int64     `orm:"source_id" json:"sourceId"`                  // 源id
	Location         string    `orm:"location" json:"location"`                   // 在源中的位置
	Business         int       `orm:"business" json:"business"`                   // 业务方:1:平台运营,2:内容运营
	Name             string    `orm:"name" json:"name"`                           // 名称
	CreateUserName   string    `orm:"create_user_name" json:"createUserName"`     // 新增用户名
	ModifiedUserName string    `orm:"modified_user_name" json:"modifiedUserName"` // 更新用户名
	GmtCreate        time.Time `orm:"gmt_create" json:"gmtCreate"`                // 创建时间
	GmtModified      time.Time `orm:"gmt_modified" json:"gmtModified"`            // 更新时间
}

type MaterialR struct {
	Id               int64     `orm:"id" json:"id"`                               // 主键id
	Status           int       `orm:"status" json:"status"`                       // 状态：删除:-1,暂存:1,审核中:2,已驳回:3,审核通过:4,已撤回:5,灰度:6,上线:7,下线:8
	MaterialId       int64     `orm:"material_id" json:"materialId"`              // 引用物料id
	ResourceStateId  int64     `orm:"resource_state_id" json:"resourceStateId"`   // 资源位id
	MaterialType     int       `orm:"material_type" json:"materialType"`          // 素材类型,
	BuiltInType      int       `orm:"built_in_type" json:"builtInType"`           // 素材类型,
	TemplateType     int       `orm:"template_type" json:"templateType"`          // 模板类型,1:底导、2:首页1
	HighlightType    int       `orm:"highlight_type" json:"highlightType"`        // 强调类型，0:无;1:红点;2:气泡
	HighlightInfo    string    `orm:"highlight_info" json:"highlightInfo"`        // 强调内容
	Name             string    `orm:"name" json:"name"`                           // 物料名称
	Title            string    `orm:"title" json:"title"`                         // 文案内容
	CoverUrl         string    `orm:"cover_url" json:"coverUrl"`                  // 封面图地址
	BfClickImgUrl    string    `orm:"bf_click_img_url" json:"bfClickImgUrl"`      // 特殊标题点击前图片
	AfClickImgUrl    string    `orm:"af_click_img_url" json:"afClickImgUrl"`      // 特殊标题点击后图片
	Ext              string    `orm:"ext" json:"ext"`                             // 特殊标题点击后图片
	CreateUserName   string    `orm:"create_user_name" json:"createUserName"`     // 新增用户名
	ModifiedUserName string    `orm:"modified_user_name" json:"modifiedUserName"` // 更新用户名
	GmtCreate        time.Time `orm:"gmt_create" json:"gmtCreate"`                // 创建时间
	GmtModified      time.Time `orm:"gmt_modified" json:"gmtModified"`            // 更新时间
}

package impl

import (
	"pikaUtils/design-combination/model/vo"
	"pikaUtils/design-combination/service/logic"
)

/**
  @Author: pika
  @Date: 2022/4/1 11:00 上午
  @Description: 用户性别过滤器
*/

type UserGenderFilter struct {
	bl logic.BaseLogic
}

func (ugf UserGenderFilter) Filter(matterValue string, treeNodeLineInfoList []vo.TreeNodeLink) int64 {
	return ugf.bl.Filter(matterValue, treeNodeLineInfoList)
}

func (ugf UserGenderFilter) MatterValue(treeId int64, userId string, decisionMatter map[string]string) string {
	// 可以从数据库、RPC接口、缓存等方式获取
	return decisionMatter["gender"]
}

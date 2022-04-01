package impl

import (
	"pikaUtils/design-combination/model/vo"
	"pikaUtils/design-combination/service/logic"
)

/**
  @Author: pika
  @Date: 2022/4/1 10:59 上午
  @Description: 用户年龄过滤
*/


type UserAgeFilter struct {
	bl logic.BaseLogic
}

func (uaf UserAgeFilter) Filter(matterValue string, treeNodeLineInfoList []vo.TreeNodeLink) int64  {
	return uaf.bl.Filter(matterValue, treeNodeLineInfoList)
}

func (uaf UserAgeFilter) MatterValue(treeId int64, userId string, decisionMatter map[string]string) string  {
	return decisionMatter["age"]
}
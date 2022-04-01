package logic

import (
	"pikaUtils/design-combination/model/vo"
	"pikaUtils/util"
	"strings"
)

/**
  @Author: pika
  @Date: 2022/3/31 7:34 下午
  @Description: 基础策略
*/

const (
	// EQ 等于
	EQ = 1
	// GT 大于
	GT = 2
	// LT 小于
	LT = 3
	// GTE 大于等于
	GTE = 4
	// LTE 小于等于
	LTE = 5
)

type BaseLogic struct {
}

func (bl BaseLogic) Filter(matterValue string, treeNodeLineInfoList []vo.TreeNodeLink) int64 {
	for _, nodeLink := range treeNodeLineInfoList {
		if bl.decisionLogic(matterValue, nodeLink) {
			return nodeLink.NodeIdTo()
		}
	}
	return 0
}

func (bl BaseLogic) MatterValue(treeId int64, userId string, decisionMatter map[string]string) string {
	return ""
}

//
// decisionLogic
// @author: pika
// @Description: 规则选择
// @receiver bl 基础策略
// @param matterValue 决策值
// @param nodeLink 决策链节点
// @return bool 决策节点返回值
//
func (bl BaseLogic) decisionLogic(matterValue string, nodeLink vo.TreeNodeLink) bool {
	// 验证规则符号限定类型
	switch nodeLink.RuleLimitType() {
	// =
	case EQ:
		return strings.EqualFold(matterValue, nodeLink.RuleLimitValue())
	// >
	case GT:
		return util.Str2F64(matterValue) > util.Str2F64(nodeLink.RuleLimitValue())
	// <
	case LT:
		return util.Str2F64(matterValue) < util.Str2F64(nodeLink.RuleLimitValue())
	// >=
	case GTE:
		return util.Str2F64(matterValue) >= util.Str2F64(nodeLink.RuleLimitValue())
	// <=
	case LTE:
		return util.Str2F64(matterValue) <= util.Str2F64(nodeLink.RuleLimitValue())
	default:
		return false
	}
}

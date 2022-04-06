package filter

import (
	"pikaUtils/czp/c_strategy/model/vo"
	"pikaUtils/util"
	"strings"
)

/**
  @Author: pika
  @Date: 2022/4/2 3:44 下午
  @Description: 基础过滤器
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
	// UEQ 不等于
	UEQ = 6
	// IN 属于
	IN = 7
	// NIN 不属于
	NIN = 8
	// INC 包含
	INC = 9
	// EXC 不包含
	EXC = 10
)

type BaseFilter struct {
	ruleLimitType  int    //限定类型；1:=;2:>;3:<;4:>=;5:<=;6:!=;7:in;8:notin;9:enum[枚举范围] 也可推出限定值类型
	ruleLimitValue string //限定值
}

func (bf BaseFilter) Filter(matterValue string, chanNodeLineList []vo.ChanNodeLink) bool {
	for _, nodeLink := range chanNodeLineList {
		if !bf.decisionLogic(matterValue, nodeLink) {
			return false
		}
	}
	return true
}

func (bf BaseFilter) MatterValue(strategy vo.Strategy) string {
	return ""
}

func (bf BaseFilter) decisionLogic(matterValue string, nodeLink vo.ChanNodeLink) bool {
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
	// !=
	case UEQ:
		return !strings.EqualFold(matterValue, nodeLink.RuleLimitValue())
	// in
	case IN:
		// 限制值包含要决策的值
		list := strings.Split(nodeLink.RuleLimitValue(), ",")
		for _, rlv := range list {
			isPass := strings.EqualFold(matterValue, rlv)
			// 包含决策值
			if isPass {
				return true
			}
		}

		// 不包含决策值
		return false

	// not in
	case NIN:
		// 限制值不包含要决策的值
		list := strings.Split(nodeLink.RuleLimitValue(), ",")
		for _, rlv := range list {
			isPass := strings.EqualFold(matterValue, rlv)
			// 包含决策值
			if isPass {
				return false
			}
		}

		// 不包含决策值
		return true
	// include
	case INC:
		// 要决策的值包含限制值
		list := strings.Split(matterValue, ",")
		for _, mv := range list {
			isPass := strings.EqualFold(mv, nodeLink.RuleLimitValue())
			//	如果包含限制值
			if isPass {
				return true
			}
		}

		// 不包含限制值
		return false
	// exclusive
	case EXC:
		// 要决策的值不包含限制值
		list := strings.Split(matterValue, ",")
		for _, mv := range list {
			isPass := strings.EqualFold(mv, nodeLink.RuleLimitValue())
			//	如果包含限制值
			if isPass {
				return false
			}
		}
		// 不包含限制值
		return true

	default:
		return false
	}

}

func (bf *BaseFilter) RuleLimitType() int {
	return bf.ruleLimitType
}

func (bf *BaseFilter) SetRuleLimitType(ruleLimitType int) {
	bf.ruleLimitType = ruleLimitType
}

func (bf *BaseFilter) RuleLimitValue() string {
	return bf.ruleLimitValue
}

func (bf *BaseFilter) SetRuleLimitValue(ruleLimitValue string) {
	bf.ruleLimitValue = ruleLimitValue
}

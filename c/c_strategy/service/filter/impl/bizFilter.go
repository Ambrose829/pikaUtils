package impl

import (
	"pikaUtils/c/c_strategy/model/vo"
	"pikaUtils/c/c_strategy/service/filter"
)

/**
  @Author: pika
  @Date: 2022/4/14 11:52 上午
  @Description: 业务线过滤器
*/

type BizFilter struct {
	baseFilter filter.BaseFilter
}

func NewBizFilter(baseFilter filter.BaseFilter) *BizFilter {
	return &BizFilter{baseFilter: baseFilter}
}

func (bf *BizFilter) BaseFilter() filter.BaseFilter {
	return bf.baseFilter
}

func (bf *BizFilter) SetBaseFilter(baseFilter filter.BaseFilter) {
	bf.baseFilter = baseFilter
}

// Filter 验证业务线规则
func (bf BizFilter) Filter(strategy vo.Strategy) bool {
	// 验证是否通过
	return bf.baseFilter.Filter(bf.MatterValue(strategy))
}

// MatterValue 获取业务线过滤器需要验证的值, 业务线
func (bf BizFilter) MatterValue(strategy vo.Strategy) string {
	return strategy.BizTypes
}

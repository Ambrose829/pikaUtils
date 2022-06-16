package impl

import (
	"pikaUtils/c/s/filter"
	"pikaUtils/c/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 2:39 下午
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
func (bf BizFilter) Filter(strategy strategy.Strategy) bool {
	// 验证是否通过
	return bf.baseFilter.Filter(bf.MatterValue(strategy))
}

// MatterValue 获取业务线过滤器需要验证的值, 业务线
func (bf BizFilter) MatterValue(strategy strategy.Strategy) string {
	return strategy.BizTypes
}

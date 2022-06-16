package impl

import (
	"pikaUtils/c/s/filter"
	"pikaUtils/c/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 10:18 上午
  @Description: 过滤器链
*/

// FilterChan 过滤器链
type FilterChan struct {
	filters []filter.LogicFilter
	Target  string //目的
}

// AddFilter 增加过滤器
func (fc *FilterChan) AddFilter(filter filter.LogicFilter) {
	fc.filters = append(fc.filters, filter)
}

// 初始化目的
func (fc *FilterChan) SetTarget(target string) {
	fc.Target = target
}

// Execute 处理过滤器
func (fc *FilterChan) Execute(strategy strategy.Strategy) bool {

	for _, filter := range fc.filters {
		// 验证过滤器链中的过滤器过滤
		isSuccess := filter.Filter(strategy)

		// 未通过决策
		if !isSuccess {
			return false
		}
	}
	return true
}

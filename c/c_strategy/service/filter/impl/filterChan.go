package impl

import (
	"pikaUtils/c/c_strategy/model/vo"
	"pikaUtils/c/c_strategy/service/filter"
)

/**
  @Author: pika
  @Date: 2022/4/2 2:27 下午
  @Description:
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
func (fc *FilterChan) Execute(strategy vo.Strategy) bool {

	for _, filter := range fc.filters {
		// todo 注释
		isSuccess := filter.Filter(strategy)

		// 未通过决策
		if !isSuccess {
			return false
		}
	}
	return true
}

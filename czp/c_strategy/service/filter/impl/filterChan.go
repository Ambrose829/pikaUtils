package impl

import (
	"pikaUtils/czp/c_strategy/model/vo"
	"pikaUtils/czp/c_strategy/service/filter"
)

/**
  @Author: pika
  @Date: 2022/4/2 2:27 下午
  @Description:
*/

// FilterChan 过滤器链
type FilterChan struct {
	filters []filter.Filter
	Target  string //目的
}

// AddFilter 增加过滤器
func (fc FilterChan) AddFilter(filter filter.Filter) {
	fc.filters = append(fc.filters, filter)
}

// 初始化目的
func (fc FilterChan) SetTarget(target string) {
	fc.Target = target
}

// Execute 处理过滤器
func (fc FilterChan) Execute(strategy vo.Strategy, chanNodeLineList []vo.ChanNodeLink) bool {

	for _, filter := range fc.filters {
		// todo 注释
		isSuccess := filter.Execute(filter.MatterValue(strategy), chanNodeLineList)

		// 未通过决策
		if !isSuccess {
			return false
		}
	}
	return true
}

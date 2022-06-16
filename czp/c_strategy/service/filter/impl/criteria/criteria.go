package criteria

import "pikaUtils/czp/c_strategy/service/filter"

/**
  @Author: pika
  @Date: 2022/4/2 12:12 下午
  @Description:
*/

// Criteria 规则
type Criteria struct {
	filters []filter.LogicFilter
}

func (ctr *Criteria) Filters() []filter.LogicFilter {
	return ctr.filters
}

func (ctr *Criteria) SetFilters(filters ...filter.LogicFilter) {
	ctr.filters = filters
}

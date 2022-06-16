package criteria

import "pikaUtils/czp/s/filter"

/**
  @Author: pika
  @Date: 2022/4/19 10:25 上午
  @Description: 过滤器聚合
*/

// Criteria 规则
type Criteria struct {
	filters []filter.LogicFilter
}

func NewCriteria(filters ...filter.LogicFilter) *Criteria {
	return &Criteria{filters: filters}
}

func (ctr *Criteria) Filters() []filter.LogicFilter {
	return ctr.filters
}

func (ctr *Criteria) SetFilters(filters ...filter.LogicFilter) {
	ctr.filters = filters
}

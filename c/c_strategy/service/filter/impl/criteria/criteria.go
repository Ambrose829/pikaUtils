package criteria

import "pikaUtils/c/c_strategy/service/filter"

/**
  @Author: pika
  @Date: 2022/4/2 12:12 δΈε
  @Description:
*/

// Criteria θ§ε
type Criteria struct {
	filters []filter.LogicFilter
}

func (ctr *Criteria) Filters() []filter.LogicFilter {
	return ctr.filters
}

func (ctr *Criteria) SetFilters(filters ...filter.LogicFilter) {
	ctr.filters = filters
}

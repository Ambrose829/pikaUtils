package criteria

import (
	"pikaUtils/c/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 10:27 上午
  @Description: 或规则
*/

// OrCriteria 或规则
type OrCriteria struct {
	ctr Criteria
}

func NewOrCriteria(ctr Criteria) *OrCriteria {
	return &OrCriteria{ctr: ctr}
}

func (oc *OrCriteria) Filter(strategy strategy.Strategy) bool {
	for _, f := range oc.ctr.Filters() {
		isPass := f.Filter(strategy)
		if isPass {
			return true
		}
	}
	return false
}

func (oc *OrCriteria) MatterValue(strategy strategy.Strategy) string {
	return ""
}

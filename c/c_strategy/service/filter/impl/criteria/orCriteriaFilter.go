package criteria

import (
	"pikaUtils/c/c_strategy/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/1 8:55 下午
  @Description: 或规则
*/

// OrCriteria 或规则
type OrCriteria struct {
	ctr Criteria
}

func (oc OrCriteria) Filter(strategy vo.Strategy) bool {
	for _, f := range oc.ctr.Filters() {
		isPass := f.Filter(strategy)
		if isPass {
			return true
		}
	}
	return false
}

func (oc OrCriteria) MatterValue(strategy vo.Strategy) string {
	return ""
}

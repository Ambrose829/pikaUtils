package criteria

import (
	"pikaUtils/c/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 10:26 上午
  @Description: 与规则
*/

// AndCriteria 与规则
type AndCriteria struct {
	ctr Criteria
}

func (ac *AndCriteria) Filter(strategy strategy.Strategy) bool {
	for _, f := range ac.ctr.Filters() {
		isPass := f.Filter(strategy)
		if !isPass {
			return false
		}
	}
	return true
}

func (ac *AndCriteria) MatterValue(strategy strategy.Strategy) string {
	return ""
}

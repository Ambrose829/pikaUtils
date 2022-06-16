package criteria

import (
	"pikaUtils/czp/c_strategy/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/1 8:54 下午
  @Description: 与规则
*/

// AndCriteria 与规则
type AndCriteria struct {
	ctr Criteria
}

func (ac AndCriteria) Filter(strategy vo.Strategy) bool {
	for _, f := range ac.ctr.Filters() {
		isPass := f.Filter(strategy)
		if !isPass {
			return false
		}
	}
	return true
}

func (ac AndCriteria) MatterValue(strategy vo.Strategy) string {
	return ""
}

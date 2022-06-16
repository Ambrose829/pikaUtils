package impl

import (
	"pikaUtils/czp/c_strategy/model/vo"
	"pikaUtils/czp/c_strategy/service/filter"
)

/**
  @Author: pika
  @Date: 2022/4/1 8:48 下午
  @Description: 城市过滤器
*/

type CityFilter struct {
	baseFilter filter.BaseFilter
}

func NewCityFilter(baseFilter filter.BaseFilter) *CityFilter {
	return &CityFilter{baseFilter: baseFilter}
}

func (cf *CityFilter) BaseFilter() filter.BaseFilter {
	return cf.baseFilter
}

func (cf *CityFilter) SetBaseFilter(baseFilter filter.BaseFilter) {
	cf.baseFilter = baseFilter
}

func (cf CityFilter) Filter(strategy vo.Strategy) bool {
	// 验证是否通过
	return cf.baseFilter.Filter(cf.MatterValue(strategy))
}

func (cf CityFilter) MatterValue(strategy vo.Strategy) string {
	return strategy.CityList
}

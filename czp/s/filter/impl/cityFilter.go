package impl

import (
	"pikaUtils/czp/s/filter"
	"pikaUtils/czp/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 2:38 下午
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

func (cf CityFilter) Filter(strategy strategy.Strategy) bool {
	// 验证是否通过
	return cf.baseFilter.Filter(cf.MatterValue(strategy))
}

func (cf CityFilter) MatterValue(strategy strategy.Strategy) string {
	return strategy.CityList
}

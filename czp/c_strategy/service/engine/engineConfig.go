package engine

//
//import (
//	"pikaUtils/czp/c_strategy/service/filter"
//	"pikaUtils/czp/c_strategy/service/filter/impl"
//)
//
///**
//  @Author: pika
//  @Date: 2022/4/2 7:20 下午
//  @Description: 引擎配置
//*/
//
//type EngineConfig struct {
//	filterMap map[string]filter.Filter
//}
//
//func (ec *EngineConfig) FilterMap() map[string]filter.Filter {
//	return ec.filterMap
//}
//
//func (ec *EngineConfig) SetFilterMap(filterMap map[string]filter.Filter) {
//	ec.filterMap = filterMap
//}
//
//var (
//	// EC 将决策节点配置到map结构中，也可以将这样的map结构抽取到数据库中
//	EC = EngineConfig{filterMap: map[string]filter.Filter{"city": impl.CityFilter{}}}
//)

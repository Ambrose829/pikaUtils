package engine

import (
	"pikaUtils/design-combination/service/logic"
	"pikaUtils/design-combination/service/logic/impl"
)

/**
  @Author: pika
  @Date: 2022/4/1 1:14 下午
  @Description: 引擎配置
*/

type EngineConfig struct {
	logicFilterMap map[string]logic.LogicFilter
}

func (ec *EngineConfig) LogicFilterMap() map[string]logic.LogicFilter {
	return ec.logicFilterMap
}

func (ec *EngineConfig) SetLogicFilterMap(logicFilterMap map[string]logic.LogicFilter) {
	ec.logicFilterMap = logicFilterMap
}

var (
	// EC 将决策节点配置到map结构中，也可以将这样的map结构抽取到数据库中
	EC = EngineConfig{logicFilterMap: map[string]logic.LogicFilter{"userAge": impl.UserAgeFilter{}, "userGender": impl.UserGenderFilter{}}}
)

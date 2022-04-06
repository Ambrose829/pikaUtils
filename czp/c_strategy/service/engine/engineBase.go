package engine

import (
	"pikaUtils/czp/c_strategy/service/filter/impl"
	"pikaUtils/design-combination/model/aggregates"
	"pikaUtils/design-combination/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/2 2:55 下午
  @Description: 基础引擎
*/

type EngineBase struct {
	IEngine
}

func (eb EngineBase) EngineDecisionMaker(filterChan impl.FilterChan, chanRich aggregates.ChanRich, decisionMatter []vo.Strategy) vo.TreeNode {
}

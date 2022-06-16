package engine

import (
	"pikaUtils/c/c_strategy/model/vo"
	"pikaUtils/c/c_strategy/service/filter/impl"
)

/**
  @Author: pika
  @Date: 2022/4/2 2:55 下午
  @Description: 基础引擎
*/

type EngineBase struct {
	IEngine
}

func (eb EngineBase) EngineDecisionMaker(filterChan impl.FilterChan, decisionMatter []vo.Strategy) (passStrategy []vo.Strategy) {
	for _, strategy := range decisionMatter {
		pass := filterChan.Execute(strategy)
		if pass {
			passStrategy = append(passStrategy, strategy)
		}
	}
	return
}

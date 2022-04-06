package engine

import (
	"pikaUtils/czp/c_strategy/model/aggregates"
	"pikaUtils/czp/c_strategy/model/vo"
	"pikaUtils/czp/c_strategy/service/filter/impl"
)

/**
  @Author: pika
  @Date: 2022/4/2 2:45 下午
  @Description: 引擎抽象
*/

type IEngine interface {

	//
	// Process
	// @author: pika
	// @Description: 决策处理
	// @param chanRich 聚合规则图
	// @param decisionMatter 需要决策的值
	//

	Process(filterChan impl.FilterChan, chanRich aggregates.ChanRich, decisionMatter []vo.Strategy)
}

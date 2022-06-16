package engine

import (
	"pikaUtils/czp/s/filter/impl"
	"pikaUtils/czp/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/19 11:55 上午
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

	Process(filterChan impl.FilterChan, chanRich int, decisionMatter []strategy.Strategy)
}

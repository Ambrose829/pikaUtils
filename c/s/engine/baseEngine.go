package engine

import (
	"pikaUtils/c/s/filter/impl"
	"pikaUtils/c/s/strategy"
	"pikaUtils/c/s/vo"
	"sort"
)

/**
  @Author: pika
  @Date: 2022/4/19 12:01 下午
  @Description: 基础策略引擎
*/

type BaseEngine struct {
	IEngine
}

//
// EngineDecisionMaker
// @author: pika
// @Description: 决策方法
// @receiver be
// @param filterChan 过滤器链
// @param decisionMatter 需要决策的策略列表
// @return strategy.Strategy 决策唯一
//
func (be BaseEngine) EngineDecisionMaker(filterChan impl.FilterChan, decisionMatter []strategy.Strategy) strategy.Strategy {

	//如果决策策略为空返回
	if len(decisionMatter) <= 0 {
		return strategy.Strategy{}
	}

	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			// todo 打印错误信息
		}
	}()
	var passStrategies vo.StgSlice

	for _, strategy := range decisionMatter {
		pass := filterChan.Execute(strategy)
		if pass {
			passStrategies = append(passStrategies, strategy)
		}
	}

	// 对通过验证的策略进行优先级排序
	sort.Sort(passStrategies)

	// 返回优先级最高的策略
	if len(passStrategies) > 0 {
		return passStrategies[0]
	}

	return strategy.Strategy{}

}

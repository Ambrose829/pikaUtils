package impl

import (
	"pikaUtils/design-combination/model/aggregates"
	"pikaUtils/design-combination/model/vo"
	"pikaUtils/design-combination/service/engine"
)

/**
  @Author: pika
  @Date: 2022/4/1 3:02 下午
  @Description: 决策引擎的实现
*/

type TreeEngineHandle struct {
	eb engine.EngineBase
}

func (teh TreeEngineHandle) Process(treeId int64, userId string, treeRich aggregates.TreeRich, decisionMatter map[string]string) vo.EngineResult {
	// 决策流程, 得出结果
	resTreeNode := teh.eb.EngineDecisionMaker(treeRich, treeId, userId, decisionMatter)

	return *vo.NewEngineResult(userId, treeId, resTreeNode.TreeNodeId(), resTreeNode.NodeValue())
}

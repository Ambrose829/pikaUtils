package engine

import (
	"pikaUtils/design-combination/model/aggregates"
	"pikaUtils/design-combination/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/1 1:13 下午
  @Description: 基础引擎
*/

type EngineBase struct {
	IEngine
}

func engineDecisionMaker(treeRich aggregates.TreeRich, treeId int64, userId string, decisionMatter map[string]string) vo.TreeNode {
	// 拿到根节点
	treeRoot := treeRich.TreeRoot()
	// 拿到规则树map
	treeNodeMap := treeRich.TreeNodeMap()
	// 规则树根ID
	rootNodeId := treeRoot.TreeNodeId()

	treeNodeInfo, _ := treeNodeMap[rootNodeId]
	//节点类型[NodeType]: 1子叶、2果实
	for treeNodeInfo.NodeType() == vo.NodeTypeLeaf {

		// 规则key
		ruleKey := treeNodeInfo.RuleKey()
		// 根据key取得相应的
		logicFilter := EC.logicFilterMap[ruleKey]
	}
}
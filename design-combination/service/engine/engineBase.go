package engine

import (
	"fmt"
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

//
// EngineDecisionMaker
// @author: pika
// @Description: 决策树处理流程
// @receiver eb 基础引擎
// @param treeRich 聚合规则树
// @param treeId 决策树id
// @param userId 用户id
// @param decisionMatter 决策值k-v map池
// @return vo.TreeNode
//

func (eb EngineBase) EngineDecisionMaker(treeRich aggregates.TreeRich, treeId int64, userId string, decisionMatter map[string]string) vo.TreeNode {

	// 拿到根节点
	treeRoot := treeRich.TreeRoot()

	// 拿到规则树map
	treeNodeMap := treeRich.TreeNodeMap()
	// 规则树根ID
	rootNodeId := treeRoot.TreeRootNodeId()

	treeNodeInfo, _ := treeNodeMap[rootNodeId]

	//节点类型[NodeType]: 1子叶、2果实
	for treeNodeInfo.NodeType() == vo.NodeTypeLeaf {

		// 规则key
		ruleKey := treeNodeInfo.RuleKey()
		// 根据key取得相应的过滤器
		logicFilter := EC.logicFilterMap[ruleKey]
		// 根据过滤器不同,拿到相应的需要决策的值
		matterValue := logicFilter.MatterValue(treeId, userId, decisionMatter)

		// 过滤器链一个节点过滤
		nextNodeId := logicFilter.Filter(matterValue, treeNodeInfo.TreeNodeLinkList())
		treeNodeInfo = treeNodeMap[nextNodeId]

		fmt.Printf("决策树引擎=>%s userId：%s treeId：%d treeNodeId：%d ruleKey：%s matterValue：%s",
			treeRoot.TreeName(), userId, treeId, treeNodeInfo.TreeNodeId(), ruleKey, matterValue)
	}
	return treeNodeInfo
}

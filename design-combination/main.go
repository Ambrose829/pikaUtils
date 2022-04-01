package main

import (
	"encoding/json"
	"fmt"
	"pikaUtils/design-combination/model/aggregates"
	"pikaUtils/design-combination/model/vo"
	"pikaUtils/design-combination/service/engine/impl"
	"pikaUtils/design-combination/service/logic"
)

/**
  @Author: pika
  @Date: 2022/4/1 3:29 下午
  @Description:
*/

var (
	treeNode_01      vo.TreeNode
	treeNode_11      vo.TreeNode
	treeNode_12      vo.TreeNode
	treeNode_111     vo.TreeNode
	treeNode_112     vo.TreeNode
	treeNode_121     vo.TreeNode
	treeNode_122     vo.TreeNode
	treeNodeLink_11  vo.TreeNodeLink
	treeNodeLink_12  vo.TreeNodeLink
	treeNodeLink_111 vo.TreeNodeLink
	treeNodeLink_112 vo.TreeNodeLink
	treeNodeLink_121 vo.TreeNodeLink
	treeNodeLink_122 vo.TreeNodeLink
	treeRoot         vo.TreeRoot
	treeNodeMap      map[int64]vo.TreeNode
	treeRich         aggregates.TreeRich
)

func init() {
	// 节点：1 Gender
	treeNode_01.SetTreeId(10001)
	treeNode_01.SetTreeNodeId(1)
	treeNode_01.SetNodeType(vo.NodeTypeLeaf)
	treeNode_01.SetNodeValue("")
	treeNode_01.SetRuleKey("userGender")
	treeNode_01.SetRuleDesc("用户性别[男/女]")

	// 节点：11 Age
	treeNode_11.SetTreeId(10001)
	treeNode_11.SetTreeNodeId(11)
	treeNode_11.SetNodeType(vo.NodeTypeLeaf)
	treeNode_11.SetNodeValue("")
	treeNode_11.SetRuleKey("userAge")
	treeNode_11.SetRuleDesc("用户年龄")

	// 节点：12 Age
	treeNode_12.SetTreeId(10001)
	treeNode_12.SetTreeNodeId(12)
	treeNode_12.SetNodeType(vo.NodeTypeLeaf)
	treeNode_12.SetNodeValue("")
	treeNode_12.SetRuleKey("userAge")
	treeNode_12.SetRuleDesc("用户年龄")

	// 节点：111 果实A
	treeNode_111.SetTreeId(10001)
	treeNode_111.SetTreeNodeId(111)
	treeNode_111.SetNodeType(vo.NodeTypeFruit)
	treeNode_111.SetNodeValue("果实A")

	// 节点：112 果实B
	treeNode_112.SetTreeId(10001)
	treeNode_112.SetTreeNodeId(112)
	treeNode_112.SetNodeType(vo.NodeTypeFruit)
	treeNode_112.SetNodeValue("果实B")

	// 节点：121 果实C
	treeNode_121.SetTreeId(10001)
	treeNode_121.SetTreeNodeId(121)
	treeNode_121.SetNodeType(vo.NodeTypeFruit)
	treeNode_121.SetNodeValue("果实C")

	// 节点：122 果实D
	treeNode_122.SetTreeId(10001)
	treeNode_122.SetTreeNodeId(122)
	treeNode_122.SetNodeType(vo.NodeTypeFruit)
	treeNode_122.SetNodeValue("果实D")

	// 链接：1->11
	treeNodeLink_11.SetNodeIdFrom(treeNode_01.TreeNodeId())
	treeNodeLink_11.SetNodeIdTo(treeNode_11.TreeNodeId())
	treeNodeLink_11.SetRuleLimitType(logic.EQ)
	treeNodeLink_11.SetRuleLimitValue("man")

	// 链接：1->12
	treeNodeLink_12.SetNodeIdFrom(treeNode_01.TreeNodeId())
	treeNodeLink_12.SetNodeIdTo(treeNode_12.TreeNodeId())
	treeNodeLink_12.SetRuleLimitType(logic.EQ)
	treeNodeLink_12.SetRuleLimitValue("woman")
	treeNode_01.SetTreeNodeLinkList([]vo.TreeNodeLink{treeNodeLink_11, treeNodeLink_12})

	// 链接：11->111
	treeNodeLink_111.SetNodeIdFrom(treeNode_11.TreeNodeId())
	treeNodeLink_111.SetNodeIdTo(treeNode_111.TreeNodeId())
	treeNodeLink_111.SetRuleLimitType(logic.LT)
	treeNodeLink_111.SetRuleLimitValue("25")

	// 链接：11->112
	treeNodeLink_112.SetNodeIdFrom(treeNode_11.TreeNodeId())
	treeNodeLink_112.SetNodeIdTo(treeNode_112.TreeNodeId())
	treeNodeLink_112.SetRuleLimitType(logic.GTE)
	treeNodeLink_112.SetRuleLimitValue("25")
	treeNode_11.SetTreeNodeLinkList([]vo.TreeNodeLink{treeNodeLink_111, treeNodeLink_112})

	// 链接：12->121
	treeNodeLink_121.SetNodeIdFrom(treeNode_12.TreeNodeId())
	treeNodeLink_121.SetNodeIdTo(treeNode_121.TreeNodeId())
	treeNodeLink_121.SetRuleLimitType(logic.LT)
	treeNodeLink_121.SetRuleLimitValue("25")

	// 链接：12->122
	treeNodeLink_122.SetNodeIdFrom(treeNode_12.TreeNodeId())
	treeNodeLink_122.SetNodeIdTo(treeNode_122.TreeNodeId())
	treeNodeLink_122.SetRuleLimitType(logic.GTE)
	treeNodeLink_122.SetRuleLimitValue("25")
	treeNode_12.SetTreeNodeLinkList([]vo.TreeNodeLink{treeNodeLink_121, treeNodeLink_122})

	// 树根
	treeRoot.SetTreeId(treeNode_01.TreeId())
	treeRoot.SetTreeRootNodeId(treeNode_01.TreeNodeId())
	treeRoot.SetTreeName("规则决策树")

	treeNodeMap = make(map[int64]vo.TreeNode, 7)
	treeNodeMap[treeNode_01.TreeNodeId()] = treeNode_01
	treeNodeMap[treeNode_11.TreeNodeId()] = treeNode_11
	treeNodeMap[treeNode_12.TreeNodeId()] = treeNode_12
	treeNodeMap[treeNode_111.TreeNodeId()] = treeNode_111
	treeNodeMap[treeNode_112.TreeNodeId()] = treeNode_112
	treeNodeMap[treeNode_121.TreeNodeId()] = treeNode_121
	treeNodeMap[treeNode_122.TreeNodeId()] = treeNode_122

	treeRich = *aggregates.NewTreeRich(treeRoot, treeNodeMap)
}

func main() {

	treeRichJson, _ := json.Marshal(&treeRich)
	fmt.Println("决策树组合结构信息: \n", string(treeRichJson))
	var treeEngineHandle impl.TreeEngineHandle

	decisionMatter := make(map[string]string, 2)
	decisionMatter["gender"] = "man"
	decisionMatter["age"] = "29"

	res := treeEngineHandle.Process(10001, "用户1", treeRich, decisionMatter)
	fmt.Println(res)
	resJson, _ := json.Marshal(&res)
	fmt.Println("测试结果:\n", string(resJson))

}

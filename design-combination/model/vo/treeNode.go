package vo

/**
  @Author: pika
  @Date: 2022/3/30 8:38 下午
  @Description: 规则树节点信息
*/

const (
	// NodeTypeLeaf 叶子节点类型
	NodeTypeLeaf = 1
	// NodeTypeFruit 果实节点类型
	NodeTypeFruit = 2
)

type TreeNode struct {
	treeId           int64          // 规则树ID
	treeNodeId       int64          // 规则树节点ID
	nodeType         int            // 节点类型；1子叶、2果实
	nodeValue        string         // 节点值[nodeType=2]；果实值
	ruleKey          string         // 规则Key
	ruleDesc         string         // 规则描述
	treeNodeLinkList []TreeNodeLink // 节点链路
}

func (tn *TreeNode) TreeId() int64 {
	return tn.treeId
}

func (tn *TreeNode) SetTreeId(treeId int64) {
	tn.treeId = treeId
}

func (tn *TreeNode) TreeNodeId() int64 {
	return tn.treeNodeId
}

func (tn *TreeNode) SetTreeNodeId(treeNodeId int64) {
	tn.treeNodeId = treeNodeId
}

func (tn *TreeNode) NodeType() int {
	return tn.nodeType
}

func (tn *TreeNode) SetNodeType(nodeType int) {
	tn.nodeType = nodeType
}

func (tn *TreeNode) NodeValue() string {
	return tn.nodeValue
}

func (tn *TreeNode) SetNodeValue(nodeValue string) {
	tn.nodeValue = nodeValue
}

func (tn *TreeNode) RuleKey() string {
	return tn.ruleKey
}

func (tn *TreeNode) SetRuleKey(ruleKey string) {
	tn.ruleKey = ruleKey
}

func (tn *TreeNode) RuleDesc() string {
	return tn.ruleDesc
}

func (tn *TreeNode) SetRuleDesc(ruleDesc string) {
	tn.ruleDesc = ruleDesc
}

func (tn *TreeNode) TreeNodeLinkList() []TreeNodeLink {
	return tn.treeNodeLinkList
}

func (tn *TreeNode) SetTreeNodeLinkList(treeNodeLinkList []TreeNodeLink) {
	tn.treeNodeLinkList = treeNodeLinkList
}

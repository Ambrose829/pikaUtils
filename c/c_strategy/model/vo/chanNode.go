package vo

/**
  @Author: pika
  @Date: 2022/4/2 10:58 上午
  @Description: 链表节点
*/

type ChanNode struct {
	nodeId           int64          //链节点id
	ruleKey          string         // 规则Key city、biz
	ruleDesc         string         // 规则描述
	treeNodeLinkList []ChanNodeLink // 节点链路

}

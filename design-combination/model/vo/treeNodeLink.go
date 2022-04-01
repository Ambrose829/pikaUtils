package vo

/**
  @Author: pika
  @Date: 2022/3/30 8:38 下午
  @Description: 规则树线信息
*/

type TreeNodeLink struct {
	nodeIdFrom     int64  //节点From
	nodeIdTo       int64  //节点To
	ruleLimitType  int    //限定类型；1:=;2:>;3:<;4:>=;5<=;6:enum[枚举范围]
	ruleLimitValue string //限定值
}

func (tnl *TreeNodeLink) NodeIdFrom() int64 {
	return tnl.nodeIdFrom
}

func (tnl *TreeNodeLink) SetNodeIdFrom(nodeIdFrom int64) {
	tnl.nodeIdFrom = nodeIdFrom
}

func (tnl *TreeNodeLink) NodeIdTo() int64 {
	return tnl.nodeIdTo
}

func (tnl *TreeNodeLink) SetNodeIdTo(nodeIdTo int64) {
	tnl.nodeIdTo = nodeIdTo
}

func (tnl *TreeNodeLink) RuleLimitType() int {
	return tnl.ruleLimitType
}

func (tnl *TreeNodeLink) SetRuleLimitType(ruleLimitType int) {
	tnl.ruleLimitType = ruleLimitType
}

func (tnl *TreeNodeLink) RuleLimitValue() string {
	return tnl.ruleLimitValue
}

func (tnl *TreeNodeLink) SetRuleLimitValue(ruleLimitValue string) {
	tnl.ruleLimitValue = ruleLimitValue
}

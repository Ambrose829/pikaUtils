package vo

/**
  @Author: pika
  @Date: 2022/4/2 11:08 上午
  @Description: 链表节点连接线
*/

type ChanNodeLink struct {
	nodeIdFrom     int64  //节点From
	nodeIdTo       int64  //节点To
	ruleLimitType  int    //限定类型；1:=;2:>;3:<;4:>=;5:<=;6:!=;7:in;8:notin;9:enum[枚举范围] 也可推出限定值类型
	ruleLimitValue string //限定值
}

func (cnl *ChanNodeLink) NodeIdFrom() int64 {
	return cnl.nodeIdFrom
}

func (cnl *ChanNodeLink) SetNodeIdFrom(nodeIdFrom int64) {
	cnl.nodeIdFrom = nodeIdFrom
}

func (cnl *ChanNodeLink) NodeIdTo() int64 {
	return cnl.nodeIdTo
}

func (cnl *ChanNodeLink) SetNodeIdTo(nodeIdTo int64) {
	cnl.nodeIdTo = nodeIdTo
}

func (cnl *ChanNodeLink) RuleLimitType() int {
	return cnl.ruleLimitType
}

func (cnl *ChanNodeLink) SetRuleLimitType(ruleLimitType int) {
	cnl.ruleLimitType = ruleLimitType
}

func (cnl *ChanNodeLink) RuleLimitValue() string {
	return cnl.ruleLimitValue
}

func (cnl *ChanNodeLink) SetRuleLimitValue(ruleLimitValue string) {
	cnl.ruleLimitValue = ruleLimitValue
}

package vo

/**
  @Author: pika
  @Date: 2022/3/30 8:38 下午
  @Description: 决策结果
*/

type EngineResult struct {
	isSuccess bool   //执行结果
	userId    string //用户ID
	treeId    int64  //规则树ID
	nodeId    int64  //果实节点ID
	nodeValue string //果实节点值

}

func NewEngineResult(userId string, treeId int64, nodeId int64, nodeValue string) *EngineResult {
	return &EngineResult{userId: userId, treeId: treeId, nodeId: nodeId, nodeValue: nodeValue}
}

func (er *EngineResult) IsSuccess() bool {
	return er.isSuccess
}

func (er *EngineResult) SetIsSuccess(isSuccess bool) {
	er.isSuccess = isSuccess
}

func (er *EngineResult) UserId() string {
	return er.userId
}

func (er *EngineResult) SetUserId(userId string) {
	er.userId = userId
}

func (er *EngineResult) TreeId() int64 {
	return er.treeId
}

func (er *EngineResult) SetTreeId(treeId int64) {
	er.treeId = treeId
}

func (er *EngineResult) NodeId() int64 {
	return er.nodeId
}

func (er *EngineResult) SetNodeId(nodeId int64) {
	er.nodeId = nodeId
}

func (er *EngineResult) NodeValue() string {
	return er.nodeValue
}

func (er *EngineResult) SetNodeValue(nodeValue string) {
	er.nodeValue = nodeValue
}

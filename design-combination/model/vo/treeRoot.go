package vo

/**
  @Author: pika
  @Date: 2022/3/30 8:39 下午
  @Description: 树根信息
*/

type TreeRoot struct {
	treeId         int64  //规则树ID
	treeRootNodeId int64  //规则树根ID
	treeName       string //规则树名称

}

func (tr *TreeRoot) TreeId() int64 {
	return tr.treeId
}

func (tr *TreeRoot) SetTreeId(treeId int64) {
	tr.treeId = treeId
}

func (tr *TreeRoot) TreeRootNodeId() int64 {
	return tr.treeRootNodeId
}

func (tr *TreeRoot) SetTreeRootNodeId(treeRootNodeId int64) {
	tr.treeRootNodeId = treeRootNodeId
}

func (tr *TreeRoot) TreeName() string {
	return tr.treeName
}

func (tr *TreeRoot) SetTreeName(treeName string) {
	tr.treeName = treeName
}

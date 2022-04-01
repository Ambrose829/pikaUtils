package aggregates

import "pikaUtils/design-combination/model/vo"

/**
  @Author: pika
  @Date: 2022/3/31 3:22 下午
  @Description: 规则树聚合
*/

type TreeRich struct {
	treeRoot    vo.TreeRoot
	treeNodeMap map[int64]vo.TreeNode
}

func NewTreeRich(treeRoot vo.TreeRoot, treeNodeMap map[int64]vo.TreeNode) *TreeRich {
	return &TreeRich{treeRoot: treeRoot, treeNodeMap: treeNodeMap}
}

func (trh *TreeRich) TreeRoot() vo.TreeRoot {
	return trh.treeRoot
}

func (trh *TreeRich) SetTreeRoot(treeRoot vo.TreeRoot) {
	trh.treeRoot = treeRoot
}

func (trh *TreeRich) TreeNodeMap() map[int64]vo.TreeNode {
	return trh.treeNodeMap
}

func (trh *TreeRich) SetTreeNodeMap(treeNodeMap map[int64]vo.TreeNode) {
	trh.treeNodeMap = treeNodeMap
}

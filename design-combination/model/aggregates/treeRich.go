package aggregates

import "pikaUtils/design-combination/model/vo"

/**
  @Author: pika
  @Date: 2022/3/31 3:22 下午
  @Description: 规则树聚合
*/

type TreeRich struct {
	treeRoot vo.TreeNode
	treeNodeMap map[int64]vo.TreeNode
}

func (trh *TreeRich) TreeRoot() vo.TreeNode {
	return trh.treeRoot
}

func (trh *TreeRich) SetTreeRoot(treeRoot vo.TreeNode) {
	trh.treeRoot = treeRoot
}

func (trh *TreeRich) TreeNodeMap() map[int64]vo.TreeNode {
	return trh.treeNodeMap
}

func (trh *TreeRich) SetTreeNodeMap(treeNodeMap map[int64]vo.TreeNode) {
	trh.treeNodeMap = treeNodeMap
}



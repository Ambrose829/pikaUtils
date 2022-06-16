package aggregates

import "pikaUtils/c/c_strategy/model/vo"

/**
  @Author: pika
  @Date: 2022/4/2 1:38 下午
  @Description: 链信息
*/

type ChanRich struct {
	chanNodeMap map[int64]vo.ChanNode
}

func (cr *ChanRich) ChanNodeMap() map[int64]vo.ChanNode {
	return cr.chanNodeMap
}

func (cr *ChanRich) SetChanNodeMap(chanNodeMap map[int64]vo.ChanNode) {
	cr.chanNodeMap = chanNodeMap
}

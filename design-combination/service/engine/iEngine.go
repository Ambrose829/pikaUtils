package engine

import (
	"pikaUtils/design-combination/model/aggregates"
	"pikaUtils/design-combination/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/1 1:09 下午
  @Description: 引擎
*/

type IEngine interface {
	process(treeId int64, userId string, rich aggregates.TreeRich, decisionMatter map[string]string) vo.EngineResult
}

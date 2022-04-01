package logic

import "pikaUtils/design-combination/model/vo"

/**
  @Author: pika
  @Date: 2022/3/30 5:43 下午
  @Description: 策略
*/


type LogicFilter interface {

	//
    // Filter
    // @author: pika
    // @Description: 逻辑决策器
    // @param matterValue 决策值
    // @param treeNodeLineInfoList 决策节点
    // @return int64 下一个节点id
    //
	Filter(matterValue string, treeNodeLineInfoList []vo.TreeNodeLink) int64

	//
    // MatterValue
    // @author: pika
    // @Description: 获取决策值
    // @param treeId 决策树id
    // @param userId 用户id
    // @param decisionMatter 决策物料
    // @return string 决策值
    //
	MatterValue(treeId int64, userId string, decisionMatter map[string]string) string
}



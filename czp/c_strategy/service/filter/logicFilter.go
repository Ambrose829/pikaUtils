package filter

import "pikaUtils/czp/c_strategy/model/vo"

/**
  @Author: pika
  @Date: 2022/4/1 8:13 下午
  @Description: 过滤器interface
*/

// LogicFilter 策略过滤器接口
type LogicFilter interface {
	// Filter 过滤
	Filter(matterValue string, chanNodeLineInfoList []vo.ChanNodeLink) bool

	MatterValue(strategy vo.Strategy) string
}

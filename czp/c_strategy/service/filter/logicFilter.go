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
	Filter(strategy vo.Strategy) bool

	// MatterValue 获取需要验证的值
	MatterValue(strategy vo.Strategy) string
}

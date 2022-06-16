package filter

import (
	"pikaUtils/c/s/strategy"
)

/**
  @Author: pika
  @Date: 2022/4/15 6:57 下午
  @Description: 过滤器interface, 新增过滤器组件需要实现此接口
*/

// LogicFilter 策略过滤器接口
type LogicFilter interface {
	// Filter 过滤
	Filter(strategy strategy.Strategy) bool

	// MatterValue 获取需要验证的值
	MatterValue(strategy strategy.Strategy) string
}

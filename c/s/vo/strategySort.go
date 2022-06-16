package vo

import (
	"pikaUtils/c/s/strategy"
	"sort"
)

/**
  @Author: pika
  @Date: 2022/4/21 11:09 上午
  @Description: 策略排序
*/

func StrategySort(list []strategy.Strategy) {
	var sl StgSlice
	sl = list
	sort.Sort(sl)
}

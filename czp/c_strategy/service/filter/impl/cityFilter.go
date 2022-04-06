package impl

import (
	"pikaUtils/czp/c_strategy/model/vo"
)

/**
  @Author: pika
  @Date: 2022/4/1 8:48 下午
  @Description: 城市过滤器
*/

type CityFilter struct {
}

func (cf CityFilter) Execute(matterValue string, chanNodeLineInfoList []vo.ChanNodeLink) bool {
	panic("implement me")
}

func (cf CityFilter) MatterValue(strategy vo.Strategy) string {
	panic("implement me")
}

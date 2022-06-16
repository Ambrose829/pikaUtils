package strategy

import (
	"pikaUtils/czp/s/strategy"
	"strings"
)

/**
  @Author: pika
  @Date: 2022/4/20 4:28 下午
  @Description: 工具类
*/

const (
	maxUnit    = 10
	cityWeight = 1
	bizWeight  = 10
)

func calculatePriority(strategy strategy.Strategy) int64 {

	// 业务线和城市占的单位权重的数量, 最高为9
	var bizUnit, cityUnit int64
	// 计算这条策略的业务线维度占有的单位权重数量
	if strategy.BizTypes != "" {
		bizs := strings.Split(strategy.BizTypes, ",")
		if len(bizs) > 0 {
			bizUnit = int64(maxUnit - len(bizs))
			if bizUnit <= 0 {
				bizUnit = 1
			}
		}

	}
	if strategy.CityList != "" {
		cities := strings.Split(strategy.CityList, ",")
		if len(cities) > 0 {
			cityUnit = int64(maxUnit - len(cities))
			if cityUnit <= 0 {
				cityUnit = 1
			}
		}
	}

	return bizUnit*bizWeight + cityUnit*cityWeight

}

type StgSlice []strategy.Strategy

func (ss StgSlice) Len() int {
	return len(ss)
}

func (ss StgSlice) Less1(i, j int) bool {
	// i的业务线为全部时
	if ss[i].BizTypes == "" {
		// 如果j的业务线也为全部
		if ss[j].BizTypes == "" {
			// 看i的城市
			// 如果i的城市为全部
			if ss[i].CityList == "" {

				//如果j的城市也为全部
				if ss[j].CityList == "" {
					return true
					// 如果j的城市不为全部,优先级出来
				} else {
					return false
				}
				// 如果i的城市不为全部
			} else {
				//判断逻辑

			}
			// 如果j的业务线不是全部，优先级出来
		} else {
			return false
		}
		// 如果i的业务线不为全部
	} else {

	}

	//iBizs := strings.Split(ss[i].BizTypes, ",")
	//jBizs := strings.Split(ss[j].BizTypes, ",")
	//iCities := strings.Split(ss[i].CityList, ",")
	//jCities := strings.Split(ss[].CityList, ",")

	return true

}

func (ss StgSlice) Less(i, j int) bool {

	// 业务线和城市占的单位权重的数量, 最高为9
	var bizUnit1, cityUnit1, bizUnit2, cityUnit2 int64
	// ss[i]
	// 计算这条策略的业务线维度占有的单位权重数量
	if ss[i].BizTypes != "" {
		bizs := strings.Split(ss[i].BizTypes, ",")
		if len(bizs) > 0 {
			bizUnit1 = int64(maxUnit - len(bizs))
			if bizUnit1 <= 0 {
				bizUnit1 = 1
			}
		}
	}
	// 计算这条策略的城市维度占有的单位权重数量
	if ss[i].CityList != "" {
		cities := strings.Split(ss[i].CityList, ",")
		if len(cities) > 0 {
			cityUnit1 = int64(maxUnit - len(cities))
			if cityUnit1 <= 0 {
				cityUnit1 = 1
			}
		}
	}

	// ss[j]
	// 计算这条策略的业务线维度占有的单位权重数量
	if ss[j].BizTypes != "" {
		bizs := strings.Split(ss[j].BizTypes, ",")
		if len(bizs) > 0 {
			bizUnit2 = int64(maxUnit - len(bizs))
			if bizUnit2 <= 0 {
				bizUnit2 = 1
			}
		}
	}
	// 计算这条策略的城市维度占有的单位权重数量
	if ss[j].CityList != "" {
		cities := strings.Split(ss[j].CityList, ",")
		if len(cities) > 0 {
			cityUnit2 = int64(maxUnit - len(cities))
			if cityUnit2 <= 0 {
				cityUnit2 = 1
			}
		}
	}

	iPriority := bizUnit1*bizWeight + cityUnit1*cityWeight
	jPriority := bizUnit2*bizWeight + cityUnit2*cityWeight
	return iPriority > jPriority
}

func (ss StgSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

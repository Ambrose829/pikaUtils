package vo

import (
	"pikaUtils/czp/s/strategy"
	"testing"
)

/**
  @Author: pika
  @Date: 2022/4/21 11:13 上午
  @Description:
*/

func BenchmarkStrategySort(b *testing.B) {
	// 测试用例
	s1 := strategy.Strategy{
		CityList: "北京,上海,深圳",
		BizTypes: "专车,快车,豪华车",
	}

	s2 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "专车,快车,豪华车",
	}

	s3 := strategy.Strategy{
		CityList: "北京",
		BizTypes: "专车,快车,豪华车",
	}

	s4 := strategy.Strategy{
		CityList: "北京,上海,深圳",
		BizTypes: "专车,快车",
	}

	s5 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "快车,专车",
	}

	s6 := strategy.Strategy{
		CityList: "北京",
		BizTypes: "快车,专车",
	}

	s7 := strategy.Strategy{
		CityList: "北京,上海,深圳",
		BizTypes: "快车",
	}

	s8 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "快车",
	}

	s9 := strategy.Strategy{
		CityList: "北京",
		BizTypes: "快车",
	}

	s10 := strategy.Strategy{
		CityList: "北京,上海,深圳",
		BizTypes: "",
	}

	s11 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "",
	}

	s12 := strategy.Strategy{
		CityList: "北京",
		BizTypes: "",
	}

	s13 := strategy.Strategy{
		CityList: "",
		BizTypes: "专车,快车,豪华车",
	}

	s14 := strategy.Strategy{
		CityList: "",
		BizTypes: "专车,快车",
	}

	s15 := strategy.Strategy{
		CityList: "",
		BizTypes: "快车",
	}

	list := []strategy.Strategy{s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrategySort(list)
	}
}

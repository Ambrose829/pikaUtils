package vo

import (
	"pikaUtils/czp/s/strategy"
	"reflect"
	"testing"
)

/**
  @Author: pika
  @Date: 2022/4/21 11:14 上午
  @Description: 测试策略排序
*/

func TestStrategySort(t *testing.T) {
	type args struct {
		list []strategy.Strategy
	}
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

	ag1 := args{[]strategy.Strategy{s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15}}

	tests := []struct {
		name             string
		args             args
		wantPassStrategy []strategy.Strategy
	}{
		{
			name:             "策略排序单元测试",
			args:             ag1,
			wantPassStrategy: []strategy.Strategy{s9, s8, s7, s15, s6, s5, s4, s14, s3, s2, s1, s13, s12, s11, s10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if StrategySort(tt.args.list); !reflect.DeepEqual(tt.args.list, tt.wantPassStrategy) {
				t.Errorf("EngineDecisionMaker() = %v, want %v", tt.args.list, tt.wantPassStrategy)
			}
		})
	}
}

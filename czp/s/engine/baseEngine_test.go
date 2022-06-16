package engine

import (
	"pikaUtils/czp/s/filter"
	"pikaUtils/czp/s/filter/impl"
	"pikaUtils/czp/s/filter/impl/criteria"
	"pikaUtils/czp/s/strategy"
	"reflect"
	"testing"
)

/**
  @Author: pika
  @Date: 2022/4/19 2:36 下午
  @Description:
*/

func TestBaseEngine_EngineDecisionMaker(t *testing.T) {
	type fields struct {
		IEngine IEngine
	}
	type args struct {
		filterChan     impl.FilterChan
		decisionMatter []strategy.Strategy
	}

	// 初始化测试数据
	var cf1 *impl.CityFilter
	var bf1 *impl.BizFilter

	// 城市规则是北京
	cf1 = impl.NewCityFilter(*filter.NewBaseFilter(filter.INC, "北京"))
	bf1 = impl.NewBizFilter(*filter.NewBaseFilter(filter.INC, "快车"))

	// 需要将全部的情况考虑
	var cf2 *impl.CityFilter
	var bf2 *impl.BizFilter
	cf2 = impl.NewCityFilter(*filter.NewBaseFilter(filter.EQ, ""))
	bf2 = impl.NewBizFilter(*filter.NewBaseFilter(filter.EQ, ""))

	// 或规则关联
	var cf *criteria.OrCriteria
	var bf *criteria.OrCriteria
	cf = criteria.NewOrCriteria(*criteria.NewCriteria(cf1, cf2))
	bf = criteria.NewOrCriteria(*criteria.NewCriteria(bf1, bf2))

	var fc impl.FilterChan
	fc.AddFilter(cf)
	fc.AddFilter(bf)
	fc.SetTarget("城市业、务线过滤器链")

	s1 := strategy.Strategy{
		CityList: "北京,广州",
		BizTypes: "快车,D1",
	}

	s2 := strategy.Strategy{
		CityList: "北京,深圳",
		BizTypes: "专车,快车",
	}

	s3 := strategy.Strategy{
		CityList: "上海,徐州",
		BizTypes: "快车,优享",
	}

	s4 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "专车,豪华车",
	}

	s5 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "快车,专车",
	}

	s6 := strategy.Strategy{
		CityList: "",
		BizTypes: "快车,专车",
	}

	s7 := strategy.Strategy{
		CityList: "北京,上海",
		BizTypes: "",
	}

	s8 := strategy.Strategy{
		CityList: "",
		BizTypes: "",
	}

	s9 := strategy.Strategy{
		CityList: "北京",
		BizTypes: "快车",
	}

	ag1 := args{
		filterChan:     fc,
		decisionMatter: []strategy.Strategy{s1, s2, s3, s4, s5, s6, s7, s8, s9},
	}

	var tests = []struct {
		name             string
		fields           fields
		args             args
		wantPassStrategy []strategy.Strategy
	}{
		{
			name:             "测试1",
			fields:           fields{IEngine: BaseEngine{}},
			args:             ag1,
			wantPassStrategy: []strategy.Strategy{s1, s2, s5, s6, s7, s8, s9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			be := BaseEngine{
				IEngine: tt.fields.IEngine,
			}
			if gotPassStrategy := be.EngineDecisionMaker(tt.args.filterChan, tt.args.decisionMatter); !reflect.DeepEqual(gotPassStrategy, tt.wantPassStrategy) {
				t.Errorf("EngineDecisionMaker() = %v, want %v", gotPassStrategy, tt.wantPassStrategy)
			}
		})
	}
}

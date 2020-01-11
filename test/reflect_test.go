package test

import (
	"reflect"
	"testing"
)

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
func walk(x interface{}, fn func(input string) )  {
	val := getValue(x)

	//通过递归函数脱去数据结构的外壳
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	//根据val的类型处理
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i< val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i:= 0; i<val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}
//reflect 常用来做参数校验,他比较消耗性能,如非必要不要使用
func TestWalk(t *testing.T)  {
	//测试表
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{ "Chris"},
			[]string{"Chris"},
		}, {
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases{
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got,test.ExpectedCalls){
				t.Errorf("got %s want %s", got[0], test.ExpectedCalls)
			}
		})
	}


}

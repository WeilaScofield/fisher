/*
sort.go封装了排序条件以及slice类型的原地正反排序
*/
package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const (
	SortByKey = "sort_by"
)

var (
	genLessFuncError = fmt.Errorf("data is not a slice")
)

type SortCondition struct {
	Condition string
	Reverse   bool
}

// return "desc" if Reverse==true else "asc"
func (sc SortCondition) ReverseString() string {
	if sc.Reverse {
		return "desc"
	}
	return "asc"
}

// 当condition不为空时, 认为需要进行排序
func (sc SortCondition) NeedSort() bool {
	return sc.Condition != ""
}

/*
条件原地排序:
对`[]struct`类型的slice按struct中指定字段进行排序
支持int, int64, string, float64, uint, uint64类型

usage:
s := []MyStruct{
	{Foo: "a", Bar: "a"},
	{Foo: "b", Bar: "b"},
}
sc := SortCondition{
	Condition: "Foo" //struct 字段名
	Reverse: true //是否倒序
}
ConditionalSort(s, sc)
*/
func ConditionalSort(slice interface{}, condition SortCondition) {
	lessfunc, err := genLessFunc(slice, condition)
	if err != nil {
		fmt.Println(err)
		return
	}
	if condition.Reverse {
		lessfunc = reverseLessFunc(lessfunc)
	}
	sort.Slice(slice, lessfunc)
}

// 反转lessfunc
func reverseLessFunc(f func(i, j int) bool) func(i, j int) bool {
	return func(i, j int) bool {
		return !f(i, j)
	}
}

// 生成less function
func genLessFunc(slice interface{}, sc SortCondition) (f func(i, j int) bool, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		f = func(i, j int) bool {
			iv := reflect.ValueOf(s.Index(i).Interface()).FieldByName(sc.Condition)
			jv := reflect.ValueOf(s.Index(j).Interface()).FieldByName(sc.Condition)
			switch iv.Kind() {
			case reflect.Int, reflect.Int64:
				return iv.Int() < jv.Int()
			case reflect.String:
				return iv.String() < jv.String()
			case reflect.Float64:
				return iv.Float() < jv.Float()
			case reflect.Uint, reflect.Uint64:
				return iv.Uint() < jv.Uint()
			}
			return false
		}
		return f, nil
	}
	return f, genLessFuncError
}

/*
解析排序条件, 默认顺序排列asc
采用k8s dashboard的url参数形式
sortBy=a,createTime 表示按照createTime字段顺序排序
sortBy=d,name 表示按照name倒叙排序
*/
func ParseSortCondition(s string) SortCondition {
	res := SortCondition{}
	cond := strings.Split(s, ",")
	if len(cond) < 2 {
		return res
	}
	if cond[0] == "d" {
		res.Reverse = true
	}
	res.Condition = cond[1]
	return res
}

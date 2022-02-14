package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func GroupBy(arr interface{}, groupFunc interface{}) interface{} {
	groupMap := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(groupFunc).Out(0), reflect.TypeOf(arr)))
	for i := 0; i < reflect.ValueOf(arr).Len(); i++ {
		groupPivot := reflect.ValueOf(groupFunc).Call([]reflect.Value{reflect.ValueOf(arr).Index(i)})[0]
		if !groupMap.MapIndex(groupPivot).IsValid() {
			groupMap.SetMapIndex(groupPivot, reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(arr).Elem()), 0, 0))
		}
		groupMap.SetMapIndex(groupPivot, reflect.Append(groupMap.MapIndex(groupPivot), reflect.ValueOf(arr).Index(i)))
	}
	return groupMap.Interface()
}

func main() {
	fmt.Println(GroupBy([]int{1, 22, 3, 44, 555}, func(value int) string {
		return strconv.Itoa(len(strconv.Itoa(value)))
	}))
}

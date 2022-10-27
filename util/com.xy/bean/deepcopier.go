package bean

/**
 * @description:
 * @author:xy
 * @date:2022/9/8 9:45
 * @Version: 1.0
 */

import (
	"github.com/ulule/deepcopier"
	"reflect"
)

//Of struct deep copy
func Of[T any](src any) (t T) {
	err := deepcopier.Copy(src).To(&t)
	if err != nil {
		panic(err)
	}
	return
}

//Ofs list deep copy
//src must be Slice or Array
//T = destType
//T1 = sourceType
//return value type
func Ofs[T any, T1 any](src any) []T {
	var ts []T
	value := reflect.ValueOf(src)
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		for i := 0; i < value.Len(); i++ {
			t1 := value.Index(i).Interface().(T1)
			of := Of[T](t1)
			ts = append(ts, of)
		}
	}
	return ts
}

//OfsRe list deep copy
//src must be Slice or Array
//T = destType
//T1 = sourceType
// warn 会导致0 nil 变量json消失
//return Reference types
func OfsRe[T any, T1 any](src any) []*T {
	var ts []*T
	value := reflect.ValueOf(src)
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		for i := 0; i < value.Len(); i++ {
			t1 := value.Index(i).Interface().(T1)
			of := Of[T](t1)
			ts = append(ts, &of)
		}
	}
	return ts
}

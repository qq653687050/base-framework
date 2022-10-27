package bean

import (
	"reflect"
)

/**
 * @description: 判空
 * @author:xy
 * @date:2022/10/25 15:20
 * @Version: 1.0
 */

// CheckIsNil .
//指针类型判空
func CheckIsNil(arg interface{}) (b bool) {
	if reflect.ValueOf(arg).IsNil() { //利用反射直接判空，指针用isNil
		b = true
	}
	return
}

// CheckNonNil .
//指针类型非空
func CheckNonNil(arg interface{}) (b bool) {
	if !reflect.ValueOf(arg).IsNil() { //利用反射直接判空，指针用isNil

		b = true
	}
	return
}

// CheckIsZero .
//基本类型判空
func CheckIsZero(arg interface{}) (b bool) {
	if reflect.ValueOf(arg).IsZero() { //利用反射直接判空，基础数据类型用isZero
		b = true
	}
	return
}

// CheckNonZero .
//基本类型非空
func CheckNonZero(arg interface{}) (b bool) {
	if !reflect.ValueOf(arg).IsZero() { //利用反射直接判空，基础数据类型用isZero
		b = true
	}
	return
}

package comm

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var (
	ErrInvalidValue = errors.New("Invalid value")
)

// ParseString 从任意类型转换成string值
func ParseString(val interface{}) string {
	if val == nil {
		return ""
	}
	// 判断val类型
	switch v := val.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case error:
		return v.Error()
	}
	// 判断val种类
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.String {
		return v.String()
	}
	// 默认打印一下
	return fmt.Sprint(val)
}

// ParseBool 从任意类型转换成bool值
func ParseBool(val interface{}) (bool, error) {
	if val == nil {
		return false, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case bool:
		return v, nil
	}
	// 判断val种类
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Bool {
		return v.Bool(), nil
	}
	// 默认先转成字符串，再解析
	str := ParseString(val)
	return strconv.ParseBool(str)
}

// ParseInt 从任意类型转换成int值
func ParseInt(val interface{}) (int, error) {
	if val == nil {
		return 0, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case int:
		return v, nil
	}
	// 判断val种类
	var i int64
	var err error
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i = v.Int()
	default:
		str := ParseString(val)
		i, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			return 0, err
		}
	}
	// 判断是否溢出
	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0, strconv.ErrRange
	}
	return int(i), nil
}

// ParseInt64 从任意类型转换成int64值
func ParseInt64(val interface{}) (int64, error) {
	if val == nil {
		return 0, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case int64:
		return v, nil
	}
	// 判断val种类
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	default:
		str := ParseString(val)
		return strconv.ParseInt(str, 10, 64)
	}
}

// ParseUint 从任意类型转换成uint值
func ParseUint(val interface{}) (uint, error) {
	if val == nil {
		return 0, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case uint:
		return v, nil
	}
	// 判断val种类
	var i uint64
	var err error
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i = v.Uint()
	default:
		str := ParseString(val)
		i, err = strconv.ParseUint(str, 10, 32)
		if err != nil {
			return 0, err
		}
	}
	// 判断是否溢出
	if i > math.MaxUint32 {
		return 0, strconv.ErrRange
	}
	return uint(i), nil
}

// ParseUint64 从任意类型转换成uint64值
func ParseUint64(val interface{}) (uint64, error) {
	if val == nil {
		return 0, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case uint64:
		return v, nil
	}
	// 判断val种类
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	default:
		str := ParseString(val)
		return strconv.ParseUint(str, 10, 64)
	}
}

// ParseFloat 从任意类型转换成float64值
func ParseFloat(val interface{}) (float64, error) {
	if val == nil {
		return 0, ErrInvalidValue
	}
	// 判断val类型
	switch v := val.(type) {
	case float64:
		return v, nil
	}
	// 判断val种类
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	default:
		str := ParseString(val)
		return strconv.ParseFloat(str, 64)
	}
}

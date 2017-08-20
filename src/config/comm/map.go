package comm

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Map struct {
	// 映射数据，支持嵌套
	Data map[string]interface{}
	// 用于分割查询用key字符串，保存的map数据里不允许有该字符串
	KeySep   string
	ArraySep string
	sync.RWMutex
}

var (
	ErrNotExistKey = errors.New("Not exist key")
	ErrInvalidKey  = errors.New("Invalid key")
)

func NewMap() Map {
	m := Map{
		Data:     make(map[string]interface{}),
		KeySep:   ".",
		ArraySep: ",",
	}
	return m
}

func parseElem(v reflect.Value) interface{} {
	k := v.Kind()
	// 先取到实际对象
	for k == reflect.Ptr || k == reflect.Interface {
		v = v.Elem()
		k = v.Kind()
	}
	// 如果是结构体的非导出字段，直接返回nil
	if !v.CanInterface() {
		return nil
	}
	switch k {
	case reflect.Struct:
		return parseStruct(v)
	case reflect.Map:
		return parseMap(v)
	default:
		return v.Interface()
	}
}

// parseStruct 解析结构体
func parseStruct(val reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})
	t := val.Type()
	n := t.NumField()
	for i := 0; i < n; i++ {
		m[t.Field(i).Name] = parseElem(val.Field(i))
	}
	return m
}

// parseMap 解析映射
func parseMap(val reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})
	for _, key := range val.MapKeys() {
		if key.CanInterface() {
			v := key.Interface()
			k := ParseString(v)
			m[k] = parseElem(val.MapIndex(key))
		}
	}
	return m
}

// handlePanic 捕获panic避免程序异常退出
func handlePanic(err *error) {
	r := recover()
	if r != nil {
		s, ok := r.(string)
		if ok {
			*err = errors.New(s)
			return
		}
		e, ok := r.(error)
		if ok {
			*err = e
			return
		}
		panic(r)
	}
}

// Parse 从结构体、映射类型对象解析数据
func (m *Map) Parse(obj interface{}) (err error) {
	m.Lock()
	defer m.Unlock()
	defer handlePanic(&err)

	v := reflect.ValueOf(obj)
	k := v.Kind()
	// 先取到实际对象
	for k == reflect.Ptr || k == reflect.Interface {
		v = v.Elem()
		k = v.Kind()
	}
	// 只能解析结构体和映射
	switch k {
	case reflect.Struct:
		m.Data = parseStruct(v)
	case reflect.Map:
		m.Data = parseMap(v)
	default:
		return fmt.Errorf("Can't parse type [%s] as a map", v.Type().String())
	}
	return nil
}

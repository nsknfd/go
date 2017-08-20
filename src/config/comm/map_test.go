package comm

import (
	"bytes"
	"testing"

	"fmt"

	"github.com/nsknfd/go/src/errors/errors"
)

type TestInfo struct {
	UUID   string
	Detail string
}

type TestStruct struct {
	Name    string
	Id      int
	Len     float32
	value   bool
	Info    TestInfo
	Values1 map[string]interface{}
	Values2 map[interface{}]interface{}
}

var testFloat32 = float32(2.2)
var testFloat64 = 5.5

func TestMap_ParseStruct(t *testing.T) {
	var testStruct = TestStruct{
		Name:  "name",
		Id:    1,
		Len:   testFloat32,
		value: true,
		Info: TestInfo{
			UUID:   "uuid",
			Detail: "detail",
		},
		Values1: map[string]interface{}{
			"bool":   true,
			"int":    3,
			"int64":  int64(4),
			"float":  testFloat64,
			"string": "string",
			"map": map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
			},
		},
		Values2: map[interface{}]interface{}{
			true:     true,
			3:        3,
			int64(4): int64(4),
			5:        testFloat64,
			"string": "string",
		},
	}
	var resultMap = map[string]interface{}{
		"Name":           "name",
		"Id":             1,
		"Len":            testFloat32,
		"value":          nil,
		"Info.UUID":      "uuid",
		"Info.Detail":    "detail",
		"Values1.bool":   true,
		"Values1.int":    3,
		"Values1.int64":  int64(4),
		"Values1.float":  testFloat64,
		"Values1.string": "string",
		"Values1.map.k1": "v1",
		"Values1.map.k2": "v2",
		"Values2.true":   true,
		"Values2.3":      3,
		"Values2.4":      int64(4),
		"Values2.5":      testFloat64,
		"Values2.string": "string",
	}

	m := NewMap()
	err := m.Parse(testStruct)
	if err != nil {
		t.Fatal(err)
	}

	for key, val := range resultMap {
		v, err := m.get(key)
		if err != nil {
			t.Fatalf("get key [%v] value [%v] err [%s]", key, val, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", val) {
			t.Fatalf("get key [%v] value, want [%v] got [%v]", key, val, v)
		}
	}
}

func TestMap_ParseMap(t *testing.T) {
	b := bytes.NewBufferString("buffer")
	e := errors.New("error")
	type TestString string
	var s TestString = "string"

	var testMap = map[interface{}]interface{}{
		b: "buffer",
		e: "error",
		s: "string",
	}
	var resultMap = map[string]interface{}{
		"buffer": "buffer",
		"error":  "error",
		"string": "string",
	}

	m := NewMap()
	err := m.Parse(7)
	if err == nil {
		t.Fatalf("should not parse int type")
	}
	err = m.Parse(&testMap)
	if err != nil {
		t.Fatal(err)
	}
	for key, val := range resultMap {
		v, err := m.get(key)
		if err != nil {
			t.Fatalf("get key [%v] value [%v] err [%s]", key, val, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", val) {
			t.Fatalf("get key [%v] value, want [%v] got [%v]", key, val, v)
		}
	}
}

package yaml

import (
	"fmt"
	"os"
	"testing"
)

func TestYaml(t *testing.T) {
	var testYaml = `bool: true
int: 1
int64: 2
float: 3.3
string: "string"
strings: "string1, string2"
`
	var resultMap = map[string]interface{}{
		"bool":    true,
		"int":     1,
		"int64":   int64(2),
		"float":   3.3,
		"string":  "string",
		"strings": []string{"string1", "string2"},
	}

	f, err := os.Create("test.yaml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(testYaml)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("test.yaml")

	config := &Config{}
	configer, err := config.ParseFile("test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	for key, val := range resultMap {
		var err error
		var v interface{}
		switch val.(type) {
		case bool:
			v, err = configer.GetBool(key)
		case int:
			v, err = configer.GetInt(key)
		case int64:
			v, err = configer.GetInt64(key)
		case float64:
			v, err = configer.GetFloat(key)
		case string:
			v, err = configer.GetString(key)
		case []string:
			v, err = configer.GetStrings(key)
		default:
			t.Fatalf("not support type")
		}
		if err != nil {
			t.Fatalf("get key [%v] value [%v] err [%s]", key, val, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", val) {
			t.Fatalf("get key [%v] value, want [%v] got [%v]", key, val, v)
		}
	}
}

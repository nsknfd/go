package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Test struct {
	ID   int
	Name string
}

func ReadValue(file string, value interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	data := make([]byte, 1000)
	n, err := f.Read(data)
	if err != nil {
		return err
	}
	fmt.Println("Read bytes:", n)
	data = data[0:n]
	err = json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return nil
}
func WriteValue(file string, value interface{}) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	fmt.Println("Marshal:", string(data))
	n, err := f.Write(data)
	if err != nil {
		return err
	}
	fmt.Println("Write bytes:", n)
	return nil
}

type Info struct {
	Id   string
	Name string
}

type Detail struct {
	Info   Info
	Num    int
	Score  float32
	Slice  []string
	Kv     map[string]Info
	InfoP  *Info
	IntP   *int
	SliceP *[]float32
	KvP    *map[string]time.Time
}

func main() {
	i := 4
	obj := &Detail{
		Info{"11", "aaa"},
		1,
		2.2,
		[]string{"a", "b", "c"},
		map[string]Info{"a": Info{"1", "a"}, "b": Info{"2", "b"}},
		&Info{"3", "c"},
		&i,
		&[]float32{5.5, 6.6, 7.7},
		&map[string]time.Time{"t1": time.Now(), "t2": time.Now()},
	}

	//obj := &Test{11, "aa"}
	file := "main.json"
	err := WriteValue(file, obj)
	if err != nil {
		fmt.Println("Write Test failed.", err)
	} else {
		fmt.Println("Write Test success.", obj)
	}
	m := make(map[string]interface{})
	err = ReadValue(file, &m)
	if err != nil {
		fmt.Println("Read Test failed.", err)
	} else {
		fmt.Println("Read Test success.", m)
	}

}

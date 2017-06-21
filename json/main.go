package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Test struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func ReadTest(file string) (*Test, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var data []byte
	n, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Println("Read bytes:", n)
	test := &Test{}
	err = json.Unmarshal(data, test)
	if err != nil {
		return nil, err
	}
	return test, nil
}
func WriteTest(file string, test *Test) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	data, err := json.Marshal(test)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err != nil {
		return err
	}
	fmt.Println("Write bytes:", n)
	return nil
}

func main() {
	test := &Test{"maoyufeng", 283059}
	file := "main.json"
	err := WriteTest(file, test)
	if err != nil {
		fmt.Println("Write Test failed.", err)
	} else {
		fmt.Println("Write Test success.", test)
	}
	test1, err := ReadTest(file)
	if err != nil {
		fmt.Println("Read Test failed.", err)
	} else {
		fmt.Println("Read Test success.", test1)
	}

}

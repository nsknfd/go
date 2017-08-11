package main

import (
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/nsknfd/go/src/errors/errors"
)

func A() error {
	return errors.New("A failed")
}

func B() error {
	_, err := os.Open("b b b b")
	return errors.New("B failed:", err)
}

func C() error {
	_, err := os.Open("c c c c")
	return errors.New(err)
}

func main() {
	errors.Init(true, 2)
	var err error
	err = A()
	logs.Error(err)
	err = B()
	logs.Error(err)
	err = C()
	logs.Error(err)
}

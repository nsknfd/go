package main

import (
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/nsknfd/go/src/errors/errors"
)

var ErrMulti = errors.New("Failed")

func A() error {
	return errors.New("A")
}

func B() error {
	_, err := os.Open("b b b b")
	return errors.New("B:", err)
}

func C() error {
	_, err := os.Open("c c c c")
	return errors.New(err)
}

func D() error {
	return ErrMulti.Trace()
}

func E() error {
	return ErrMulti
}

func main() {
	var err error
	err = A()
	logs.Error(err)
	err = B()
	logs.Error(err)
	err = C()
	logs.Error(err)
	err = D()
	logs.Error(err)
	err = E()
	logs.Error(err)
}

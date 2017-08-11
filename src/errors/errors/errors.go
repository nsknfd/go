package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type CallStackError interface {
	error
	Trace() CallStackError
}

type callStackError struct {
	message string
	stack   []uintptr
}

const (
	skipStack int = 2
	maxStack  int = 10
)

func New(f interface{}, v ...interface{}) CallStackError {
	err := &callStackError{}
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
	default:
		msg = fmt.Sprint(f)
	}

	if len(v) > 0 {
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}

	}
	err.message = fmt.Sprintf(msg, v...)
	return err
}

func (e *callStackError) Trace() CallStackError {
	// 添加堆栈信息时新建一个error
	err := &callStackError{}
	err.message = e.message
	stack := make([]uintptr, maxStack)
	l := runtime.Callers(skipStack, stack)
	err.stack = stack[:l]
	return err
}

func (e *callStackError) Error() string {
	message := e.message
	for _, pc := range e.stack {
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(pc)
		message += fmt.Sprintf("\n\t%s:%d %s", file, line, f.Name())
	}
	return message
}

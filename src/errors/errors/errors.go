package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type CallStackError interface {
	error
}

type callStackError struct {
	message string
}

const (
	skipStack         int = 2
	defaultStackLevel int = 5
	minStackLevel     int = 1
	maxStackLevel     int = 50
)

var (
	withStack  bool = false
	stackLevel int  = defaultStackLevel
)

func Init(withCallStack bool, callStackLevel int) {
	withStack = withCallStack
	if callStackLevel < minStackLevel || callStackLevel > maxStackLevel {
		stackLevel = defaultStackLevel
	} else {
		stackLevel = callStackLevel
	}
}
func formatError(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

func New(f interface{}, v ...interface{}) CallStackError {
	e := &callStackError{}

	e.message = formatError(f, v...)

	if withStack {
		stack := make([]uintptr, stackLevel)
		l := runtime.Callers(skipStack, stack)
		for i := 0; i < l; i++ {
			f := runtime.FuncForPC(stack[i])
			file, line := f.FileLine(stack[i])
			e.message += fmt.Sprintf("\n\t%s:%d %s", file, line, f.Name())
		}
	}
	return e
}

func (e *callStackError) Error() string {
	return e.message
}

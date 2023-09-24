package errors

import (
	"fmt"
	"runtime"
	"strings"
)

func newStackTrace() stacktrace {
	const depth = 16

	pcs := make([]uintptr, depth)
	n := runtime.Callers(3, pcs)
	return pcs[:n:n]
}

type stacktrace []uintptr

func (s stacktrace) String() string {
	var buf strings.Builder

	frames := runtime.CallersFrames(s)
	for {
		frame, more := frames.Next()
		buf.WriteString(fmt.Sprintf("> %s\t%s:%d\n", frame.Func.Name(), frame.File, frame.Line))
		if !more {
			break
		}
	}
	return buf.String()
}

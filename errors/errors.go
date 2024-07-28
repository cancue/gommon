package errors

import (
	"errors"
	"fmt"
	"strings"
)

func New(a ...any) error {
	return &fundamental{
		msg:   fmt.Sprint(a...),
		stack: newStackTrace(),
		err:   nil,
	}
}

func Newf(format string, a ...any) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, a...),
		stack: newStackTrace(),
		err:   nil,
	}
}

type fundamental struct {
	msg   string
	stack stacktrace
	err   error
}

func (f *fundamental) Error() string {
	if f.err != nil {
		return fmt.Sprintf("%s: %s", f.msg, f.err.Error())
	}
	return f.msg
}

func (f *fundamental) Unwrap() error {
	return f.err
}

func (f *fundamental) Format(s fmt.State, verb rune) {
	if verb == 'v' && s.Flag('+') {
		s.Write([]byte(formatChain(f)))
		return
	}

	s.Write([]byte(f.Error()))
}

func Wrap(cause error, a ...any) error {
	if cause == nil {
		return nil
	}

	return &fundamental{
		msg:   fmt.Sprint(a...),
		stack: newStackTrace(),
		err:   cause,
	}
}

func Wrapf(cause error, format string, a ...any) error {
	if cause == nil {
		return nil
	}

	return &fundamental{
		msg:   fmt.Sprintf(format, a...),
		stack: newStackTrace(),
		err:   cause,
	}
}

func Cause(err error) error {
	for err != nil {
		e, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			return err
		}
		unwrapped := e.Unwrap()
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
	return nil
}

func formatChain(err error) string {
	var buf strings.Builder
	for err != nil {
		if f, ok := err.(*fundamental); ok {
			msg := f.msg
			if msg != "" {
				msg += "\n"
			}
			buf.WriteString(fmt.Sprintf("%s%v", msg, f.stack))
			err = f.err
		} else {
			buf.WriteString(fmt.Sprintf("%s\n", err.Error()))
			err = nil
		}
	}
	return buf.String()
}

var Is = errors.Is
var As = errors.As
var Unwrap = errors.Unwrap

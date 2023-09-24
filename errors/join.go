package errors

import (
	"bytes"
	"fmt"
)

func Join(errs ...error) error {
	var joinedErrs Joined
	for _, err := range errs {
		if err == nil {
			continue
		}
		if j, ok := err.(Joined); ok {
			joinedErrs = append(joinedErrs, j...)
			continue
		}
		joinedErrs = append(joinedErrs, err)
	}
	return joinedErrs
}

type Joined []error

func (errs Joined) Error() string {
	var buf bytes.Buffer

	if len(errs) > 1 {
		fmt.Fprintf(&buf, "%d errors: ", len(errs))
	}
	for i, err := range errs {
		if i != 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(err.Error())
	}

	return buf.String()
}

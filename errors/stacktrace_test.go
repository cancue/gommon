package errors

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacktraceOutput(t *testing.T) {
	st := caller()
	output := st.String()
	reg := regexp.MustCompile(`> github\.com\/cancue\/gommon\/errors\.TestStacktraceOutput	.*\/errors\/stacktrace_test\.go:12+`)
	assert.True(t, reg.MatchString(output))
	assert.Equal(t, len(st), len(strings.Split(output, "\n"))-1)
}

func caller() stacktrace {
	return newStackTrace()
}

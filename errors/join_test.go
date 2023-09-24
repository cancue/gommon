package errors_test

import (
	"testing"

	"github.com/cancue/gommon/errors"
	"github.com/stretchr/testify/assert"
)

func TestJoinReturnsNil(t *testing.T) {
	assert.Nil(t, errors.Join())
	assert.Nil(t, errors.Join(nil))
	assert.Nil(t, errors.Join(nil, nil))
}

func TestJoin(t *testing.T) {
	err1 := errors.New("foo")
	err2 := errors.New("bar")
	for _, test := range []struct {
		errs []error
		want errors.Joined
		msg  string
	}{{
		errs: []error{err1},
		want: errors.Joined{err1},
		msg:  "foo",
	}, {
		errs: []error{err1, err2},
		want: errors.Joined{err1, err2},
		msg:  "2 errors: foo; bar",
	}, {
		errs: []error{err1, nil, err2},
		want: errors.Joined{err1, err2},
		msg:  "2 errors: foo; bar",
	}} {
		got := errors.Join(test.errs...).(errors.Joined)
		assert.Equal(t, test.want, got)
		assert.Equal(t, test.msg, got.Error())
	}
}

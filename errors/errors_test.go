package errors_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/cancue/gommon/errors"
	"github.com/stretchr/testify/assert"
)

func ExampleWrap() {
	if err := func() error {
		// Do something...
		return errors.New("error!")
	}(); err != nil {
		wrapped := errors.Wrap(err, "doing something")
		fmt.Println(wrapped)
	}

	// Output: doing something: error!
}

func TestNew(t *testing.T) {
	msg := "foo"

	err := errors.New(msg)
	assert.Error(t, err)
	assert.Equal(t, msg, err.Error())

	t.Run("The %+v triggers stacktrace print", func(t *testing.T) {
		reg := regexp.MustCompile(msg + `[ \n]+> github\.com\/cancue\/gommon\/errors_test\.TestNew	.*\/errors\/errors_test\.go:\d+`)
		assert.True(t, reg.MatchString(fmt.Sprintf("%+v", err)))
	})
}

func TestWrap(t *testing.T) {
	cause := "foo"
	comment := "bar"

	err := errors.New(cause)
	err = errors.Wrap(err, comment)

	assert.Equal(t, comment+": "+cause, err.Error())

	t.Run("should return all stacks", func(t *testing.T) {
		reg := regexp.MustCompile(
			comment + `[ \n]+> github\.com\/cancue\/gommon\/errors_test\.TestWrap	.*\/errors\/errors_test\.go:\d+[[:ascii:]]+` +
				cause + `[ \n]+> github\.com\/cancue\/gommon\/errors_test\.TestWrap	.*\/errors\/errors_test\.go:\d+`)
		assert.True(t, reg.MatchString(fmt.Sprintf("%+v", err)))
	})

	t.Run("should return nil if cause is nil", func(t *testing.T) {
		assert.Nil(t, errors.Wrap(nil, comment))
	})
}

func TestWrappedErrors(t *testing.T) {
	cause := fmt.Errorf("foo")
	wrapped := errors.Wrap(cause, "bar")
	doubled := errors.Wrap(wrapped, "baz")

	t.Run("Unwrap should return the preceding error or nil", func(t *testing.T) {
		assert.Equal(t, wrapped, errors.Unwrap(doubled))
		assert.Equal(t, cause, errors.Unwrap(wrapped))
		assert.Nil(t, errors.Unwrap(cause))
	})

	t.Run("Cause should return the causing error", func(t *testing.T) {
		assert.Equal(t, cause, errors.Cause(doubled))
		assert.Equal(t, cause, errors.Cause(wrapped))
		assert.Equal(t, cause, errors.Cause(cause))
		cause = errors.New("foo")
		assert.Equal(t, cause, errors.Cause(cause))
	})
}

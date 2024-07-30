package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cancue/gommon/json"
)

func TestUnmarshalJSON(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		data := []byte("{\"A\":\"foo\",\"b\":\"bar\",\"c\":\"baz\"}")
		result, err := json.UnmarshalJSON[struct {
			A string
			B string `json:"b"`
			c string
		}](data)
		assert.NoError(t, err)
		assert.Equal(t, "foo", result.A)
		assert.Equal(t, "bar", result.B)
		assert.Equal(t, "", result.c)
	})

	t.Run("should work with plain string", func(t *testing.T) {
		data := []byte("\"foobar\"")
		result, err := json.UnmarshalJSON[string](data)
		assert.NoError(t, err)
		assert.Equal(t, "foobar", *result)
	})

	t.Run("should work with plain number", func(t *testing.T) {
		data := []byte("1234567890123456789")
		result, err := json.UnmarshalJSON[uint64](data)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1234567890123456789), *result)
	})

	t.Run("should work with array", func(t *testing.T) {
		data := []byte("[\"foobar\",3,{\"A\":5}]")
		result, err := json.UnmarshalJSON[[]any](data)
		assert.NoError(t, err)
		assert.Equal(t, "foobar", (*result)[0].(string))
		assert.Equal(t, float64(3), (*result)[1].(float64))
		assert.Equal(t, map[string]any{"A": float64(5)}, (*result)[2])
	})
}

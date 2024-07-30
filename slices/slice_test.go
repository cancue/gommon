package slices_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/cancue/gommon/slices"
	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		a := []string{"a", "b", "c"}
		b := []string{"b", "c", "d"}
		result := slices.Difference(a, b)

		assert.Equal(t, []string{"a"}, result)
	})
}

func TestFilter(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		a := []string{"a", "b", "c"}
		result := slices.Filter(a, func(el string) bool {
			return el == "a"
		})

		assert.Equal(t, []string{"a"}, result)
	})
}

func TestMap(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		a := []string{"a", "b", "c"}
		result := slices.Map(a, func(el string) string {
			return el + el
		})

		assert.Equal(t, []string{"aa", "bb", "cc"}, result)
	})
}

func TestSet(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		set := slices.NewSet[string]([]string{"foo"})
		assert.Equal(t, 1, len(set.Keys()))
		assert.True(t, set.Has("foo"))

		assert.False(t, set.Has("a"))
		set.Add("a")
		assert.True(t, set.Has("a"))

		set.Add("b")
		set.Add("c")
		assert.Equal(t, 4, len(set.Keys()))
	})

	t.Run("NewSetFrom should work", func(t *testing.T) {
		set := slices.NewSetFrom[string, string](
			[]string{"foo", "bar"},
			func(el string) string {
				return el + el
			},
		)
		assert.Equal(t, 2, len(set.Keys()))
		assert.Contains(t, set.Keys(), "foofoo")
		assert.Contains(t, set.Keys(), "barbar")
	})
}

func TestOrderableDict(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		dict := slices.NewOrderableDict[string, string]()
		sort.Strings(dict.Keys)

		assert.Equal(t, 0, dict.Size())

		aValue := "aValue"
		bValue := "bValue"
		cValue := "cValue"
		assert.False(t, dict.Has("b"))
		dict.Add("b", &bValue)
		assert.True(t, dict.Has("b"))
		dict.Add("c", &cValue)
		dict.Add("a", &aValue)

		assert.Equal(t, "[b c a]", fmt.Sprintf("%v", dict.Keys))
		sort.Strings(dict.Keys)
		assert.Equal(t, "[a b c]", fmt.Sprintf("%v", dict.Keys))

		assert.Equal(t, 3, dict.Size())
		assert.Equal(t, &aValue, dict.Get("a"))
		dict.Delete("a")
		assert.Equal(t, 2, dict.Size())
		assert.Nil(t, dict.Get("a"))
	})

	t.Run("NewOrderableDictFrom should work", func(t *testing.T) {
		dict := slices.NewOrderableDictFrom[string, string](
			[]string{"foo", "bar"},
			func(el string) string {
				return el + el
			},
		)
		assert.Equal(t, 2, dict.Size())
		assert.Equal(t, []string{"foofoo", "barbar"}, dict.Keys)
	})

	t.Run("should ignore duplicate keys", func(t *testing.T) {
		dict := slices.NewOrderableDict[string, string]()
		prev := "foo"
		next := "bar"
		dict.Add("a", &prev)
		dict.Add("a", &next)

		assert.Equal(t, 1, dict.Size())
		assert.Equal(t, &prev, dict.Get("a"))
	})
}

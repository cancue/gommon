package slices

import (
	std "slices"
)

func Difference[E comparable](a []E, b []E) (result []E) {
	for _, el := range a {
		if !std.Contains[[]E, E](b, el) {
			result = append(result, el)
		}
	}
	return
}

func Filter[E any](a []E, f func(el E) bool) (result []E) {
	for _, el := range a {
		if f(el) {
			result = append(result, el)
		}
	}
	return
}

func Map[E any, T any](a []E, f func(el E) T) (result []T) {
	for _, el := range a {
		result = append(result, f(el))
	}
	return
}

// Unordered Set
type Set[T comparable] struct {
	set map[T]bool
}

func (s *Set[T]) Has(el T) bool {
	return s.set[el]
}

func (s *Set[T]) Add(el T) {
	s.set[el] = true
}

func (s *Set[T]) Keys() (result []T) {
	for k := range s.set {
		result = append(result, k)
	}
	return
}

func NewSetFrom[E any, T comparable](a []E, f func(el E) T) (result Set[T]) {
	result.set = make(map[T]bool)

	for _, el := range a {
		result.set[f(el)] = true
	}
	return
}

func NewSet[E comparable](a []E) (result Set[E]) {
	result.set = make(map[E]bool)

	for _, el := range a {
		result.set[el] = true
	}
	return
}

// Orderable Dict
type OrderableDict[K comparable, V any] struct {
	dict map[K]*V
	Keys []K
}

func (s *OrderableDict[K, V]) Has(key K) bool {
	return s.dict[key] != nil
}

func (s *OrderableDict[K, V]) Add(key K, value *V) {
	if s.Has(key) {
		return
	}

	s.dict[key] = value
	s.Keys = append(s.Keys, key)
}

func (s *OrderableDict[K, V]) Delete(key K) {
	delete(s.dict, key)
	s.Keys = Filter(s.Keys, func(el K) bool {
		return el != key
	})
}

func (s *OrderableDict[K, V]) Size() int {
	return len(s.Keys)
}

func (s *OrderableDict[K, V]) Get(key K) *V {
	return s.dict[key]
}

func NewOrderableDictFrom[K comparable, V any](values []V, f func(v V) K) (result OrderableDict[K, V]) {
	result.dict = make(map[K]*V)

	for _, value := range values {
		result.Add(f(value), &value)
	}
	return
}

func NewOrderableDict[K comparable, V any]() (result OrderableDict[K, V]) {
	result.dict = make(map[K]*V)
	return
}

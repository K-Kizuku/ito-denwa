package null

import (
	"errors"
)

type Null[T any] struct {
	Value T
	Valid bool
}

func New[T any](value *T) Null[T] {
	if value == nil {
		return Null[T]{Valid: false}
	}
	return Null[T]{Value: *value, Valid: true}
}

func (n Null[T]) IsNull() bool {
	return !n.Valid
}

func (n *Null[T]) SetValue(value *T) {
	if value == nil {
		n.Valid = false
	} else {
		n.Value = *value
		n.Valid = true
	}
}

func (n Null[T]) GetValue() (T, error) {
	if n.IsNull() {
		return n.Value, errors.New("null value")
	}
	return n.Value, nil
}

func (n Null[T]) GetValueOptional() *T {
	if n.IsNull() {
		return nil
	}
	return &n.Value
}

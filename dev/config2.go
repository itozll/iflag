package dev

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func convertFloat[T constraints.Float](s string) (T, error) {
	return _convert[T](func(s string) (float64, error) { return strconv.ParseFloat(s, 64) }, s)
}

func convertBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func convertUnsigned[T constraints.Signed](s string) (T, error) {
	return _convert[T](func(s string) (uint64, error) { return strconv.ParseUint(s, 10, 64) }, s)
}

func convertSigned[T constraints.Signed](s string) (T, error) {
	return _convert[T](func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }, s)
}

func _convert[T, N constraints.Integer | constraints.Float](fn func(string) (N, error), s string) (T, error) {
	v, err := fn(s)
	if err != nil {
		return 0, err
	}

	return T(v), nil
}

var _ Config[int] = (*configInteger[int])(nil)

type configInteger[T constraints.Integer] struct {
	convert func(s string) (T, error)
}

func (cfg configInteger[T]) IsNil(v T) bool  { return false }
func (cfg configInteger[T]) IsZero(v T) bool { return v == 0 }

func (cfg configInteger[T]) Convert(arg string) (T, error) {
	return cfg.convert(arg)
}

func (cfg configInteger[T]) Append(src *T, arg string) error {
	v, err := cfg.Convert(arg)
	if err != nil {
		return err
	}

	*src = v
	return nil
}

package iflag

import (
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Int Flag[int]

func NewIntFlag(f *Int) *Flag[int] {
	_f := (*Flag[int])(f)
	_f.FnStrConvert = strconv.Atoi
	_f.FnAppend = func(v string) { _f.SetWithString(v, _f.ChangeLevel) }

	return _f
}

type Int64 Flag[int64]

func NewInt64Flag(f *Int64) *Flag[int64] {
	_f := (*Flag[int64])(f)
	_f.FnStrConvert = func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }
	_f.FnAppend = func(v string) { _f.SetWithString(v, _f.ChangeLevel) }

	return _f
}

type Bool Flag[bool]

func NewBoolFlag(f *Bool) *Flag[bool] {
	_f := (*Flag[bool])(f)
	_f.FnStrConvert = strconv.ParseBool
	_f.FnAppend = func(v string) { _f.SetWithString(v, _f.ChangeLevel) }

	return _f
}

type String Flag[string]

func NewStringFlag(f *String) *Flag[string] {
	_f := (*Flag[string])(f)
	_f.FnStrConvert = _string2string
	_f.FnAppend = func(v string) { *_f.Destination = v }

	return _f
}
func _string2string(v string) (string, error) { return v, nil }

type StringSlice Flag[[]string]

func NewStringSliceFlag(f *StringSlice) *Flag[[]string] {
	_f := (*Flag[[]string])(f)
	_f.FnStrConvert = _string2stringslice
	_f.FnAppend = func(v string) { *_f.Destination = append(*_f.Destination, v) }

	return _f
}
func _string2stringslice(s string) ([]string, error) { return strings.Split(s, ","), nil }

type Duration Flag[time.Duration]

func NewDurationFlag(f *Duration) *Flag[time.Duration] {
	_f := (*Flag[time.Duration])(f)
	_f.FnStrConvert = _string2duration
	_f.FnAppend = func(v string) { _f.SetWithString(v, _f.ChangeLevel) }

	return _f
}
func _string2duration(s string) (time.Duration, error) {
	if len(s) == 0 {
		return 0, nil
	}

	if unicode.IsDigit(rune(s[len(s)-1])) {
		s += "s"
	}

	return time.ParseDuration(s)
}

type ChangeLevel uint8

const (
	ChangeLevelMin    = iota // DefaultValue
	ChangeLevelEnv           // EnvVar
	ChangeLevel_1            // placeholder_1
	ChangeLevelConfig        // config file
	ChangeLevel_2            // placeholder_2
	ChangeLevelArg           // --arg ...
	ChangeLevelTop
)

type Flag[T any] struct {
	*Flag[T]

	Name         string
	Aliases      []string
	Usage        string
	EnvVars      []string
	Required     bool
	Destination  *T
	DefaultValue T

	_val         T
	FnStrConvert func(string) (T, error)
	FnAppend     func(string)

	ChangeLevel ChangeLevel
}

func (f *Flag[T]) IsChanged() bool {
	return f.ChangeLevel > 0
}

func (f *Flag[T]) Value() T { return *f.Destination }

func (f *Flag[T]) PreExec() {
	if f.Destination == nil {
		f.Destination = &f._val
	}

	*f.Destination = f.DefaultValue

	if len(f.EnvVars) > 0 {
		for _, envVar := range f.EnvVars {
			val := os.Getenv(strings.ToUpper(envVar))
			if val != "" {
				v, err := f.FnStrConvert(val)
				if err != nil {
					continue
				}

				*f.Destination = v
				f.ChangeLevel = ChangeLevelEnv
				return
			}
		}
	}
}

func (f *Flag[T]) SetWithString(v string, changeLevel ChangeLevel) error {
	val, err := f.FnStrConvert(v)
	if err != nil {
		return err
	}

	return f.Set(val, changeLevel)
}

//  无法修改当前级别比 changeLevel 高的数据
func (f *Flag[T]) Set(v T, changeLevel ChangeLevel) error {
	if changeLevel < f.ChangeLevel {
		return nil
	}

	f.ChangeLevel = changeLevel
	*f.Destination = v
	return nil
}

func (f *Flag[T]) Append(v string) {
	if f.FnAppend != nil {
		f.FnAppend(v)
	}
}

func (f *Flag[T]) flagText() string {
	var options []string
	for _, alias := range f.Aliases {
		if len(alias) > 1 {
			options = append(options, "--"+alias)
		} else if len(alias) == 1 {
			options = append(options, "-"+alias)
		}
	}

	options = append(options, "--"+f.Name)
	return strings.Join(options, ",")
}

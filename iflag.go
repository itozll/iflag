package iflag

import (
	"unsafe"

	"github.com/spf13/pflag"
)

var (
	// 单个变量的多重绑定设置（默认禁止）
	// 如果单个变量多重绑定的默认值不一致，可能会导致意料之外的结果
	allowVariableDuplicates = false

	// 当前被绑定变量的地址列表
	binders = map[unsafe.Pointer]bool{}
)

func AllowVariableDuplicates() bool {
	allowVariableDuplicates = true
	return true
}

var (
	NewBytes          = NewFlag((*pflag.FlagSet).BytesBase64VarP).call
	NewBool           = NewFlag((*pflag.FlagSet).BoolVarP).call
	NewBoolSlice      = NewFlag((*pflag.FlagSet).BoolSliceVarP).call
	NewString         = NewFlag((*pflag.FlagSet).StringVarP).call
	NewStringSlice    = NewFlag((*pflag.FlagSet).StringSliceVarP).call
	NewFloat32        = NewFlag((*pflag.FlagSet).Float32VarP).call
	NewFloat32Slice   = NewFlag((*pflag.FlagSet).Float32SliceVarP).call
	NewFloat64        = NewFlag((*pflag.FlagSet).Float64VarP).call
	NewFloat64Slice   = NewFlag((*pflag.FlagSet).Float64SliceVarP).call
	NewDuration       = NewFlag((*pflag.FlagSet).DurationVarP).call
	NewDurationSlice  = NewFlag((*pflag.FlagSet).DurationSliceVarP).call
	NewInt            = NewFlag((*pflag.FlagSet).IntVarP).call
	NewIntSlice       = NewFlag((*pflag.FlagSet).IntSliceVarP).call
	NewInt8           = NewFlag((*pflag.FlagSet).Int8VarP).call
	NewInt16          = NewFlag((*pflag.FlagSet).Int16VarP).call
	NewInt32          = NewFlag((*pflag.FlagSet).Int32VarP).call
	NewInt32Slice     = NewFlag((*pflag.FlagSet).Int32SliceVarP).call
	NewInt64          = NewFlag((*pflag.FlagSet).Int64VarP).call
	NewInt64Slice     = NewFlag((*pflag.FlagSet).Int64SliceVarP).call
	NewUint           = NewFlag((*pflag.FlagSet).UintVarP).call
	NewUintSlice      = NewFlag((*pflag.FlagSet).UintSliceVarP).call
	NewUint8          = NewFlag((*pflag.FlagSet).Uint8VarP).call
	NewUint16         = NewFlag((*pflag.FlagSet).Uint16VarP).call
	NewUint32         = NewFlag((*pflag.FlagSet).Uint32VarP).call
	NewUint64         = NewFlag((*pflag.FlagSet).Uint64VarP).call
	NewIP             = NewFlag((*pflag.FlagSet).IPVarP).call
	NewIPSlice        = NewFlag((*pflag.FlagSet).IPSliceVarP).call
	NewIPNet          = NewFlag((*pflag.FlagSet).IPNetVarP).call
	NewStringToInt    = NewFlag((*pflag.FlagSet).StringToIntVarP).call
	NewStringToInt64  = NewFlag((*pflag.FlagSet).StringToInt64VarP).call
	NewStringToString = NewFlag((*pflag.FlagSet).StringToStringVarP).call

	NewCount = NewFlagNoDefaultValue((*pflag.FlagSet).CountVarP).call
)

type (
	Argument interface {
		Bind(set *pflag.FlagSet)
	}

	handlerNewFlag[T any]  func(set *pflag.FlagSet, p *T, name, shorthand string, value T, usage string)
	handlerNewFlag2[T any] func(set *pflag.FlagSet, p *T, name, shorthand string, usage string)

	flag[T any] struct {
		Name    string
		Alias   string
		Usage   string
		Binder  *T
		pf      handlerNewFlag[T]
		Default T
	}
)

func NewFlag[T any](pf handlerNewFlag[T]) *flag[T] { return &flag[T]{pf: pf} }
func NewFlagNoDefaultValue[T any](pf handlerNewFlag2[T]) *flag[T] {
	return NewFlag(func(set *pflag.FlagSet, p *T, name, shorthand string, _ T, usage string) {
		pf(set, p, name, shorthand, usage)
	})
}

// 设置变量信息
func (f *flag[T]) call(p *T, name, shorthand string, value T, usage string) *flag[T] {
	if p != nil {
		if !allowVariableDuplicates {
			if binders[unsafe.Pointer(p)] {
				panic("重复绑定变量: " + name)
			}

			binders[unsafe.Pointer(p)] = true
		}
	}

	return &flag[T]{
		Name:    name,
		Alias:   shorthand,
		Usage:   usage,
		Binder:  p,
		pf:      f.pf,
		Default: value,
	}
}

// 参数值
func (f *flag[T]) Value() T { return *f.Binder }

func (arg *flag[T]) Bind(set *pflag.FlagSet) {
	if arg.Binder == nil {
		arg.Binder = new(T)
	}

	arg.pf(
		set,
		arg.Binder,
		arg.Name,
		arg.Alias,
		arg.Default,
		arg.Usage,
	)
}

func Bind(set *pflag.FlagSet, args ...Argument) {
	for _, arg := range args {
		arg.Bind(set)
	}
}

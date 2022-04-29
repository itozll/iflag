package dev

type Config[T any] interface {
	// 检测变量是否为空
	IsNil(v T) bool

	// 检测变量是否为0
	IsZero(v T) bool

	// 将字符串回去的为实际类型
	Convert(arg string) (T, error)

	// 追加或替换
	Append(src *T, arg string) error
}

type ConfigNormal[T any] struct {
	Convert func(arg string) (T, error)
}

package iflag

import "strconv"

var (
	IntConfig = Config[int]{
		IsNil:  func(v int) bool { return false },
		IsZero: func(v int) bool { return v == 0 },

		Convert: func(arg string) (int, error) {
			return strconv.Atoi(arg)
		},
		Append: func(src *int, arg string) error {
			v, err := strconv.Atoi(arg)
			if err != nil {
				return err
			}

			*src = v
			return nil
		},
	}
)

type Config[T any] struct {
	// 检测变量是否为空
	IsNil func(v T) bool

	// 检测变量是否为0
	IsZero func(v T) bool

	// 将字符串回去的为实际类型
	Convert func(arg string) (T, error)

	// 追加或替换
	Append func(src *T, arg string) error
}

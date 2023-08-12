package options

import (
	"github.com/itozll/iflag"
)

var (
	// 允许绑定同一变量
	_ = iflag.AllowVariableDuplicates()

	// 定义变量
	Verbose int
	// 定义命令行参数，绑定变量
	OVerbose = iflag.NewCount(&Verbose, "verbose", "v", 0, "verbose ...")
	// 绑定同一变量
	// 如果在此之前没有调用 iflag.AllowVariableDuplicates()，将导致 panic
	OVerbose2 = iflag.NewCount(&Verbose, "verbose", "v", 0, "verbose2 ...")

	// 定义命令行参数，不绑定变量
	// 使用时通过 OToggle.Value() 获取参数值
	OToggle = iflag.NewBool(nil, "toggle", "t", false, "toggle ...")
)

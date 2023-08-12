# iflag 增强 cobra 参数功能

## 安装

```
go get github.com/itozll/iflag
```

## 依赖

 ```
go 1.19+
 ```

## 使用

> 参考 `examples/iflag-test`
+ 已有 `cobra` 项目，或通过 `cobra-cli` 生成新项目
+ 定义命令参数

```
// 允许绑定同一变量
// 该功能是为了支持不同的命令对同一参数的usage差异性
// 但，多个参数绑定同一变量，必须使用相同的初始值，否则可能产生意料之外的初始值
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
```

+ 支持哪些类型

> 支持 cobra 中已有的所有类型，定义方式与 pflag.XXXVarP 相同 （Count是唯一例外，比cobra多了无意义的默认值参数）

```
var (
    // 第 1 个为绑定的变量地址，为 nil 时表示不绑定，通过 .Value() 获取参数值
    // 第 2 个为长参数名，verbose
    // 第 3 个为短参数名，v
    // 第 4 个为参数默认值
    // 第 5 个为参数提示信息
    _ = iflag.NewCount(&Verbose, "verbose", "v", 0, "verbose ...")
)
```

+ 与 `cobra` 差异

    + cobra 在 init() 中命令添加参数，当多个命令使用相同的参数时，需要多次定义，一旦需要修改会比较繁琐
    + iflag 通过预先定义变量，达到仅需修改一处的目的

```
// cobra
rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

// iflag
package options
var OToggle = iflag.NewBool(nil, "toggle", "t", false, "toggle ...")
var OInfo = iflag.NewString(nil, "info", "i", "default", "info message ...")

--
iflag.Bind(rootCmd.PersistentFlags(), options.OVerbose, options.OInfo)
```

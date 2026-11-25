# Caculator

使用 Go 和 MyGo 原生 UI 编写的桌面计算器。保留 Qt 版的键位布局和逐步计算方式，顶部显示当前输入以及待执行的运算。

## 运行

安装 Go 1.27.3 或更高版本，在项目根目录运行：

```sh
go run .
```

开发时可用 `go tool mygo dev` 启用热重载。MyGo 及其命令行工具的版本由 go.mod 固定。

## 按键

| 操作 | 键盘 |
| --- | --- |
| 输入数字或小数点 | `0`–`9`、`.` |
| 加、减、乘、除 | `+`、`-`、`*`、`/` |
| 计算结果 | `Enter` 或 `=` |
| 删除最后一位 | `Backspace` |
| 清空 | `Esc` |

连续运算按输入顺序执行，例如 `2 + 3 × 4 =` 得到 `20`。连续按两个运算符时采用后一个。除以零时显示 `Error`，直接输入数字或小数点即可重新开始。

## 测试与 macOS 打包

```sh
go test ./...
go tool mygo build -skip-dmg
open "build/darwin-$(go env GOARCH)/计算器.app"
```

运行 `go tool mygo build` 可同时生成磁盘映像。

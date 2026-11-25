# Caculator 开发约定

Caculator 是用 Go 和 MyGo 原生 UI 实现的桌面计算器。应用、计算引擎和测试目前都在仓库根目录的 `main` 包中。

## 开发命令

从仓库根目录使用 mise。`mise.toml` 固定开发用的 Go 和 golangci-lint 版本，`go.mod` 声明最低 Go 版本并固定 MyGo 依赖及 CLI。首次使用时执行 `mise trust` 和 `mise install`。

| 命令 | 用途 |
| --- | --- |
| `mise run start` | 启动应用 |
| `mise run dev` | 启动带热重载的 MyGo 开发环境 |
| `mise run fmt` | 应用 `modernize` 自动修复并用 `gofmt` 格式化，会写回源码 |
| `mise run lint` | 检查 Go 代码和格式 |
| `mise run vet` | 运行 Go 静态检查 |
| `mise run test` | 运行计算逻辑与界面测试 |
| `mise run build` | 为当前平台打包，产物位于 `build/` |

修改 Go 源码后运行 `fmt`、`lint` 和 `test`。涉及窗口启动或打包的改动再运行 `build`。

## 代码入口和行为

- [calculator.go](calculator.go) 管理逐位输入和运算状态。`entry` 保留 `0.`、`12.00` 等编辑状态，内部累加器用 `math/big.Rat` 保存精确值。链式计算直接使用累加器，循环小数和过长结果只在显示时舍入。
- 连续运算按输入顺序进行，`2 + 3 × 4 =` 得到 `20`。等待下一操作数时再按运算符会替换上一个，连续按 `=` 会保留当前结果。除以零时显示 `Error`，可以用数字或小数点开始下一次输入。
- [view.go](view.go) 构建显示区和键盘。屏幕按钮和物理键盘通过 `Calculator.Press` 调用同一套输入逻辑。
- [main.go](main.go) 创建可缩放的窗口。布局变更应检查最小窗口、窄高及宽矮窗口中的长数字、待运算提示和按键。
- 计算行为变更用 [calculator_test.go](calculator_test.go) 中的按键序列验证。键鼠交互和缩放使用 [view_test.go](view_test.go) 中的 MyGo `ui.NewTester`。

## 代码风格

- 用有含义的常量表达控制字符、输入限制、显示精度和布局约束。含义直接的数字键、运算符和配色值可以保留字面量。
- 注释解释行为背后的原因和边界条件，例如为何退格可按字节删除，以及为何继续计算时要使用显示舍入前的值。

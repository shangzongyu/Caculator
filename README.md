# Caculator

使用 Go 和 MyGo 原生 UI 编写的桌面计算器。保留 Qt 版的键位布局和逐步计算方式，顶部显示当前输入以及待执行的运算。窗口可拖动边缘调整大小，也可最大化，键盘和显示区会随之调整。

## 开发与运行

先安装 [mise](https://mise.jdx.dev/getting-started.html)，然后在项目根目录执行：

```sh
mise trust
mise install
mise run start
```

mise.toml 固定项目使用的 Go 和 golangci-lint 版本。MyGo 和它的命令行工具由 Go 模块管理，首次运行会自动下载。

| 命令 | 用途 |
| --- | --- |
| `mise run start` | 启动计算器 |
| `mise run dev` | 启动应用并监听源码变化，保存后热重载 |
| `mise run test` | 运行计算逻辑和界面测试 |
| `mise run fmt` | 写入 `modernize` 的自动修复，并用 `gofmt` 格式化 Go 源码 |
| `mise run lint` | 运行 golangci-lint，检查代码和 Go 格式 |
| `mise run vet` | 运行 Go 静态检查 |
| `mise run build` | 为当前平台打包，产物位于 `build/` |

macOS 打包会生成应用和 DMG。Apple Silicon 上可用 `open build/darwin-arm64/计算器.app` 打开构建产物。

## 按键

| 操作 | 键盘 |
| --- | --- |
| 输入数字或小数点 | `0`–`9`、`.` |
| 加、减、乘、除 | `+`、`-`、`*`、`/` |
| 计算结果 | `Enter` 或 `=` |
| 删除最后一位 | `Backspace` |
| 清空 | `Esc` |

连续运算按输入顺序执行，例如 `2 + 3 × 4 =` 得到 `20`。连续按两个运算符时采用后一个。除以零时显示 `Error`，直接输入数字或小数点即可重新开始。

计算使用精确分数。循环小数和过长的结果最多显示 15 位有效数字，继续运算时仍使用精确值。

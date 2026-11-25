# Caculator

使用 C++ / Qt Widgets 编写的桌面计算器，支持基本的加减乘除。

## 构建依赖

- Qt 6（Core、Gui、Widgets 模块及 qmake）
- C++ 编译器和 make；macOS 可使用 Xcode 提供的工具链

本机已使用 Apple Silicon、Qt 6.12.0 和 Xcode 工具链编译、运行验证。
无需额外安装 Qt Creator，也不需要 Qt 5 兼容模块。

## macOS 编译与运行

在项目根目录执行：

```sh
mkdir -p build-local
cd build-local
qmake ../Caculator.pro CONFIG+=release
make -j4
open bin/Caculator.app
```

也可以直接运行可执行文件，查看终端日志：

```sh
./bin/Caculator.app/Contents/MacOS/Caculator
```

界面样式和图片均已嵌入 Qt 资源文件，启动时无需复制图片或设置工作目录。
连续运算按输入顺序执行。
应用使用无边框窗口，可拖动顶部区域移动窗口，macOS 使用 `⌘Q` 退出。
生成的应用依赖本机 Qt 动态库；若要分发到其他 Mac，需另行使用
`macdeployqt` 打包依赖。

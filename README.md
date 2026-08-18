<div align="center">
  <!-- 这里引用了开源社区著名的 Gopher 吉祥物图片，非常可爱和有活力 -->
  <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GO_LEARN.png" alt="Go Learn Gopher" width="300" />

  <h1>🚀 我的 Go 语言奇妙探索之旅 (LearnGo)</h1>

  <p>
    记录我从零开始征服 Go (Golang) 的点点滴滴！在这里，没有枯燥的代码，只有不断解锁新技能的快乐！🎉
  </p>

  <!-- 炫酷的徽章 -->
  <p>
    <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go" />
    <img src="https://img.shields.io/badge/Status-Learning-success?style=for-the-badge" alt="Status" />
    <img src="https://img.shields.io/badge/Author-CJJ0206-FF69B4?style=for-the-badge" alt="Author" />
  </p>
</div>

---

## 🌟 为什么是 Go？

Go 语言就像代码界的“瑞士军刀”——**简洁、高效、天生支持高并发**！无论是写微服务还是小工具，跑起来都像一阵风 🌪️。

## 🗺️ 我的打怪升级路线图

这里是我接下来的学习计划，每完成一项我就会打个勾 `[x]`！

### 🛡️ 新手村：基础语法
- [x] Hello World! (万物起源)
- [x] 变量、常量与基本数据类型 (int/float/bool/string，零值，`unsafe.Sizeof`)
- [x] 字符串专题 (不可变性、反引号原始字符串、拼接的坑)
- [x] 类型转换 (`fmt.Sprintf` / `strconv` 全家桶)
- [x] 运算符 (算术/关系/逻辑/位运算，二进制补码这种硬核细节也啃了)
- [x] 指针 (Pointer) 入门 (取地址、解引用、多级指针)
- [x] 包与可见性 (多包项目结构，导出规则，别名导入)
- [x] 输入输出 (`Scanln` / `Scanf` 花式读取)
- [x] 流程控制 (if 单/双/多分支、switch 与 fallthrough) — `P76_FlowControl`
- [x] 循环控制 (for 三种写法、for-range、while/do-while 模拟、嵌套循环)
- [x] 跳转控制 (break/continue 及标签用法，goto 了解但不用)
- [x] 函数入门 (带参数与返回值，函数内 switch 分支) — `P108`
- [x] 函数进阶 (多返回值、可变参数、闭包、defer) — `P122`
- [x] 数组、切片 (Slice) 和 字典 (Map) — `P142`(数组/切片)、`P169`(二维数组)、`P173`(Map)

### ⚔️ 进阶之路：面向对象与接口
- [x] 结构体 (Struct) 与方法 — `P182`(结构体)、`P191`(方法)
- [ ] 接口 (Interface) 的奇妙用法
- [x] 错误处理与 Panic/Recover — `P139/errorHandling`(panic/recover)、`P122`(defer)

### 🧬 面向对象进阶（设计篇）
承接上面的结构体与方法，进一步用 Go 模拟经典 OOP 与设计模式，集中放在 `P200/` 下：
- [x] 封装 (Encapsulation) — `P200/Encapsulation`
- [x] 继承 (Inheritance) — `P200/inherit`（含多级嵌套继承 `deeper` / `deeper2`）
- [x] 抽象类 / 抽象方法 (Abstract) — `P200/abstract`
- [x] 工厂模式 (Factory Mode) — `P200/factoryMode`
- [x] 对象转换 (toObject) — `P200/toObiect`

### 🚀 核心大招：并发编程 (Concurrency)
- [ ] Goroutine (轻量级线程，Go 的灵魂)
- [ ] Channel (通道通信)
- [ ] sync 包 (锁与并发同步)

### 🛠️ 实战演练
- [ ] 编写简单的命令行工具 (CLI)
- [ ] 搭建一个简单的 Web 服务器
- [ ] 连接数据库增删改查

### 🎲 支线任务：好玩的尝试
- [x] 用 [goja](https://github.com/dop251/goja) 在 Go 里跑一段 JavaScript，体验一下跨语言的乐趣

---

## 📚 每日巩固

学到的知识点容易忘，所以给自己配了两套“每日”练习，都按日期分文件夹存放，并且和源码 `package` 旁的日期注释（如 `// 8.13`）一一对应，方便回看某天学了什么：

- **每日测试 `ZDailyTest/`**：巩固当天 + 之前容易遗忘的内容，目录名形如 `2026-8-12`。目前覆盖 7/31、8/1–8/7、8/9、8/11、8/12、8/17（部分日期改用下面的每日函数）。
- **每日函数 `ZDailyFunction/`**：每天一个小函数专题，目录名形如 `8_14`，覆盖 `sync.OnceFunc`、对象克隆等进阶用法。

---

## 💻 怎么运行这里的代码？

如果你想克隆并在本地运行我的代码，只需要确保你安装了 [Go 环境](https://go.dev/)，然后在终端输入：

```bash
# 1. 克隆仓库
git clone https://github.com/CJJ0206/LearnGo.git

# 2. 进入目录
cd LearnGo

# 3. 项目按知识点分文件夹存放（如 P76_FlowControl、P108、P200 等），
#    进入对应文件夹的 main 目录运行即可，例如：
cd P76_FlowControl/main
go run .
```

> 整个仓库共用根目录下的 `go.mod`（module 名为 `learn`），所以各个文件夹之间可以互相 import 本地包～

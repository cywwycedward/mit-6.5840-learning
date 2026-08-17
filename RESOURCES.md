# Go 后端与 MIT 6.5840 资源

## Knowledge

- [MIT 6.5840 官方课程调研](reference/mit-6.5840-official-research.md)
  课程定位、实验依赖和适合当前起点的分阶段门槛；每次规划课程主线时先查阅。
- [MIT 6.5840：课程主页](https://pdos.csail.mit.edu/6.5840/)
  关于课程范围、先修和实验的权威入口；核实课程事实时使用。
- [MIT 6.5840：Go 环境说明](https://pdos.csail.mit.edu/6.824/labs/go.html)
  实验使用 Go 1.22+；Windows 的实验环境为 WSL2。开始课程实验环境配置时使用。
- [A Tour of Go](https://go.dev/tour/)
  Go 官方交互式入门；用于补齐语法、方法、接口和并发原语的基础。
- [Go 官方入门教程](https://go.dev/doc/tutorial/getting-started)
  第一次创建 Go 练习项目时使用：`go mod init <模块路径>` 在项目根目录创建依赖清单 `go.mod`；`go run .` 编译并运行当前目录的 `main` 程序。
- [Go Data Race Detector](https://go.dev/doc/articles/race_detector)
  数据竞争的定义、`-race` 的使用边界和报告解释；所有并发练习和课程实验的必读工具文档。
- [`testing` 标准库文档](https://pkg.go.dev/testing)
  `go test` 如何发现 `TestXxx` 函数，以及 `testing.T`、`Errorf` 和 `Fatalf` 的语义；第一次阅读或编写测试时使用。
- [Go 语言规范：函数、方法与 `for`/`range`](https://go.dev/ref/spec)
  函数参数、方法接收者和循环语法的权威定义；对语法位置或行为有疑问时使用。
- [`sync` 标准库文档](https://pkg.go.dev/sync)
  `Mutex`、`WaitGroup`、`Cond` 等同步原语的精确语义；实现并发状态时查阅。
- [Pro Git（官方中文书）](https://git-scm.com/book/zh/v2)
  Git 的提交、分支、合并和远端协作的权威基础；执行本工作区的 Git 规范和复盘历史时使用。
- [`gitignore` 官方文档](https://git-scm.com/docs/gitignore)
  忽略规则的精确匹配语义；调整 `.gitignore` 前查阅，避免误排除学习成果。

## Wisdom (Communities)

- 暂未选定。先通过短练习建立能提出具体问题的基础；之后再选择适合中文交流偏好的 Go 社区或线下活动。

## Gaps

- 尚未建立真实 Go 服务的部署、可观测性和数据库实践材料；在完成 Lab 1–2 的并发与 RPC 基础后补充。

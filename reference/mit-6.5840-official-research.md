# MIT 6.5840 官方课程调研与入门路线

> 调研日期：2026-08-17。除本仓库已有的课表速览外，本文件的外部来源均为 MIT PDOS 的课程官方页面。课程曾名为 6.824；2023 年起改为 6.5840，但官网 URL 仍多保留 `/6.824/`。完整 2026 春季日程见同目录的 [`schedule-spring-2026.md`](schedule-spring-2026.md)。

## 课程定位与学习形式

- 6.5840 是一门 12-unit 的研究生核心课程，包含讲座、阅读、编程实验、可选期末项目、期中和期末；重点是构建分布式系统所需的容错、复制与一致性抽象和实现技术。[课程主页](https://pdos.csail.mit.edu/6.824/)
- 多数课堂将论文讨论和讲授结合；学生应课前阅读指定论文。实验约每一到两周一次，目的是深入理解概念，并获得编写、调试分布式系统的经验。[课程说明](https://pdos.csail.mit.edu/6.824/general.html)
- 官方建议的先修基础是计算机系统（6.1910/6.004）以及 6.1800/6.033 或 6.1810 中至少一门，且强调熟练的调试、实现与软件设计能力。也就是说，这不是以 Go 或分布式系统零基础为目标设计的入门课。[课程说明](https://pdos.csail.mit.edu/6.824/general.html)

## 课程材料与实践主线

2026 春季的讲授/阅读顺序是：MapReduce → RPC 与线程/Go → GFS → Paxos → Go 模式 → Raft → 线性一致性 → ZooKeeper → 分布式事务 → Spanner → Chain Replication → 乐观并发控制 → 验证 → 缓存一致性 → Lambda/Ray → Fork consistency → Bitcoin → BFT。课程课表直接链接对应论文、讲义和问题；未来日期的材料可能调整，应以官网为准。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html)

实验按一个逐层扩展的系统推进：

| 实验 | 你将构建的东西 | 必须掌握的核心问题 |
| --- | --- | --- |
| Lab 1 MapReduce | 通过 RPC 协调的 coordinator 和并发 worker；worker 故障后重新分配任务 | RPC、并发、任务状态、重试、原子落盘；测试会检查并行与 crash recovery。[实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html) |
| Lab 2 KV Server | 单机线性一致 KV 与基于 KV 的锁 | 至多一次语义、丢包/重传、并发客户端与线性一致性。[实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-kvsrv1.html) |
| Lab 3 Raft | 复制状态机协议 | 选主、日志复制、多数派、持久化、快照；之后的实验都以它为基础。[实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-raft1.html) |
| Lab 4 KV Raft | 基于 Raft 的容错 KV | 将请求提交给复制日志、应用已提交操作、处理领导者变化/分区与快照。[实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-kvraft1.html) |
| Lab 5 Sharded KV | 多个 Raft 副本组构成的分片 KV | 配置变更、数据迁移与 controller 故障期间保持线性一致；它可替代课程项目。[实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-shard1.html) |

## 针对当前学习者的起点判断

已具备后端项目经验，但没有 Go 项目、并发编程和分布式系统基础。这意味着 HTTP、请求处理和常见后端业务经验是优势；但不宜直接从 Raft 开始。Lab 1 已同时要求并行执行和 worker 崩溃恢复，Lab 2 已要求在线性一致语义下处理网络故障，因此应先建立 Go 并发和故障语义的共同语言，再进入实验主线。这是依据官方实验依赖得出的学习建议，并非 MIT 的正式入学要求。[Lab 1](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html) [Lab 2](https://pdos.csail.mit.edu/6.824/labs/lab-kvsrv1.html)

## 建议路线：先补基础，再按实验推进

以下分阶段目标是根据上述官方前置条件、课表顺序和实验依赖作出的建议；可按每周可投入时间伸缩，不把它当作课程官方进度。

1. **Go 起步（先完成，不跳过）**：安装并使用 Go 1.22+，完成课表指定的 Online Go Tour；随后练熟 module、结构体/接口、错误处理、文件 I/O、`go test`、`go test -race`。在 Windows 上应使用 WSL2；官方明确 WSL1 不适用于实验。[Go 环境说明](https://pdos.csail.mit.edu/6.824/labs/go.html) [RPC 与线程课前任务](https://pdos.csail.mit.edu/6.824/schedule.html)
2. **并发与本机服务（达到可做 Lab 1 的门槛）**：用小练习掌握 goroutine、channel、`sync.Mutex`、`sync.Cond`、定时器、取消/退出以及数据竞争。目标是能解释某共享状态由谁保护、何时阻塞、为何不会泄漏 goroutine，并用 race detector 验证。课程 Lab 1 特别提示 coordinator 的 RPC server 是并发的，必须保护共享数据。[Lab 1 提示](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html)
3. **网络故障语义（Lab 1–2）**：先读 MapReduce，再实现 Lab 1；接着以 Lab 2 建立“请求重发不应重复执行”和“线性一致”的直觉，并实现 KV 锁。完成标准不是仅让测试绿，而是能画出一次 RPC 丢失回复时 client/server 的状态变化，并说明为何结果正确。[Lab 2 语义](https://pdos.csail.mit.edu/6.824/labs/lab-kvsrv1.html)
4. **一致性与复制（Lab 3）**：阅读课表指定的 Raft extended paper（重点 Figure 2），先写出 follower/candidate/leader、term、日志和提交下标的状态表，再实现选主、复制、持久化和快照。官方要求使用测试驱动实现，并建议始终用 `-race` 测试。[Raft 实验说明](https://pdos.csail.mit.edu/6.824/labs/lab-raft1.html)
5. **端到端存储系统（Lab 4–5）**：在已稳定的 Raft 上实现容错 KV，再处理分片、迁移和配置变更。此阶段才把课程论文扩展到事务、Spanner、链式复制等横向比较：它们都在回答“故障、并发、扩展性”三者如何权衡的问题。[Lab 4](https://pdos.csail.mit.edu/6.824/labs/lab-kvraft1.html) [Lab 5](https://pdos.csail.mit.edu/6.824/labs/lab-shard1.html)

## 每阶段的学习产出

| 阶段 | 可验证产出 | 进入下一阶段的检查问题 |
| --- | --- | --- |
| Go 与并发 | 一个带单元测试和 `-race` 的本机并发程序 | 能否定位并修复 race、死锁或 goroutine 无法退出？ |
| Lab 1–2 | 通过官方测试，并能用时序图解释重试和超时 | 网络可能丢失任一消息时，是否仍能说清“正确”是什么？ |
| Lab 3 | 通过 Raft 分阶段测试 | 为什么多数派是进展条件？故障节点恢复后如何追上日志？ |
| Lab 4–5 | 一个可承受测试中故障/分区的 KV 系统 | 客户端从旧 leader 或旧分片组得到回复时，如何防止破坏线性一致？ |

## 使用课程材料的边界

如果按官方课程实验学习，应独立完成代码、理解并能解释全部提交内容，也不要公开实验解答；课程说明明确将这些要求作为协作政策的一部分。[协作政策](https://pdos.csail.mit.edu/6.824/general.html)

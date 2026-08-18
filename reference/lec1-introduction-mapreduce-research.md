# L1：Introduction 与 MapReduce 调研笔记

> 调研日期：2026-08-17。仅使用 MIT 6.5840 官方课表、L1 讲义、Lab 1 说明，以及课程链接的 Google 原始论文；用于设计课程，不提供 Lab 1 的实现方案。

## L1 的准确范围

- 2026 春季官方课表将第一讲列为 **LEC 1: Introduction**；课前阅读是 *MapReduce (2004)*，同日发布 Lab 1: MapReduce。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-3)
- 官方 L1 讲义先介绍分布式系统、课程关注的基础设施（storage、communication、computation），以及 fault tolerance、consistency、performance、trade-offs 和 implementation；随后把 MapReduce 作为串起这些主题、也是 Lab 1 焦点的 case study。[L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)

因此，归属 L1 的课程应忠实教学“分布式系统为什么需要这些抽象”以及“MapReduce 如何作为入门案例”，而不应把 L1 窄化成完整的 Lab 1 编程课。MapReduce 可拆成多节小课，但它们都应服务于这个 Introduction 的范围。

## MapReduce：应用写什么，框架负责什么

MapReduce 是处理、生成大数据集的 programming model 及其实现。用户只提供两个函数；运行时负责分割输入、调度、跨机器通信、并行执行和机器故障处理。[MapReduce 原始论文：摘要与 §1](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

| 部分 | 输入与输出 | 责任 |
| --- | --- | --- |
| `Map` | `(k1, v1) → list(k2, v2)` | 对一个输入记录发出零个或多个 intermediate key/value pairs。 |
| shuffle / grouping | 所有 intermediate pairs → 按 `k2` 分组 | 框架把同一 intermediate key 的所有 values 汇集到一起。 |
| `Reduce` | `(k2, list(v2)) → list(v2)` | 对一个 key 及其全部 values 做合并，产生最终输出。 |

这些是论文 §2、§2.2 的概念类型，不是 Go 的具体函数签名。[原始论文 §2–2.2](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf) 词频统计是最小例子：每个 `Map` 为见到的每个词发出 `(word, "1")`，同一词的 `Reduce` 将所有计数相加。[原始论文 §2.1](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

## 数据流与并行边界

1. 运行时把输入分为 `M` 个 splits；每个 split 对应一个 Map task，因此多个 Map 可以并行。
2. Map 产生 intermediate pairs，并使用 partition function（例如 `hash(key) mod R`）将 key space 分到 `R` 个分区。
3. 每个 Reduce task 取得属于其分区的所有 Map 输出，按 key 排序/分组，再对每个 key 调用 `Reduce`。
4. 成功完成后，结果是 `R` 个 output files（每个 Reduce task 一个），并非必然合并为一个文件。

这份数据流、`M` 个 Map task 和 `R` 个 Reduce task、以及 Map/Reduce 两阶段各自可并行，均来自原始论文 §3.1 和 L1 讲义的 word-count 图示。[原始论文 §3.1](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf) [L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)

L1 同时强调这个抽象的代价：程序固定为 `Map → shuffle → Reduce`，应用之间不能直接交互或保存跨阶段状态（只能借 intermediate output），并且它是 batch 而非 real-time/streaming 模型。[L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)

## Coordinator、worker 与任务状态

原论文中的 **master** 选择 idle worker 来分派 `M` 个 Map 和 `R` 个 Reduce task；它保存每个任务的 `idle`、`in-progress` 或 `completed` 状态，并把完成的 Map 输出位置传递给 Reduce worker。[原始论文 §3.1–3.2](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

MIT 的 2026 Lab 1 使用 **coordinator** 这个名字而非论文的 **master**。在该 Lab 的具体环境中，有一个 coordinator 和一个或多个并行 worker；worker 循环请求任务、读输入、执行、写输出、再请求下一任务，双方用 RPC 通信。[Lab 1：Introduction 与 Your Job](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html) L1 讲义也把 coordinator 概括为“跟踪任务状态、向 worker 分发任务”。[L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)

教学时可以说明这两个术语承担相似的调度角色；但要区分原论文描述的 Google 集群实现，和本课程 Lab 在单机、共享文件系统上的教学化实现。

## 失败、重试与完成语义

### Lab 1 明确要求的边界

- 若一个 worker 未在合理时间内完成任务，coordinator 必须把同一任务交给另一个 worker；2026 Lab 1 指定使用 **10 秒**。官方测试检查 Map/Reduce 的并行性，以及 worker 执行任务时 crash 后的恢复。[Lab 1：Your Job 与测试](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html)
- 官方说明也明确：coordinator 不能可靠地区分 crash、已存活但停滞、仍在工作但太慢的 worker；它能做的是等待一段时间后重新发放任务。因此，**超时是重试决策，不是“已经证明 worker 死亡”**这一结论的直接教学推论。[Lab 1：Hints](https://pdos.csail.mit.edu/6.824/labs/lab-mr.html)

### 原论文解释的生产级语义

- 原论文的 master 定期 ping worker。worker 失联后，正在运行的任务变回 `idle`；该 worker 已完成的 Map 也会重跑，因为 intermediate data 只在失效机器的 local disk。已完成的 Reduce 输出在 global file system 中，因此不必重跑。[原始论文 §3.3](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)
- 仅当用户的 `Map`/`Reduce` 对输入是 deterministic，且任务输出用 atomic commit 发布时，分布式、有失败的执行才等价于无故障的顺序执行。论文的机制是任务先写 private temporary files；Reduce 完成时以 atomic rename 发布最终文件。[原始论文 §3.3](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)
- **straggler** 是慢任务，不等于 crash；论文通过在任务接近完成时启动 backup execution 来缩短尾部时间。另一方面，论文的单一 master 失效时，现有实现会中止整个 job（客户端可自行重试），并非 coordinator 已高可用。[原始论文 §3.3、§3.6](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)

不能把论文的 ping、GFS、master checkpoint 或 backup task 当成 2026 Lab 1 已规定的实现细节；它们用于解释设计动机和更完整的失败语义。

## 学习者易混淆的点

- **Map 不是“每个 key 做一次”。** 它按 input split/record 被调用，负责产生 intermediate pairs；同 key 的聚合发生在 shuffle 之后的 Reduce。
- **Reduce 不是“接收一个 value”。** 它接收一个 intermediate key 和该 key 的全部 values；同一个 Reduce partition 往往包含多个不同 key，因而需要分组/排序。
- **“任务完成”不等于“worker 永远可靠”。** Map 输出位置和落盘位置决定是否必须重算；重发任务还可能让同一逻辑任务有多个 execution。
- **并行并不自动等于线性加速。** L1 指出负载不均衡以及网络、磁盘、CPU、内存都可能成为瓶颈；更多 task than worker 可帮助动态负载均衡。[L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)
- **deterministic 是故障重试的前提之一。** 若函数读外部可变状态、依赖随机数或有外部副作用，重复执行可能不再对应一个清晰的顺序结果；L1 将此归为应用程序作者的责任。[L1 讲义](https://pdos.csail.mit.edu/6.824/notes/l01.txt)

## 适合 L1 的可验证理解

给出两三个输入 split 的 word-count 中间输出，学习者应能独立说明：哪些 Map 可并行、同一个 key 的 values 怎样在 shuffle 后进入一个 Reduce、为什么某 Map 需要重试时要求 deterministic，以及 worker crash、straggler 与 coordinator/master crash 的不同。这个检查只验证模型与故障语义，不要求实现 Lab 1。

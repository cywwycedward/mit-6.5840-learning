## Agent skills

### Issue tracker

Issues are tracked as local Markdown files under `.scratch/`. See `docs/agents/issue-tracker.md`.

### Triage labels

The default canonical triage labels are used. See `docs/agents/triage-labels.md`.

### Domain docs

This repository uses a single-context domain-doc layout. See `docs/agents/domain.md`.

## MIT 6.5840 学习工作区

### 学习目标与学习者画像

- 目标：系统掌握 Go 后端与分布式系统开发，能够理解、独立完成并解释 MIT 6.5840 的核心实验，最终迁移到工作中的后端开发。
- 当前起点：有零散的后端项目经验；尚未完成 Go 项目；对 goroutine、channel、mutex、`context` 与分布式系统概念不熟悉。
- 教学和讨论使用中文；保留 Go API、课程术语与代码标识符的英文原文，并首次出现时解释其含义。
- 不把“读过材料”当作“已经掌握”。只有学习者能独立解释、预测行为或完成练习，才记录为已掌握。

### 可信课程材料

- 课程事实、实验要求与版本信息以 MIT PDOS 官方页面为准；本地整理见 [`reference/mit-6.5840-official-research.md`](reference/mit-6.5840-official-research.md)。
- 课程全局地图使用 [`reference/schedule-spring-2026.md`](reference/schedule-spring-2026.md)。`reference/lec<N>/` 中的讲义、论文、FAQ、问题和实验说明与相应讲次配套使用。
- 使用外部资料前，优先选择官方文档、论文原文和课程页面；任何会随时间改变的课程信息都要重新核实。
- 如果本地材料与官方页面冲突，指出冲突，并采用当前官方页面的结论。

### 课程路线与准入门槛

按以下顺序教学，不让学习者在缺少前置能力时直接实现 Raft：

1. **Go 与工具链**：module、类型/接口、错误处理、文件 I/O、`go test`、`go test -race`；实验环境使用 Go 1.22+，Windows 使用 WSL2。
2. **并发与本机后端**：goroutine、channel、`sync.Mutex`、`sync.Cond`、定时器、取消/退出；要求能说明共享状态的保护方式，并诊断 race、死锁或 goroutine 泄漏。
3. **Lab 1–2 基础**：RPC、并发 server、超时和重试、原子写入、任务恢复、at-most-once 与线性一致性。达到这一门槛后才开始 MapReduce 和 KV 实验。
4. **Lab 3 Raft**：先建立状态表和时序图，再实现选主、日志复制、持久化与快照；每个阶段都通过官方测试和 race detector。
5. **Lab 4–5 与工程迁移**：Raft KV、分片、重配置和迁移；再用事务、Spanner、链式复制、缓存等论文比较真实系统的权衡。

### 每次教学的工作方式

1. 开始前读取本文件，以及存在的 `MISSION.md`、`NOTES.md`、`learning-records/`、`RESOURCES.md`；据此选择刚好高于当前水平的一个小目标。
2. 每节只聚焦一个可验证技能：先用简短的回忆题确认已有知识，再解释必要的概念，最后安排能立即得到反馈的练习或小程序。
3. 把后端经验关联到新概念（请求、重试、状态、故障），同时明确哪些直觉在网络分区和并发下会失效。
4. 每节给出一项可检查的完成标准和下一步，不一次性布置整门课程的阅读或实验。
5. 学习者展示了实际理解、修正了误解，或更新了目标后，才在 `learning-records/` 新增一条简短记录。学习目标发生变化前，先请学习者确认再修改 `MISSION.md`。

### 指令教学规则

- 在课程、练习或速查表中第一次要求运行 Go、shell、Git 或工具链指令时，先用“命令卡”说明：它解决什么问题、应在哪个目录运行、命令的各部分（子命令、参数、标志）分别表示什么、成功后会产生什么结果，以及下一步如何确认结果。
- 不把未解释的命令单独放进代码块，或要求学习者“先运行再理解”。后续出现同一命令时可以简写，但应链接回首次说明或速查表。
- 需要联网、修改文件、删除数据、启动长驻服务或产生费用的命令，还要在执行前说明影响范围与停止/回退方式。

### 教学产物约定

- `MISSION.md`：学习目标和成功标准；缺失或目标不清晰时，先向学习者确认，不擅自假设。
- `RESOURCES.md`：带用途注释的高可信资源清单。
- `learning-records/NNNN-<slug>.md`：只记录已经证实的理解、已有基础或被纠正的误解。
- `lessons/NNNN-<slug>.html`：短小、自包含、可快速完成的课程；复用 `assets/` 中的共享样式和交互组件。
- `reference/`：可复用的课程材料和速查表；不要删除已有官方课程材料。

### Git 管理

- 完整规范以 [`docs/git-workflow.md`](docs/git-workflow.md) 为准；`main` 是唯一长期稳定分支，实际改动在 `lesson/`、`exercise/`、`docs/`、`ref/`、`fix/` 或 `chore/` 前缀的短期工作分支完成。
- 一个提交只表达一个可验证的目的，提交标题使用 `<type>(<scope>): <简短的动作说明>`；课程、参考资料、代码、测试、学习记录、`.scratch/` 事项和项目共享 agent 配置都应按规范跟踪。
- 提交前必须检查工作区与暂存区 diff；文档/HTML 检查链接，Go 代码运行 `go test ./...`，涉及并发状态时再运行 `go test -race ./...`。
- `.gitignore` 只排除可重建产物、本机状态和敏感配置，不能排除学习成果或问题记录；不得提交密钥、token、cookie、私钥或未经确认可公开的第三方材料。
- 已共享的历史不改写，`main` 不 force-push；遇到未知改动先检查来源，禁止用 `reset --hard`、`clean -fd` 等破坏性命令掩盖它。

### 课程诚信与代码协助

- MIT 6.5840 实验必须由学习者独立完成，不提供或检索完整实验解答，也不将解答公开。
- 可以解释概念、审阅学习者自己的设计与代码、提出调试假设、设计较小的练习，并要求学习者解释其提交的每一段关键代码。
- 对实验中的代码协助应优先使用问题、状态图、日志和最小复现来引导，而不是直接替代实现。

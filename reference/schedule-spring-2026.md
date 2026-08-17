# MIT 6.5840（原链接为 6.824）2026 春季课表速览

来源：MIT PDOS 官方 [6.5840 Schedule: Spring 2026](https://pdos.csail.mit.edu/6.824/schedule.html)。访问日期：2026-08-17。

## 页面身份与使用方式

- 虽然 URL 路径仍为 `/6.824/`，页面标题、课程主页链接和日期单元格均显示这是 **MIT 6.5840，2026 年春季**的课表；上课时间为周二、周四 13:00–14:30，地点 54-100。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html)
- 课表明确标为暂定；尚未到来的讲义和论文问题可能是往年副本，之后会更新。以下是该页面当前列出的内容，而不是对未来安排的保证。[官方说明](https://pdos.csail.mit.edu/6.824/schedule.html)
- 所有实验截止时间均为当天 23:59。[Lab 1 截止说明](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-13)

## 讲次与课前阅读

下表的“阅读”是课表中标为 `Preparation` 的材料；每项还可由同一页面进入讲义、FAQ 和论文问题。

| 日期 | 讲次与主题 | 课前阅读 |
|---|---|---|
| 2/3 | L1：Introduction | MapReduce（2004）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-3) |
| 2/5 | L2：RPC and Threads | Online Go tutorial（附 FAQ）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-5) |
| 2/10 | L3：GFS | GFS（2003）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-10) |
| 2/12 | L4：Paxos | Paxos[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-12) |
| 2/19 | L5（Russ Cox 客座）：Go patterns | *The Go Programming Language and Environment*[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-19) |
| 2/26 | L6：Fault Tolerance: Raft (1) | Raft extended（2014），至第 5 节末[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-26) |
| 3/3 | L7：Fault Tolerance: Raft (2) | Raft extended，第 7 节至结尾（不读第 6 节）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-3) |
| 3/5 | L8：Consistency and Linearizability | *Linearizability*，仅至 §3.1[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-5) |
| 3/10 | L9：Zookeeper | ZooKeeper（2010）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-10) |
| 3/12 | L10：Lab 3A+B Q&A | 无指定论文；有课程问题[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-12) |
| 3/17 | L11：Distributed Transactions | 6.033 Chapter 9 的 §9.1.5、§9.1.6、§9.5.2、§9.5.3、§9.6.3[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-17) |
| 3/31 | L12：Spanner | Spanner（2012）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-31) |
| 4/2 | L13：Chain Replication | Chain Replication（2004）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-2) |
| 4/7 | L14：Optimistic Concurrency Control | FaRM（2015）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-7) |
| 4/9 | L15（Upamanyu Sharma）：Verification of distributed systems | IronFleet（2015）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-9) |
| 4/16 | L16：Cache Consistency: Memcached at Facebook | *Memcached at Facebook*（2013）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-16) |
| 4/21 | L17（Marc Brooker，Zoom）：AWS Lambda | *On-demand Container Loading*（2023）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-21) |
| 4/23 | L18：Ray | Ray（2021）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-23) |
| 4/28 | L19：Fork Consistency, SUNDR | SUNDR（2004），读至 §3.3.2 末[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-28) |
| 4/30 | L20：Peer-to-peer: Bitcoin | Bitcoin（2008）及协议概要[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-30) |
| 5/7 | L21（Derek Leung）：Byzantine Fault Tolerance | Practical BFT（1999）[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-7) |
| 5/12 | L22：Project demos；最后上课日 | 无指定阅读[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-12) |

## 实验和期末项目

| 发布日 | 模块 | 截止日 |
|---|---|---|
| 2/3 | Lab 1：MapReduce | 2/13[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-3) [截止](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-13) |
| 2/10 | Lab 2：Key/Value server | 2/20[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-10) [截止](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-20) |
| 2/17 | Lab 3：Raft | 3A：2/27；3B：3/6；3C：3/13；3D：4/3[发布](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-17) [3A](https://pdos.csail.mit.edu/6.824/schedule.html#2026-2-27) [3B](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-6) [3C](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-13) [3D](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-3) |
| 3/10 | Lab 4：KV Raft | 4A：4/10；4B+C：4/17[发布](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-10) [4A](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-10) [4B+C](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-17) |
| 4/7 | Lab 5：Sharded KV | 5A：4/29；5B+C+D：5/8[发布](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-7) [5A](https://pdos.csail.mit.edu/6.824/schedule.html#2026-4-29) [5B+C+D](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-8) |
| 3/3 | Final Project | proposal：3/20（仅做项目者）；报告和代码：5/8；演示：5/12[发布](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-3) [提案](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-20) [报告/代码](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-8) [演示](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-12) |

## 考试与日历节点

- 期中考试：3/19 13:00–14:30，Walker 50-340；开卷，可用笔记和笔记本电脑；范围 L1–L11、Lab 1、Lab 2、Lab 3A–C。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-3-19)
- 期末考试：5/15 09:00–11:00，34-101；开卷，可用笔记和笔记本电脑；范围 L12–L21、Lab 3D、Lab 4A–C。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html#2026-5-15)
- 2/2 开课；2/16 总统日放假；2/24 雪天停课；3/23–3/27 春假；4/14、5/5 为 Hacking day、无讲课；4/20 Patriots' Day 放假。3/6 为加课截止日，4/21 为退课截止日。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html)

## 学习主线（归纳）

课表按 **分布式计算入门（MapReduce、RPC、GFS）→ 一致性与容错（Paxos、Raft、线性一致性、ZooKeeper）→ 事务和复制（2PC、Spanner、Chain Replication、FaRM）→ 工程与新主题（验证、缓存、Serverless、Ray、安全与 P2P、BFT）** 展开；对应实践从 MapReduce/KV，进展到 Raft、KV-Raft 与分片 KV，并可完成期末项目。此段为根据上述官方讲次和实验顺序作出的归纳。[官方课表](https://pdos.csail.mit.edu/6.824/schedule.html)

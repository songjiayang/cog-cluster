# cog-cluster

熟悉 SD 服务部署的朋友，可能对 [cog](https://github.com/replicate/cog) 并不陌生，它能够统一机器学习单推理节点的生命周期管理，包括镜像构建、服务构建、服务部署等。

但因为它集群能力的缺陷（不支持中心化的任务队列），导致基于它构建类似 [replicate](https://replicate.com/) 的服务存在不小的挑战 ，故 cog-cluster 诞生，意在尽量简化 cog 集群的搭建。

## 项目架构

![cog-cluster](https://github.com/user-attachments/assets/2e0c39c1-c934-4955-8efb-d4c502a66051)

架构说明：

- cog-api: 负责异步推理任务的创建，查询和耗时评估。
- cog-agent: 作为 cog-server 的 sidecar 进行部署，属于 cog-worker 的一部分，负责从 redis 拉取同 worker 类型的任务。
- redis： 负责待处理异步推理任务的中心式存储。

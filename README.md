# cog-cluster

熟悉 SD 服务部署的朋友，可能对 [cog](https://github.com/replicate/cog) 并不陌生，它能够统一机器学习单推理节点的生命周期管理，包括镜像构建、服务构建、服务部署等。

但因为它集群能力的缺陷（不支持中心化的任务队列），导致基于它构建类似 [replicate](https://replicate.com/) 的服务存在不小的挑战 ，故 cog-cluster 诞生，意在尽量简化 cog 集群的搭建。

## 项目架构

![cog-cluster](https://github.com/user-attachments/assets/2e0c39c1-c934-4955-8efb-d4c502a66051)

架构说明：

- cog-api: 负责异步推理任务的创建，查询和耗时评估。
- cog-agent: 作为 cog-server 的 sidecar 进行部署，属于 cog-worker 的一部分，负责从 redis 拉取同 worker 类型的任务。
- redis： 负责待处理异步推理任务的中心式存储。

## 使用

1. 制作 rsnet cog-server 镜像，可以参考 [cog#getting-started](https://cog.run/getting-started/)。
2. 使用命令 `make up` 启动 examples 样例所有容器，主要包括 redis, cog-api, cog-agent, cog-server。
3. 发起 rsnet 推理请求到 cog-api:

```bash
curl http://localhost:8000/predictions -X POST \
    -H 'Content-Type: application/json' \
    -d '{"input": {"image": "https://gist.githubusercontent.com/bfirsh/3c2115692682ae260932a67d93fd94a8/raw/56b19f53f7643bb6c0b822c410c366c3a6244de2/mystery.jpg"}}'

=>

{"task_id":"8f36ddc5-132e-4e2e-8f3b-1c56eb366659"}
```

该请求会返回 task id，后续我们会用到该id进行任务结果查询

4. 等待大概1s左右，通过  cog-api 获取推理结果：

```
curl http://localhost:8000/predictions/8f36ddc5-132e-4e2e-8f3b-1c56eb366659

=> 

{"id":"8f36ddc5-132e-4e2e-8f3b-1c56eb366659","output":[["n02123159","tiger_cat",0.4898366928100586],["n02123045","tabby",0.23457567393779755],["n02124075","Egyptian_cat",0.09744952619075775]],"status":"succeeded","started_at":"2024-12-02T01:25:22.632050+00:00","completed_at":"2024-12-02T01:25:23.109751+00:00"}
```

可以看到根据 task id 我们可以获取最终推理的结果，结果还包含任务推理开始和完成时间。

## 致谢

感谢以下开源项目，没有它们就不会有本项目的存在。

- [cog](https://github.com/replicate/cog)
- [cog-comfyui](https://github.com/fofr/cog-comfyui)
- [asynq](https://github.com/hibiken/asynq)
  

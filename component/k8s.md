# K8s

### 架构体系

```
k8s的master节点实现了对集群的管理，主要有四个组件：api-server、controller-mananger、kube-scheduler、etcd
```

- api-server：提供restful接口，实现整个k8s集群通信
- controller_manager：集群的管理控制中心，对集群的资源进行管理
- kube-scheduler：实现调度算法和策略，为pod选择合适的节点
- etcd：非关系型数据库，集群内资源的存储

### 创建一个pod的流程

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210310100538.png)

- api-server处理用户请求，将pod信息存储到etcd中
- kube-scheduler预选和优选为pod选择最优的节点
- 节点的kubelet从节点点中获取pod清单，下载镜像启动容器

### 参考

- 从零开始入门 K8s：详解 K8s 核心概念：https://www.infoq.cn/article/knmavdo3jxs3qpkqtzbw
- 官方文档：https://kubernetes.io/zh/docs/home/
- https://www.jianshu.com/p/2de643caefc1
- 学习路线：https://www.infoq.cn/article/9dtx*1i1z8hsxkdrpmhk
- minikube：https://github.com/kubernetes/minikube
- minikube docs：https://minikube.sigs.k8s.io/docs/start/
- kubectl安装：https://github.com/caicloud/kube-ladder/blob/master/tutorials/lab1-installation.md
# collector

用于收集离线安装包的工具。



## 安装

## 使用说明

程序读取预先定义的配置文件（collector/config/v1.0.0.yaml），根据命令行参数选择性的下载离线包。

```yaml
# 定义application
- name: argocd
  version: v2.7.6

  # 定义etcd应用的镜像
  imageGroups:
  - images:
    - source: 172.18.60.199:5000/dexidp/dex:v2.36.0
    - source: 172.18.60.199:5000/argoproj/argocd:v2.7.6
    - source: 172.18.60.199:5000/redis/7.0.11-alpine

  # 定义etcd应用的二进制文件
  binaries:
  - url: https://github.com/argoproj/argo-cd/releases/download/v2.7.6/argocd-linux-amd64
    sha256sum: 4f2f56548ad042cc168108263fd0e92bdccc0f61957e4df32dabf3bc53ef4b87
    arch: amd64
    filename: argocd-linux-amd64
  - url: https://github.com/argoproj/argo-cd/releases/download/v2.7.6/argocd-linux-arm64
    sha256sum: b4fa9bac5dfe7e3677ab2ec2d369443a158d8e3265984964cd568a6057f8906b
    arch: arm64
    filename: argocd-linux-arm64

- name: keepalived
  imageGroups:
  - images:
    - source: keepalived:latest
```

默认收集所有应用，所有类型的资源：

```
collector --config collector/config/v1.0.0.yaml --output-path /srv/collector/2023-11-23
```

收集指定应用的资源：

```
collector etcd --config collector/config/v1.0.0.yaml --output-path /srv/collector/2023-11-23
```

收集指定类型的资源：

```
collector --type image --config collector/config/v1.0.0.yaml --output-path /srv/collector/2023-11-23
```

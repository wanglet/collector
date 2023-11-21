# collector
用于收集离线安装的工具




## 支持的类型

### image

```
imageGroups:
- name: cluster-proportional-autoscaler
  output:
    type: docker
  images:
  - source: 172.18.60.199:5000/cpa/cluster-proportional-autoscaler-arm64:1.8.5
  - source: 172.18.60.199:5000/cpa/cluster-proportional-autoscaler-amd64:1.8.5
- name: httpd
  images:
  - source: httpd:2.4
```
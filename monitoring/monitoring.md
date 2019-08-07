# Prometheus监控系统

## 安装Prometheus Server

执行下面命令安装prometheus:

```
docker run --name=prometheus -d \
-p 9090:9090 \
prom/prometheus  \
--web.enable-lifecycle
```
如下图所示：
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/index.jpeg'>
</p>
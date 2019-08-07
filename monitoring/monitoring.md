# Prometheus监控系统

## 安装Prometheus Server

执行下面命令安装prometheus:

```
docker run --name=prometheus --restart=always  -d -p 9090:9090 prom/prometheus  --web.enable-lifecycle --config.file=/etc/prometheus/prometheus.yml
```
说明：
- 启动时加上 `--web.enable-lifecycle` 启用远程热加载配置文件
- 调用指令是：curl -X POST 127.0.0.1:9090/-/reload

使用如下命令查看:
```
docker container ls     // 执行此命令

// 看到如下信息表示安装成功
prometheus docker container ls
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                               NAMES
37cb8259141d        prom/prometheus         "/bin/prometheus --w…"   6 minutes ago       Up 6 minutes        0.0.0.0:9090->9090/tcp              prometheus
```

访问： 127.0.0.1:9090 看到如下界面表示安装成功



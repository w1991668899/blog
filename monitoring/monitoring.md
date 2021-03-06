# Prometheus监控系统

## 安装Prometheus Server  [官网](https://prometheus.io/download/)

### 部署 prometheus:

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
docker container ls
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                               NAMES
37cb8259141d        prom/prometheus         "/bin/prometheus --w…"   6 minutes ago       Up 6 minutes        0.0.0.0:9090->9090/tcp              prometheus
```

访问： 106.15.95.51:9090 看到如下界面表示安装成功:
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/343242342.png'>
</p>

访问：106.15.95.51:9090/metrics 看到如下界面表示prometheus在抓取自己的 `/metrics`接口新信息：
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/22222.png'>
</p>

## metrics客户端

### 部署

````
mkdir -p /home/wt/promethues/client/golang/src
cd !$
export GOPATH=/home/chenqionghe/promethues/client/golang/
#克隆项目
git clone https://github.com/prometheus/client_golang.git
#编译
cd $GOPATH/src/client_golang/examples/random
go build -o random main.go
````

运行程序：

```
./random -listen-address=:8080 &
./random -listen-address=:8081 &
./random -listen-address=:8082 &
```


## node-exporter [官网](https://github.com/prometheus/node_exporter)

### 部署流程

1. 使用如下命令步骤安装
```
docker run -d --name=node-exporter --restart=always -p 9100:9100 prom/node-exporter
```

使用 `docker container ls` 命令查看如下则表示安装成功:
```

docker container ls   //执行命令
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                               NAMES
1b6f877d1858        prom/node-exporter      "/bin/node_exporter"     8 seconds ago       Up 7 seconds        0.0.0.0:9100->9100/tcp
```

2. 更改配置信息：

1. 执行 `touch prometheus.yml` 命令，生成  prometheus.yml 文件
2. 执行 `vim ./prometheus.yml` 编辑文件，将下面信息复制进文件中，注意更新需要监控的服务器IP信息：

```
global:
  scrape_interval:     15s # 默认抓取间隔, 15秒向目标抓取一次数据。
  external_labels:
    monitor: 'codelab-monitor'
rule_files:
  #- 'prometheus.rules'
# 这里表示抓取对象的配置
scrape_configs:
  #这个配置是表示在这个配置内的时间序例，每一条都会自动添加上这个{job_name:"prometheus"}的标签  - job_name: 'prometheus'
  - job_name: 'prometheus'
    scrape_interval: 5s # 重写了全局抓取间隔时间，由15秒重写成5秒
    static_configs:
      - targets: ['106.15.95.51:9090']
      - targets: ['106.15.95.51:8080', '106.15.95.51:8081','106.15.95.51:8082']
        labels:
          group: 'client-golang'
      - targets: ['106.15.95.51:9100']
        labels:
          group: 'client-node-exporter'
```

3. 执行 `docker cp ./prometheus.yml 37cb8259141d:/etc/prometheus/prometheus.yml` 命令替换配置文件

注：
- ./prometheus.yml 指刚才创建的服务器地址
- 37cb8259141d 指容器 prometheus 容器的ID可以使用， `docker container ls` 命令查看
- /etc/prometheus/prometheus.yml  指容器内部文件地址

4. 使用 `docker container restart 37cb8259141d` 重启容器

5. 访问 `106.15.95.51:9090/targets` 接口，如出现如下图所示表示成功：

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/555552343.png'>
</p>

## pushgateway

### 部署

```
docker run -d -p 9091:9091 --restart=always  --name pushgateway prom/pushgateway
```

访问：http://106.15.95.51:9091 看到如下界面即可:

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/44444234324.png'>
</p>

## Grafana可视化工具  

### 部署流程

1.  docker安装

```
docker run -d \
  -p 3000:3000 \
  -e INFLUXDB_HOST=localhost \
  -e INFLUXDB_PORT=8086 \
  -e INFLUXDB_NAME=cadvisor \
  -e INFLUXDB_USER=root -e INFLUXDB_PASS=root \
  --link influxdb:influxdb \
  --name grafana \
grafana/grafana
```

访问 127.0.0.1:3000 

默认账号： admin
默认密码： admin

2. 登录后如下：

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/555552343.png'>
</p>

## 添加source 如下图

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/777esfdfd.png'>
</p>

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/883432df.png'>
</p>

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/999343433.png'>
</p>

## 结果如下图

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/12345dfsf.png'>
</p>


















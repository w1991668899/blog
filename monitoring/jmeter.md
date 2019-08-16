# 分布式压测与监控系统搭建流程

## 概述

介绍使用 JMeter+InfluxDB+Cadvisor+Grafana 搭建分布式可视化压测实时监控系统

## 分布式一主多从

<p align='center'>
<img src='https://jmeter.apache.org/images/screenshots/distributed-jmeter.svg'>
</p>

## 压测流程

<p align='center'>
<img src='https://jmeter.apache.org/images/screenshots/distributed-names.svg'>
</p>


## 分布式环境与压力服务器要求

### 硬件

- 控制机： 独立linux服务器, 负责测试脚本发布与测试数据汇总
- 压力机： linux服务器， 压力机的带宽要比服务器带宽高，负责实际的压力测试。执行完成后，压力机会把结果回传给控制机，控制机会收集所有压力机的信息并汇总


### 软件

- 关闭系统上的防火墙或打开正确的端口
- 所有压力机IP都在同一个子网上
- 确保在所有系统上使用相同版本的JMeter和Java，混合版本无法正常工作


- Grafana 数据的可视化展示， 需要部署在独立的服务器

## 部署jmeter服务端/客户端  [官网地址](https://jmeter.apache.org/download_jmeter.cgi)

### jmeter 服务端部署

```
docker run --detach --publish 1099:1099 -d --name=jmeter_server --net=host  egaillardon/jmeter -Jserver.rmi.ssl.disable=true -Djava.rmi.server.hostname=192.168.3.14 -Jserver.rmi.localport=1099 -Dserver_port=1099 --server
```

`192.168.3.14` 为当前压力机IP

多台压力机重复上面操作

### jmeter 客户端部署

```
docker run --detach --restart=always -d --name=jmeter_client --net=host  --volume `pwd`:/jmeter egaillardon/jmeter -Jserver.rmi.ssl.disable=true --nongui --testfile test.jmx --remotestart 192.168.3.14:1099,192.168.3.15:1099 --logfile result.jtl
```

- `--remotestart 192.168.3.14:1099,192.168.3.15:1099`   指定远程压力机
- `--testfile test.jmx`  指定测试样本

## influxDB  [官网](https://www.influxdata.com/)

时序数据库，负责压测数据的存储，控制机和压力机之间的时间需要保证同步

### 部署

```
docker run -d --restart=always -p 8083:8083 -p 8086:8086 --expose 8090 --expose 8099 --name influxdb tutum/influxdb
```

部署完成后访问 `http://106.15.95.51:8083`  访问可视化界面, 换成自己的IP, 如下图所示：

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/bbbb2332.png'>
</p>

命令行操作：

```
CREATE DATABASE "cadvisor"
```

```
CREATE USER "cadvisor" WITH PASSWORD 'cadvisor'
grant all privileges on "cadvisor" to "cadvisor"
```

## cadvisor  [官网](https://github.com/google/cadvisor)

Google用来监测单节点的资源信息的监控工具

### 部署

```
docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  -p 8080:8080 \
  --detach=true --link influxdb:influxdb \
  --name=cadvisor \
  google/cadvisor:latest \
  -storage_driver=influxdb \
  -storage_driver_db=cadvisor \
  -storage_driver_host=influxdb:8086
```

- `-p 8080:8080` 指定默认端口
- `--link infludb:influxdb` 指定数据类型
- `--name=cadvisor` 指定数据库名称, 用来收集数据


访问（换成自己的IP）： `http://106.15.95.51:8080/containers/`  出现如下图所示：

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/aa213213.png'>
</p>

## granfana  

### 部署

```
docker run -d \
  -p 3000:3000 \
  -e INFLUXDB_HOST=localhost \        // 安装influxdb的host
  -e INFLUXDB_PORT=8086 \
  -e INFLUXDB_NAME=cadvisor \
  -e INFLUXDB_USER=root -e INFLUXDB_PASS=root \
  --link influxdb:influxdb \
  --name grafana --restart=always -d \
grafana/grafana
```

- `-e INFLUXDB_NAME=cadvisor` 指定数据库,收集granfana所在服务器信息

## 分布式环境配置








# 分布式压测与监控系统搭建流程

## 概述
介绍使用 JMeter+InfluxDB+Grafana 搭建分布式可视化压测实时监控系统

## 架构简介
<p align='center'>
<img src='https://jmeter.apache.org/images/screenshots/distributed-jmeter.svg'>
</p>

## 分布式环境与压力服务器要求

- 控制机 可在本机电脑或独立服务器, 负责测试脚本发布
- 压力机 linux服务器， 压力机的带宽要比服务器带宽高，负责实际的压力测试。执行完成后，压力机会把结果回传给控制机，控制机会收集所有压力机的信息并汇总
- Jmeter Apache组织开发的基于Java的压力测试工具
- InfluxDB 时序数据库， 负责压测数据的存储。控制机和压力机之间的时间需要保证同步
- cAdvisor Google用来监测单节点的资源信息的监控工具
- Grafana 数据的可视化展示， 需要部署在独立的服务器

## 部署jmeter  [官网地址](https://jmeter.apache.org/download_jmeter.cgi)
多台压力机重复上面操作

## 部署 influxDB
使用docker安装:

```
docker run -d --restart=always -p 8083:8083 -p 8086:8086 --expose 8090 --expose 8099 --name influxdb tutum/influxdb
```
http://106.15.95.51:8083/  访问可视化界面, 换成自己的IP

分别执行以下命令：

```
CREATE DATABASE "cadvisor"
```

```
CREATE USER "cadvisor" WITH PASSWORD 'cadvisor'
grant all privileges on "cadvisor" to "cadvisor"
```

访问（换成自己的IP）： `http://106.15.95.51:8080/containers/`  出现如下图所示：
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/monitoring/aa213213.png'>
</p>

## 安装cadvisor  [官网](https://github.com/google/cadvisor)

使用docker 安装
```
docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  -p 8080:8080 \
  --detach=true --link influxsrv:influxdb \
  --name=cadvisor \
  google/cadvisor:latest \
  -storage_driver=influxdb \
  -storage_driver_db=cadvisor \
  -storage_driver_host=influxdb:8086
```

## 安装 granfana  

```
docker run -d \
  -p 3000:3000 \
  -e INFLUXDB_HOST=localhost \        // 安装influxdb的host
  -e INFLUXDB_PORT=8086 \
  -e INFLUXDB_NAME=cadvisor \
  -e INFLUXDB_USER=root -e INFLUXDB_PASS=root \
  --link influxdb:influxdb \
  --name grafana \
grafana/grafana
```

## 分布式环境配置

```
docker run --detach --publish 11098:11098 --restart=always -d --name=jmeter_server_001  egaillardon/jmeter -Jserver.rmi.ssl.disable=true -Djava.rmi.server.hostname=192.168.3.14 -Jserver.rmi.localport=11098 -Dserver_port=11098 --server
```

```
docker run --detach --restart=always -d --name=jmeter_client  --volume `pwd`:/jmeter egaillardon/jmeter -Jserver.rmi.ssl.disable=true --nongui --testfile test.jmx --remotestart 192.168.3.14:11098,192.168.3.14:21098 --logfile result.jtl
```






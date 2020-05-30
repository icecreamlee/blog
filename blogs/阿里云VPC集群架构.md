# 阿里云VPC集群架构



## 1. VPC



### 1.1. 简介



![img](https://static.ilibing.com/images/blogs/ad7c026864935eedb7dbf3006e166bd3.png)



专有网络（VPC）是您自己独有的云上私有网络。您可以完全掌控自己的专有网络，例如选择IP地址范围、配置路由表和网关等，您可以在自己定义的专有网络中使用阿里云资源如云服务器、云数据库RDS版和负载均衡等。每个VPC都由一个私网网段、一个路由器和至少一个交换机组成。

### 1.2. 创建

创建专有网络（VPC），创建后购买 ECS 时网络选择创建的 VPC 和 其下面的交换机：



![img](https://static.ilibing.com/images/blogs/4b3d40472cf8c2f6ed8724d65a51e67b.png)





## 2. NAT



### 2.1. 简介



![img](https://static.ilibing.com/images/blogs/2bab21c399ac9f75be59e27543562d30.png)



NAT网关（NAT Gateway）是一款企业级的VPC公网网关，提供NAT代理（SNAT和DNAT）、高达10Gbps级别转发能力以及跨可用区的容灾能力。

NAT网关作为一个网关设备，需要绑定公网IP才能正常工作。创建NAT网关后，您可以为NAT网关绑定弹性公网IP（EIP）。

NAT 可以使 VPC 下的多台服务器使用同一 IP 访问公网。

功能：支持多台 VPC ECS 访问公网（SNAT）和用户从公网访问 VPC ECS（DNAT）。

优势：NAT 网关和 EIP 的核心区别是 NAT 网关可用于多台 VPC ECS 和公网通信，而 EIP 只能用于一台VPC ECS和公网通信。

NAT网关提供SNAT和DNAT功能：

**SNAT**

NAT网关提供SNAT功能，为VPC内无公网IP的ECS实例提供访问互联网的代理服务。

此外，NAT网关的SNAT功能还可以作为一个简易防火墙使用。通过SNAT保护后端的服务器，只有后端服务主动和外部终端建立连接后，外部终端才可以访问内部服务器，而未建立连接的外部不可信终端是无法访问后端服务器的。

**DNAT**

NAT网关支持DNAT功能，将NAT网关上的公网IP映射给ECS实例使用，使ECS实例能够提供互联网服务。

DNAT支持端口映射和IP映射。

**宽带共享**

您可以为NAT网关绑定EIP，再将EIP加入到[共享带宽](https://www.aliyun.com/product/cbwp?spm=5176.8142029.388261.320.3836dbcc85aZVp)中。EIP加入到共享带宽后，EIP原本的计费模式失效，不额外收取流量或带宽费，只收取绑定NAT网关的EIP实例费。

### 2.2. 购买



![img](https://static.ilibing.com/images/blogs/868acc974cb222aaec7630fd3f420a75.png)



### 2.3. 设置SNAT

将要使用同一IP访问公网的 ECS/交换机 添加进来：



![img](https://static.ilibing.com/images/blogs/02a3e89ab41c67c0ec8ec9810ab6dfdf.png)





## 3. SLB



### 3.1. 简介



![img](https://static.ilibing.com/images/blogs/ac187969abbebf415b42ef701f3dccd7.png)



负载均衡通过设置虚拟服务地址，将添加的ECS实例虚拟成一个高性能、高可用的应用服务池，并根据转发规则，将来自客户端的请求分发给云服务器池中的ECS实例。

负载均衡默认检查云服务器池中的ECS实例的健康状态，自动隔离异常状态的ECS实例，消除了单台ECS实例的单点故障，提高了应用的整体服务能力。此外，负载均衡还具备抗DDoS攻击的能力，增强了应用服务的防护能力。

负载均衡由以下三个部分组成：

负载均衡实例 （Server Load Balancer instances）

一个负载均衡实例是一个运行的负载均衡服务，用来接收流量并将其分配给后端服务器。要使用负载均衡服务，您必须创建一个负载均衡实例，并至少添加一个监听和两台ECS实例。

监听 （Listeners）

监听用来检查客户端请求并将请求转发给后端服务器。监听也会对后端服务器进行健康检查。

后端服务器（Backend Servers）

一组接收前端请求的ECS实例。您可以单独添加ECS实例到服务器池，也可以通过虚拟服务器组或主备服务器组来批量添加和管理。

### 3.2. 配置

| **前端协议/端口** | **后端协议/端口** | **健康检查** | **调度算法**            | **会话保持** |
| ----------------- | ----------------- | ------------ | ----------------------- | ------------ |
| HTTP:80           | HTTP:80           | √            | 加权轮询(100[主]:1[备]) | 开启         |
| HTTPS:443         | HTTP:80           | √            | 加权轮询(100[主]:1[备]) | 开启         |

注意：

1. 后端协议/端口只使用 HTTP:80，服务器不使用 HTTPS，HTTPS 交给 SLB 去做。

1. 开启会话保持（测试时，可不开启，以保证负载均衡测试）

1. 开启HTTP2.0（可选，仅HTTPS下可配置）

1. 开启 Gzip 压缩（可选）

1. 附加（X-Forwarded-For，X-Forwarded-Proto）HTTP 头字段。

1. 健康检查默认使用服务器 IP 访问，列如: http://http://172.18.15.79/，所以请确保能使用IP地址访问的状态码为 2xx 或者 3xx，或者是更改健康检查的配置。

1. 调度算法使用加权轮询，主备服务器加权比为：100 : 1，因权重不能设置为0，所以只能设为1，尽量不使用备用服务器。（测试时，可设置为1:1，以保证负载均衡测试）



![img](https://static.ilibing.com/images/blogs/f37efa80a38847d49ebfccbca86caf40.png)
# 使用Nginx负载均衡实现基于Cookie的灰度发布



## 1. 起因

灰度发布是指在黑与白之间，能够平滑过渡的一种发布方式。ABtest就是一种灰度发布方式，让一部用户继续用A，一部分用户开始用B，如果用户对B没有什么反对意见，那么逐步扩大范围，把所有用户都迁移到B上面来。

目前我公司的项目是2B的，客户反馈的需求不断，有些是非常急的，所以发布的周期非常短，短暂测试后就发布上线。这也带来的严重的副作用，往往上线后Bug反馈不断，严重时甚至需要回滚发布。如何降低发布的风险，减少影响范围？

使用灰度发布，设置两组服务器，一组提供相对稳定的服务(稳定版)，一组提供最新上线的服务(开发版)。需要新功能的用户使用开发版版网站的服务器，其他的客户使用稳定版的服务器。这样解决催着新需求上线的客户，也同时不影响到其他用户。



## 2. 配置



### 2.1. 概述

新建两组服务器，为a，b两组服务器。a组使用两台服务器，分别为a1，a2，提供稳定版的网站服务。b组使用1台服务器，提成开发版的网站服务。

为了方便将3台服务器的3个站点全部部署在同一台服务器中，启用三个不同的域名或端口以区分。

### 2.2. 新建站点目录

a1/

index.html（内容为：a1）

a2/

index.html（内容为：a2）

b/

index.html（内容为：b）

### 2.3. 配置nginx负载均衡

```
server {
    listen 80; 

    root /var/www/html;

    index index.php index.html index.htm index.nginx-debian.html;

    server_name ab.ilibing.com;
    
    // 默认设置转发的服务组为a
    set $group "a";

    // 当用户携带的cookie中含"SERVERGID=b"时，设置转发的服务组为b
    if ($http_cookie ~* "SVRVERGID=b") {
        set $group "b";
    }   

    location / { 
        proxy_pass http://$group; 

        proxy_set_header Host $host; 

        proxy_set_header X-Real-IP $remote_addr; 

        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }   
}

upstream a {
    server 127.0.0.1:8081;
    server 127.0.0.1:8082;
}

upstream b {
    server 127.0.0.1:8083;
}

server {
    listen 8081;

    root /var/www/html/a1;

    index index.html;

    server_name ab.ilibing.com;

    location / {
        try_files $uri $uri/ =404;
    }
}

server {
    listen 8082;

    root /var/www/html/a2;

    index index.html;

    server_name ab.ilibing.com;

    location / {
        try_files $uri $uri/ =404;
    }
}

server {
    listen 8083;

    root /var/www/html/b;

    index index.html;

    server_name ab.ilibing.com;

    location / {
            try_files $uri $uri/ =404;
    }
}
```

## 2.4. 测试结果

最后访问ab.ilibing.com，终端输入：

curl ab.ilibing.com

查看是否输出为：a1

再次终端输入：

curl ab.ilibing.com

查看是否输出为：a2

再次终端输入：

curl --cookie="SERVERGID=b" ab.ilibing.com

查看是否输出为：b

如测试结果一致，则说明培植成功。
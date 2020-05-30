# Apache configuration & install on windows



## 1. download apache

download url : https://www.apachehaus.com/cgi-bin/download.plx

选择相应的版本下载：

![img](https://static.ilibing.com/images/blogs/ba8c1e67f09986c184816c410555ee75.png)



## 2. install apache



### 2.1. 打开cmd

p.s. 注意需要用管理员权限运行cmd

![img](https://static.ilibing.com/images/blogs/c4464a873b29ff1b0f68b4a028c4d0f1.png)



### 2.2. 运行安装apache为服务命令

将cmd目录切换到apache/bin目录：

```
httpd.exe -k install
```

看到以下提示表示已经安装成功：
```
Installing the 'Apache2.4' service
The 'Apache2.4' service is successfully installed.
Testing httpd.conf....
Errors reported here must be corrected before the service can be started.
httpd.exe: Syntax error on line 39 of D:/Dev/Apache24/conf/httpd.conf: ServerRoot must be a valid directory
```

安装完成后，可以用以下命令启动和关闭apache服务

```
// 启动apahce服务
httpd.exe -k start
// 停止apache服务
httpd.exe -k stop
// 重启apache服务
httpd.exe -k restart
```



## 3. configuration apache



### 3.1. DirectoryIndex添加index.php

```
// 修改前
<IfModule dir_module>
    DirectoryIndex index.html
</IfModule>

// 修改后
<IfModule dir_module>
    DirectoryIndex index.html index.php
</IfModule>
```



### 3.2. 启用访问控制

```
// 修改前
#LoadModule access_compat_module modules/mod_access_compat.so
// 修改后
LoadModule access_compat_module modules/mod_access_compat.so
```



### 3.3. 启用rewrite

```
// 修改前
#LoadModule rewrite_module modules/mod_rewrite.so
// 修改后
LoadModule rewrite_module modules/mod_rewrite.so
```



### 3.4. 配置fcgi

```
# 添加以下配置代码
LoadModule fcgid_module "modules/mod_fcgid.so"

<IfModule fcgid_module>
FcgidInitialEnv PATH "D:/laragon/bin/php/php-7.1.24-nts-Win32-VC14-x64;C:/Windows/system32;C:/Windows;C:/Windows/System32/Wbem;"
FcgidInitialEnv SystemRoot "C:/Windows"
FcgidInitialEnv SystemDrive "C:"
FcgidInitialEnv TEMP "C:/Windows/Temp"
FcgidInitialEnv TMP "C:/Windows/Temp"
FcgidInitialEnv windir "C:/Windows"

# 10 hrs: in case you have long running scripts, increase FcgidIOTimeout 
FcgidIOTimeout 36000
FcgidConnectTimeout 16
FcgidMaxRequestsPerProcess 1000
FcgidMaxProcesses 50
FcgidMaxRequestLen 81310720
# Location php.ini:
FcgidInitialEnv PHPRC "D:/laragon/bin/php/php-7.1.24-nts-Win32-VC14-x64"
FcgidInitialEnv PHP_FCGI_MAX_REQUESTS 1000

<Files ~ "\.php$>"
AddHandler fcgid-script .php
Options +ExecCGI
FcgidWrapper "D:/laragon/bin/php/php-7.1.24-nts-Win32-VC14-x64/php-cgi.exe" .php
</Files>
</IfModule>
```



### 3.5. 配置vhosts

```
<VirtualHost *:80>
    ServerName www.app.cn
	ErrorDocument 500 /503.html
    DocumentRoot "D:/Web/app/public"
	<Directory "D:/Web/app/public">
		Options -Indexes +FollowSymLinks +ExecCGI
		AllowOverride All
		Order allow,deny
		Allow from all
		Require all granted
	</Directory>
</VirtualHost>
```

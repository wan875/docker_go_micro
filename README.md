前言：
该安装包为个人学习使用，开放出来供goer学习交流，单机环境下修改配置文件即可一键安装。

项目技能应用:
熟悉go web框架gin的基本使用
熟悉go rpc,grpc微服务框架go-micro的基本使用
熟悉consul 服务发现工具
熟悉docker docker-compose

目录/镜像说明：
.env : 配置文件
category : GO实现的文章分类服务,安装包中为服务方,供文章发布内容时调用 ;Dockerfile
post : GO实现的文章内容服务,安装包中角色为客户端 ;Dockerfile
mysql : mysql镜像配置文件，数据库目录及测试sql文件
consul : consul Dockerfile 和安装包
docker-compose.yml : 不解释


 install:
一.修改配制文件，可默认，mysql ROOT 密码 为 root
二.进入安装包目录安装:
 docker-compose up -d
二.初始化测试数据:
1.进入容器: docker exec -it mysql bash
2.进入mysql客户端: mysql -uroot -proot
3.创建测试数据库:create database blog;
4.选择数据库:use blog;
5.设置导入时编码:set names utf8;
6.导入测试数据:source /var/test/blog.sql;
7.退出mysql:exit
8.退出容器:exit

四.访问
分类列表 ip:8000/categorys
文章列表 ip:8001/posts
增加文章:POST发送数据 ip:8001/post

ps:微服务由category提供分类服务，增加文章时调用服务判断分类是否存在.

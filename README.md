---
title: 老男孩秒杀项目ETCD的搭建 
tags: 新建,模板,小书匠
grammar_cjkRuby: true
---

### SecKill
> 基于老男孩的go语言教学班,beego框架换成了gin
> 这是基于Go语言的一个秒杀系统，这个系统分三层，接入层、逻辑层、管理层。


##### 秒杀接入层
1. 从Etcd中加载秒杀活动数据到内存当中。
2. 监听Etcd中的数据变化，实时加载数据到内存中。
3. 从Redis中加载黑名单数据到内存当中。
4. 设置白名单。
5. 对用户请求进行黑名单限制。
6. 对用户请求进行流量限制、秒级限制、分级限制。
7. 将用户数据进行签名校验、检验参数的合法性。
8. 接收逻辑层的结果实时返回给用户。


##### 秒杀逻辑层
1. 从Etcd中加载秒杀活动数据到内存当中。
2. 监听Etcd中的数据变化，实时加载数据到内存中。
3. 处理Redis队列中的请求。
4. 限制用户对商品的购买次数。
5. 对商品的抢购频次进行限制。
5. 对商品的抢购概率进行限制。
6. 对合法的请求给予生成抢购资格Token令牌。

##### 秒杀管理层
1. 添加商品数据。
2. 添加抢购活动数据。
3. 将数据同步到Etcd。
4. 将数据同步到数据库。




起初直接安装就可以使用了
但由于V2  V3版本不同 存储的内容不通
以及ETCDCTL的命令不熟和 curl的bug问题,做了很多无用的尝试,都没找到 项目中存储的key
后在同事的帮助下采用docker来跑etcd(据说etcd本机会删不干净数据)
然后就找了很久的镜像,默认的镜像是2.0版本很久以前的
新地址是quay.io/coreos/etcd:v3.3

docker run -d --name etcd \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=etcd-data:/etcd-data \
    quay.io/coreos/etcd:v3.3 \
    /usr/local/bin/etcd \
    --data-dir=/etcd-data \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379

其中也是遇到了docker里面正常,本机访问不了的问题
加入最后两行改了地址映射就可以了
然后curl的调试
   curl -L http://localhost:2379/v3beta/kv/put \
  -X POST -d '{"key": "Zm9v", "value": "YmFy"}'
  
  curl -L http://localhost:2379/v3beta/kv/range \
  -X POST -d '{"key": "Zm9v"}'
  demo正常 其中3.*版本的  /v3/  参数有些区别
  但发送的value 貌似是字节的形式  所有curl无法使用
  
  用回etcdctl  学习了一些命令 
  export ETCDCTL_API=3
  export ETCD_ENDPOINTS="https://0.0.0.0:2379" #本机就不用指定了
  
  etcdctl  get / --prefix --keys-only  
  查看 /开头的key
  
  docker 隐射下的文件地址
  /var/lib/docker/volumes/etcd-data/_data/member/wal


![enter description here](http://image.wandog.top/xsj/framework.png "framework")



测试接口数据

> 接入层接口
```
//查询秒杀接口信息
http://127.0.0.1:8082/sec/info?product_id=3

//查询秒杀商品列表
http://127.0.0.1:8082/sec/list

//秒杀商品接口
http://127.0.0.1:8082/sec/kill
//参数
product_id: 1
user_id: 1
src: 192.168.199.1
auth_code: userauthcode
time: 1530928164
nance: dsdsdjkdjskdjksdjhuieurierei
```

> 管理层接口
```
//添加商品接口
http://127.0.0.1:8081/product/create
//参数
product_name:梨子
product_total:100
status:1

//商品列表接口接口
http://127.0.0.1:8081/product/list

//秒杀活动列表接口
http://127.0.0.1:8081/activity/list

//添加秒杀活动接口
http://127.0.0.1:8081/activity/create
//参数
activity_name: 梨子大甩卖
product_id:4
start_time:1530928052
end_time:1530989052
total:20
status:1
speed:1
buy_limit:1
buy_rate:0.2
```

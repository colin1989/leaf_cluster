# 服务剥离

## Gate 服务
* ws server
* agent 
  * 跨服务绑定 session
* 消息队列 channel
* 获取服务
* 消息队列发送与接收

## Game 服务
* Agent 管理
* 消息队列 channel
* 获取服务
* 消息队列发送与接收

## 路由
在 Gate 解析消息包，转发至游戏服

## 服务发现
* etcd
* consul

## 消息队列
* nats
* grpc

# 游戏改动
* 登录流程
* 雷神战: 直连跨服服务器
* 最强王国: 应该影响不大
*
* 消息顺序的依赖
* 原 grpc 业务
* 跨服活动
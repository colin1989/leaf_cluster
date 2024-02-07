# Session

在原`leaf`框架中，`Agent`对象是一个代表客户端连接的实体，扮演消息收发和处理的角色。
而`Session`则是代表服务之间连接的实体。

以下有几点重要细节：
* 每个服务之间有且只有一条`Session`。
* `Session`只能作为消息收发的角色。
* `Session`同样实现了Agent方法集。
* `Route(msg interface{}, agent interface{}, data interface{}) error`第三个输入参数为本次`Request`中的`SessionData`。
* `Session`所保存的数据是不可靠的，因为`Route`是使用了协程进行了转发。所以在`Route`时，需要把本次`Session`的数据作为输入参数一同转发。

## 方法
* Bind: 通知网关绑定`Agent`用户ID。
* Kick: 通知网关断开客户端连接，成功后网关通知`Node`一条`protos.Disconnect`消息。
* WriteResponse: 把推送给客户端的消息包装成`protos.Response`推送给网关，然后发给客户端。
* Server(): 本条`Session`所对应的服务器信息。
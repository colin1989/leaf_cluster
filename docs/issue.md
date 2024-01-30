# Issue

## Login

- put Username/password
- get Token

- get Gate address



## Gate

* Gate 管理 Game：

* Gate 管理 Player：agent(client conn)

* Gate 和 Game 通信：
  双边通信都需要一个 key 值，来查找相应对象:
  
  * Gate -> Game : 需要携带一个 UserID
  * Game -> Gate : 需要携带一个 AgentID
  
* **消息包客户端是否需要包装一层？**
  
  1. **消息头** ([8]byte)， 校验与转发
  2. 是否登录过(session / agent userdata )
  3. 如果不包装服务端如何解析这个消息包？
  ```go
  通过 MessageHeader 路由 [8]byte
  
  type C2S struct {
    Name string
  }
  
  type XXX struct{
    Header string
    Body []byte
  }
  ```
  
* 服务端踢人操作流程: game -> (kick) -> gate -> (disconnect) -> game -> (player offline) 

* 外挂包如何处理？



## World (Master)

> 未来拆业务时候的消息中转和广播

服务注册/发现：

Gate -> World

Game -> World (broadcast)-> Gate(s)



## Game
* // 服务启动注册 Gate
* 玩家登录成功 bind(userId)
* Gate -> Game : 需要一个 S2S 的消息包类型
  ```go
  type S2S struct {
    From    string
    To      string
    MsgID   string
    UserID int
    Body []byte 
  }
  ```
  
## 网关动态链接方案
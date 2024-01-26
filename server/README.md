# server cluster

```bash
# session
clientID
*agent
serverID
*serverAgent

# game server


# gate server
定时获取game server列表
> timer: update server list from redis(db)

gate server 启动后：
1. 获取 game server list
2. 遍历 list, 分别连接到 game server, Gate.InitClient(1)
3. 将 game server 的连接保存到 map[serverID]*gameServerConn

客户端连接到 gate server后
1. make session: <clientID,serverID>
2. 获取 server: map[serverID]
3. 转发消息 server.WriteMsg(m)

4. 收到 game server的消息，session.Agent.WriteMsg(m)


```

```go
type (
    gameServer struct{
        ServerID int
        gameGate *gate.Gate
    }

    Session struct{
        ClientID int
        clientAgent *gate.Agent
        ServerID int
    }
)

var (
    mapSvr map[int]*gameServer
    mapSession map[int]*Session
)


```
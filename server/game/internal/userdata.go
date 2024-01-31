package internal

import "github.com/name5566/leaf/gate"

type UserData struct {
	Agent   gate.Agent
	UserID  int
	AgentID int
}

func (ud *UserData) SetAgent(a gate.Agent) {
	ud.Agent = a
}

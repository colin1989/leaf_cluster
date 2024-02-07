package internal

import (
	"github.com/name5566/leaf/agent"
)

type UserData struct {
	Agent   agent.Agent
	UserID  int64
	AgentID int64
}

func (ud *UserData) SetAgent(a agent.Agent) {
	ud.Agent = a
}

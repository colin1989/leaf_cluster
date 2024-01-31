package internal

import (
	"message"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

// GateProcessor
// @Description: 客户端消息处理
type GateProcessor struct {
	processor network.Processor
}

var _ network.Processor = (*GateProcessor)(nil)

func NewGateProcessor(p network.Processor) *GateProcessor {
	return &GateProcessor{
		processor: p,
	}
}

func (g *GateProcessor) Route(msg interface{}, userData interface{}) error {
	log.Infof("GateProcessor Route: %+v", msg)

	return g.processor.Route(msg, userData)
}

func (g *GateProcessor) Unmarshal(data []byte) (interface{}, error) {
	v, err := g.processor.Unmarshal(data)
	if err != nil {
		// 反序列化失败，封装成 C2S 走玩家已登录处理流程
		c2s := &message.C2S_Msg{
			//MsgID: message.GetMsgId(data),
			Body: data,
		}
		return c2s, nil
	}
	return v, err
}

func (g *GateProcessor) Marshal(msg interface{}) ([][]byte, error) {
	return g.processor.Marshal(msg)
}

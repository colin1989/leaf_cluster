package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"message"
)

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
		c2s := &message.C2S_Msg{
			MsgID: message.GetMsgId(data),
			Body:  data,
		}
		return c2s, nil
	}
	return v, err
}

func (g *GateProcessor) Marshal(msg interface{}) ([][]byte, error) {
	return g.processor.Marshal(msg)
}

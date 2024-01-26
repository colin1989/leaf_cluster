package internal

import (
	"github.com/name5566/leaf/network"
	"server/msg"
)

type GateMsgProcessor struct {
	processor network.Processor
}

var _ network.Processor = (*GateMsgProcessor)(nil)

func NewGateMsgProcessor(p network.Processor) *GateMsgProcessor {
	return &GateMsgProcessor{
		processor: p,
	}
}

func (f *GateMsgProcessor) Route(m interface{}, userData interface{}) error {
	return f.processor.Route(m, userData)
}

func (f *GateMsgProcessor) Unmarshal(data []byte) (interface{}, error) {
	v, err := f.processor.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	switch m := v.(type) {
	case *msg.S2S_Msg:
		// 解析body
		body, e := f.processor.Unmarshal(m.Body)
		if e != nil {
			return nil, e
		}
		return body, nil
	default:
		return m, nil
	}
}

func (f *GateMsgProcessor) Marshal(v interface{}) ([][]byte, error) {
	body, err := f.processor.Marshal(v)
	if err != nil {
		return nil, err
	}

	// 包装为S2S_Msg消息
	m := &msg.S2S_Msg{
		MsgID: msg.GetMsgId(v),
		Body:  body[0],
	}
	return f.processor.Marshal(m)
}

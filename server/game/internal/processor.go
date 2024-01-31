package internal

import (
	"message"

	"github.com/name5566/leaf/network"
)

type GameProcessor struct {
	processor network.Processor
}

var _ network.Processor = (*GameProcessor)(nil)

func NewGameMsgProcessor(p network.Processor) *GameProcessor {
	return &GameProcessor{
		processor: p,
	}
}

func (f *GameProcessor) Route(m interface{}, userData interface{}) error {
	return f.processor.Route(m, userData)
}

func (f *GameProcessor) Unmarshal(data []byte) (interface{}, error) {
	v, err := f.processor.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	switch m := v.(type) {
	case *message.Gate_Forward:
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

func (f *GameProcessor) Marshal(v interface{}) ([][]byte, error) {
	//body, err := f.processor.Marshal(v)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 包装为S2S_Msg消息
	//m := &message.Gate_Forward{
	//	MsgID: message.GetMsgId(v),
	//	Body:  body[0],
	//}
	//return f.processor.Marshal(m)

	return f.processor.Marshal(v)
}

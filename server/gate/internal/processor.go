package internal

import (
	"fmt"
	"github.com/name5566/leaf/network"
	"reflect"
	"server/msg"
)

type ForwardProcessor struct {
	processor network.Processor
	mapMsg    map[string]reflect.Type
}

var _ network.Processor = (*ForwardProcessor)(nil)

func NewForwardProcessor(p network.Processor) *ForwardProcessor {
	return &ForwardProcessor{
		processor: p,
	}
}

func (f *ForwardProcessor) Route(m interface{}, userData interface{}) error {
	fmt.Printf("ForwardProcessor Route: %+v", m)

	if _, ok := m.(*msg.S2S_Msg); ok {
		return f.processor.Route(m, userData)
	}

	if _, ok := m.(*msg.C2S_Msg); ok {
		return f.processor.Route(m, userData)
	}

	return f.Forward(m, userData)
}

func (f *ForwardProcessor) Unmarshal(data []byte) (interface{}, error) {
	v, err := f.processor.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	switch m := v.(type) {
	case *msg.S2S_Msg, *msg.C2S_Msg:
		return m, nil
	default:
		// 需要转发的消息
		return data, nil
	}
}

func (f *ForwardProcessor) Marshal(msg interface{}) ([][]byte, error) {
	return f.processor.Marshal(msg)
}

func (f *ForwardProcessor) Forward(m interface{}, userData interface{}) error {
	//TODO implement me
	msgID := msg.GetMsgId(m)
	if _, ok := f.mapMsg[msgID]; !ok {
		f.mapMsg[msgID] = reflect.TypeOf(m)
	}

	b, ok := m.([]byte)
	if ok {
		c2s := &msg.C2S_Msg{
			MsgID: msg.GetMsgId(m),
			Body:  b,
		}
		return f.processor.Route(c2s, userData)
	}

	return f.processor.Route(m, userData)
}

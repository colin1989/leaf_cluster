package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"message"
	"reflect"
)

type ForwardProcessor struct {
	processor network.Processor
	mapMsg    map[string]reflect.Type
}

var _ network.Processor = (*ForwardProcessor)(nil)

func NewForwardProcessor(p network.Processor) *ForwardProcessor {
	return &ForwardProcessor{
		processor: p,
		mapMsg:    map[string]reflect.Type{},
	}
}

func (f *ForwardProcessor) Route(m interface{}, userData interface{}) error {
	log.Infof("ForwardProcessor Route: %+v", m)

	if _, ok := m.(*message.S2S_Msg); ok {
		return f.processor.Route(m, userData)
	}
	if _, ok := m.(*message.S2C_Msg); ok {
		return f.processor.Route(m, userData)
	}

	//if _, ok := m.(*message.C2S_Msg); ok {
	//	return f.processor.Route(m, userData)
	//}

	return f.Forward(m, userData)
}

func (f *ForwardProcessor) Unmarshal(data []byte) (interface{}, error) {
	v, err := f.processor.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	switch m := v.(type) {
	case *message.S2S_Msg, *message.S2S_Reg, *message.Bind:
		return m, nil
	case *message.S2C_Msg:
		return m, nil
	//case *message.S2S_Msg, *message.C2S_Msg:
	//	return m, nil
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
	msgID := message.GetMsgId(m)
	if _, ok := f.mapMsg[msgID]; !ok {
		f.mapMsg[msgID] = reflect.TypeOf(m)
	}

	//b, ok := m.([]byte)
	//if ok {
	//	c2s := &message.C2S_Msg{
	//		MsgID: message.GetMsgId(m),
	//		Body:  b,
	//	}
	//	return f.processor.Route(c2s, userData)
	//}

	return f.processor.Route(m, userData)
}

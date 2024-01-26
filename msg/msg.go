package msg

type (
	C2S_Msg struct {
		MsgID   string
		MsgType interface{}
		Body    []byte
	}

	S2S_Msg struct {
		From  string
		To    string
		MsgID string
		Body  []byte
	}

	Login struct {
		Code    int
		Message string
	}

	Greeting struct {
		Code    int
		Message string
	}
)

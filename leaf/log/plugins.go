package log

type ZapCnfKey struct {
}

type ZapCnf struct {
	Debug bool //调试模式 本地打印
	Net   string
	Addr  string
	Tag   string
}

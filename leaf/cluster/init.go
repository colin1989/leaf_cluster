package cluster

var defaultCluster *Cluster

func Init() {
	if defaultCluster == nil {
		return
	}

	register()
	go defaultCluster.listen()
}

func Destroy() {
	if defaultCluster == nil {
		return
	}
	defaultCluster.destroy()
	close(defaultCluster.stop)
}

func register() {
	//defaultCluster.chanRPC.Register("NewWorldFunc", NewWorldFunc)
	//defaultCluster.chanRPC.Register("NewSessionFunc", NewSessionFunc)
}

func NewSessionFunc(args []interface{}) {
}

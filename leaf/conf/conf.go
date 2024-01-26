package conf

var (
	LenStackBuf = 4096

	// log
	LogDebug bool // log debug
	LogLevel string
	LogPath  string
	LogAddr  string
	LogTag   string

	// console
	ConsolePort   int
	ConsolePrompt string = "dhf# "
	ProfilePath   string

	// cluster
	ListenAddr      string
	ConnAddrs       []string
	PendingWriteNum int

	MsgSizeLimit = Megabyte / 2
)

const (
	ByteSize = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
)

package log

// Logger is a logger interface.
type Logger interface {

	// Init initialises options
	Init(options ...Option) error

	// The Logger options
	Options() Options

	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger

	// Log writes a log entry
	Log(level Level, keyvals ...interface{})

	// Logf writes a formatted log entry
	Logf(level Level, format string, keyvals ...interface{})

	// LogW writes a field log
	LogW(level Level, msg string, args ...Field)

	Printf(format string, args ...interface{})

	// Level 变更
	OnChangeLevel(level Level)

	// String returns the name of logger
	String() string
}

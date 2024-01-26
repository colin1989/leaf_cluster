package log

import (
	"sync"
)

// DefaultLogger is default logger.
var (
	DefaultLogger Logger
)

// globalLogger is designed as a global logger in current process.
var global = &loggerAppliance{}

// loggerAppliance is the proxy of `Logger` to
// make logger change will affect all sub-logger.
type loggerAppliance struct {
	lock sync.Mutex
	Logger
	helper *Helper
}

// todo
func init() {
	c := &Config{}
	c.Build()
	SetLogger(DefaultLogger)
}

func (l *loggerAppliance) SetLogger(in Logger) {
	if in == nil {
		return
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	l.Logger = in
	l.helper = NewHelper(l.Logger)
}

// SetLogger should be called before any other log call.
// And it is NOT THREAD SAFE.
func SetLogger(logger Logger) {
	global.SetLogger(logger)
}

// Debug logs a message at `debug` level.
func Debug(a ...interface{}) {
	global.helper.Debug(a...)
}

// Debugf logs a message at `debug` level.
func Debugf(template string, a ...interface{}) {
	global.helper.Debugf(template, a...)
}

// DebugW writes msg along with fields into access log at `debug` level.
func DebugW(template string, a ...Field) {
	global.helper.DebugW(template, a...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	global.helper.Info(a...)
}

// Infof logs a message at `info` level.
func Infof(template string, a ...interface{}) {
	global.helper.Infof(template, a...)
}

// InfoW writes msg along with fields into access log at `info` level.
func InfoW(template string, a ...Field) {
	global.helper.InfoW(template, a...)
}

func Printf(format string, args ...interface{}) {
	global.helper.Infof(format, args...)
}

// Warn logs a message at `warn` level.
func Warn(a ...interface{}) {
	global.helper.Warn(a...)
}

// Warnf logs a message at `warn` level.
func Warnf(format string, a ...interface{}) {
	global.helper.Warnf(format, a...)
}

// WarnW writes msg along with fields into access log at `warn` level.
func WarnW(template string, a ...Field) {
	global.helper.WarnW(template, a...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	global.helper.Error(a...)
}

// Errorf logs a message at `error` level.
func Errorf(format string, a ...interface{}) {
	global.helper.Errorf(format, a...)
}

// ErrorW writes msg along with fields into access log at `error` level.
func ErrorW(template string, a ...Field) {
	global.helper.ErrorW(template, a...)
}

// Fatal logs a message at `fatal` level.
func Fatal(a ...interface{}) {
	global.helper.Fatal(a...)
}

// Fatalf logs a message at `fatal` level.
func Fatalf(format string, a ...interface{}) {
	global.helper.Fatalf(format, a...)
}

// FatalW writes msg along with fields into access log at `fatal` level.
func FatalW(template string, a ...Field) {
	global.helper.FatalW(template, a...)
}

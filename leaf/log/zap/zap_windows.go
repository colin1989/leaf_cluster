//go:build windows

package zap

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

type zaplog struct {
	cfg  zap.Config
	opts log.Options

	// logger
	zap *zap.Logger
	//
	sync.RWMutex
	fields map[string]interface{}
}

// Init (opts...) should only overwrite provided options
func (z *zaplog) Init(opts ...log.Option) error {
	for _, o := range opts {
		o(&z.opts)
	}
	zcfg, ok := z.opts.Context.Value(log.ZapCnfKey{}).(log.ZapCnf)
	if !ok {
		return errors.Errorf("zaplog  config failed not ok: %v", zcfg)
	}
	log.InfoW("zap log config", log.Any("config", zcfg))

	// Set log Level if not default
	// DPanicLevel, PanicLevel, FatalLevel 这三个日志级别同时输出stacktrace
	stacktracePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DPanicLevel
	})

	var core zapcore.Core
	//if zcfg.Debug {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.Level = zap.NewAtomicLevel()
	writer := zapcore.Lock(os.Stdout)                                                               // 设置日志输出的设备，这里还是使用标准输出，也可以传一个 File 类型让它写入到文件
	core = zapcore.NewCore(zapcore.NewJSONEncoder(zapConfig.EncoderConfig), writer, zap.DebugLevel) // 设置日志的默认级别， 前置有日志级别设置，此处最低级别即可
	z.zap = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(z.opts.CallerSkipCount), zap.AddStacktrace(stacktracePriority))
	return nil
}

// Log writes a log entry
func (z *zaplog) Log(level log.Level, keyvals ...interface{}) {

	z.RLock()
	data := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		data = append(data, zap.Any(k, v))
	}
	z.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(keyvals...)

	switch lvl {
	case zap.DebugLevel:
		z.zap.Debug(msg, data...)
	case zap.InfoLevel:
		z.zap.Info(msg, data...)
	case zap.WarnLevel:
		z.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		z.zap.Error(msg, data...)
	case zap.FatalLevel:
		z.zap.Fatal(msg, data...)
	}
}

// Logf writes a formatted log entry
func (z *zaplog) Logf(level log.Level, format string, keyvals ...interface{}) {
	z.RLock()
	data := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		data = append(data, zap.Any(k, v))
	}
	z.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, keyvals...)
	switch lvl {
	case zap.DebugLevel:
		z.zap.Debug(msg, data...)
	case zap.InfoLevel:
		z.zap.Info(msg, data...)
	case zap.WarnLevel:
		z.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		z.zap.Error(msg, data...)
	case zap.FatalLevel:
		z.zap.Fatal(msg, data...)
	}
}

func (z *zaplog) Printf(format string, args ...interface{}) {
	z.Logf(log.InfoLevel, format, args...)
}

func (z *zaplog) LogW(level log.Level, msg string, args ...log.Field) {
	z.RLock()
	data := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		data = append(data, zap.Any(k, v))
	}
	z.RUnlock()

	for _, v := range args {
		data = append(data, fieldToZapEncode(v))
	}

	lvl := loggerToZapLevel(level)

	switch lvl {
	case zap.DebugLevel:
		z.zap.Debug(msg, data...)
	case zap.InfoLevel:
		z.zap.Info(msg, data...)
	case zap.WarnLevel:
		z.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		z.zap.Error(msg, data...)
	case zap.FatalLevel:
		z.zap.Fatal(msg, data...)
	}
}

func (z *zaplog) String() string {
	return "zap"
}

// 日志变更
func (z *zaplog) OnChangeLevel(level log.Level) {
	fmt.Printf("OnChangeLevel msg: %s", level.String())
	z.opts.Level = level
	return
}

func (z *zaplog) Options() log.Options {
	// not guard against options Context values
	return z.opts
}

// Fields set fields to always be logged
func (z *zaplog) Fields(fields map[string]interface{}) log.Logger {
	return nil
}

// NewLogger New builds a new logger based on options
func NewLogger(opts ...log.Option) log.Logger {
	// Default options
	options := log.Options{
		Level:           log.InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 2,
	}

	l := &zaplog{opts: options}
	if err := l.Init(opts...); err != nil {
		log.ErrorW("new logger", log.FieldErr(err))
		//panic(err)
		return nil
	}

	return l
}

func loggerToZapLevel(level log.Level) zapcore.Level {
	switch level {
	case log.TraceLevel, log.DebugLevel:
		return zap.DebugLevel
	case log.InfoLevel:
		return zap.InfoLevel
	case log.WarnLevel:
		return zap.WarnLevel
	case log.ErrorLevel:
		return zap.ErrorLevel
	case log.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func zapToLoggerLevel(level zapcore.Level) log.Level {
	switch level {
	case zap.DebugLevel:
		return log.DebugLevel
	case zap.InfoLevel:
		return log.InfoLevel
	case zap.WarnLevel:
		return log.WarnLevel
	case zap.ErrorLevel:
		return log.ErrorLevel
	case zap.FatalLevel:
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}

func fieldToZapEncode(arg log.Field) zap.Field {
	return zap.Any(arg.Key, arg.Val.V)
}

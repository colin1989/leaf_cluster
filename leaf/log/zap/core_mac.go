//go:build darwin

package zap

import (
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"log/syslog"
)

type syslogCore struct {
	zapcore.LevelEnabler
	enc zapcore.Encoder
	out *syslogWriteSyncer
}

// NewCore creates a Core that writes logs to a WriteSyncer.
func NewCore(enc zapcore.Encoder, ws *syslogWriteSyncer, enab zapcore.LevelEnabler) zapcore.Core {
	return &syslogCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
	}
}

func (s *syslogCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if s.Enabled(ent.Level) {
		return ce.AddCore(ent, s)
	}
	return ce
}

func (s *syslogCore) With(fields []zapcore.Field) zapcore.Core {

	clone := s.clone()
	for i := range fields {
		fields[i].AddTo(clone.enc)
	}
	return clone

}

func (s *syslogCore) Enabled(zapcore.Level) bool {
	return true
}

func (c *syslogCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}

	level, err := toSyslogLevel(ent.Level)
	c.out.Write(level, buf.String())
	buf.Free()
	if err != nil {
		return err
	}
	return nil
}

func (s *syslogCore) Sync() error {
	return nil
}

//

func (s *syslogCore) clone() *syslogCore {
	return &syslogCore{
		LevelEnabler: s.LevelEnabler,
		enc:          s.enc.Clone(),
		out:          s.out,
	}
}

// level
func toSyslogLevel(level zapcore.Level) (syslog.Priority, error) {
	switch level {
	case zapcore.DebugLevel:
		return syslog.LOG_DEBUG, nil
	case zapcore.InfoLevel:
		return syslog.LOG_INFO, nil
	case zapcore.WarnLevel:
		return syslog.LOG_WARNING, nil
	case zapcore.ErrorLevel:
		return syslog.LOG_ERR, nil
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return syslog.LOG_CRIT, nil
	default:
		return syslog.LOG_DEBUG, errors.Errorf("unknown log level: %v", level)
	}
}

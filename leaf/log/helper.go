package log

import "os"

type Helper struct {
	logger Logger
}

func NewHelper(logger Logger) *Helper {
	return &Helper{logger: logger}
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.logger.Log(TraceLevel, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.logger.Logf(TraceLevel, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.logger.Log(DebugLevel, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.logger.Logf(DebugLevel, template, args...)
}

func (h *Helper) DebugW(template string, args ...Field) {
	if !h.logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.logger.LogW(DebugLevel, template, args...)
}

func (h *Helper) Info(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.logger.Log(InfoLevel, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.logger.Logf(InfoLevel, template, args...)
}

func (h *Helper) InfoW(template string, args ...Field) {
	if !h.logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.logger.LogW(InfoLevel, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.logger.Log(WarnLevel, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.logger.Logf(WarnLevel, template, args...)
}

func (h *Helper) WarnW(template string, args ...Field) {
	if !h.logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.logger.LogW(WarnLevel, template, args...)
}
func (h *Helper) Error(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.logger.Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.logger.Logf(ErrorLevel, template, args...)
}

func (h *Helper) ErrorW(template string, args ...Field) {
	if !h.logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.logger.LogW(ErrorLevel, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.logger.Log(FatalLevel, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.logger.Logf(FatalLevel, template, args...)
	os.Exit(1)
}

func (h *Helper) FatalW(template string, args ...Field) {
	if !h.logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.logger.LogW(FatalLevel, template, args...)
	os.Exit(1)
}

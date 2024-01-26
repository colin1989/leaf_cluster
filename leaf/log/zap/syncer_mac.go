//go:build darwin

package zap

import (
	"fmt"
	"log/syslog"
	"sync"
	"time"
)

type logEntry struct {
	level syslog.Priority
	buf   string
}

type syslogWriteSyncer struct {
	writer  *syslog.Writer
	entryCh chan logEntry

	//
	quitCh   chan struct{}
	onceQuit sync.Once
}

type syncCnf struct {
	network, raddr string
	priority       syslog.Priority
	tag            string
	debug          bool
}

// newSyslogSyncer creates a syslog writer
func newSyslogWriteSyncer(cnf *syncCnf) (*syslogWriteSyncer, error) {
	w, err := newSyslog(cnf)
	if err != nil {
		return nil, err
	}

	syncer := &syslogWriteSyncer{
		writer:  w,
		entryCh: make(chan logEntry, 10240),
		quitCh:  make(chan struct{}, 1),
	}

	go syncer.RunLoop(cnf)
	return syncer, nil
}

func newSyslog(cnf *syncCnf) (*syslog.Writer, error) {
	return syslog.Dial("tcp", cnf.raddr, syslog.LOG_WARNING|syslog.LOG_DAEMON, cnf.tag)
}

func (w *syslogWriteSyncer) Write(level syslog.Priority, buf string) {
	entry := logEntry{
		level: level,
		buf:   buf,
	}

	select {
	case w.entryCh <- entry:
	default:
		fmt.Println("write log failed")
	}
}

func (w *syslogWriteSyncer) flush(entry logEntry) (err error) {

	switch entry.level {
	case syslog.LOG_DEBUG:
		err = w.writer.Debug(entry.buf)
	case syslog.LOG_INFO:
		err = w.writer.Info(entry.buf)
	case syslog.LOG_WARNING:
		err = w.writer.Warning(entry.buf)
	case syslog.LOG_ERR:
		err = w.writer.Err(entry.buf)
	case syslog.LOG_CRIT:
		err = w.writer.Crit(entry.buf)
	default:
		err = w.writer.Debug(entry.buf)
	}
	return
}

func (w *syslogWriteSyncer) Quit() {
	w.onceQuit.Do(func() {
		close(w.entryCh)
		close(w.quitCh)
	})

}
func (w *syslogWriteSyncer) RunLoop(cnf *syncCnf) {

	for {

		select {
		case _, ok := <-w.quitCh:
			if !ok {
				return
			}
		case entry, ok := <-w.entryCh:

			if !ok {
				return
			}
			err := w.flush(entry)
			// 推送错误，则重连后再发送
			if err != nil {
				fmt.Printf("send to syslog failed. err[%v]\n", err)
				// 等待1s进行重连
				time.Sleep(time.Second)
				if writer, err := newSyslog(cnf); err == nil {
					w.writer = writer
					// one more try
					w.flush(entry)
				}
			}

		}
	}

}

package util

//
//import (
//	"common/lib/OpenTelemetry/Tracing"
//	"github.com/name5566/leaf/conf"
//	"github.com/name5566/leaf/log"
//	"runtime"
//)
//
//func ProtectRun(fn interface{}, args ...interface{}) {
//	defer func() {
//		if r := recover(); r != nil {
//			if conf.LenStackBuf > 0 {
//				buf := make([]byte, conf.LenStackBuf)
//				l := runtime.Stack(buf, false)
//				log.Errorf("%v: %s", r, buf[:l])
//			} else {
//				log.Errorf("%v", r)
//			}
//		}
//	}()
//
//	// execute
//	switch fn.(type) {
//	case func():
//		fn.(func())()
//	case func(tracer *Tracing.TracerObj):
//		fn.(func(tracer *Tracing.TracerObj))(args[0].(*Tracing.TracerObj))
//	default:
//		panic("bug")
//	}
//	return
//}

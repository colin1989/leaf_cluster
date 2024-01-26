package log

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type stdLogger struct {
	sync.RWMutex
	opts Options

	fields []Field
	// std
	pool *sync.Pool
}

// Init (opts...) should only overwrite provided options
func (l *stdLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	l.pool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	return nil
}

func (l *stdLogger) String() string {
	return "std"
}

func (l *stdLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	nfields := make(map[string]interface{}, len(l.opts.Fields))
	for k, v := range l.opts.Fields {
		nfields[k] = v
	}
	l.Unlock()

	for k, v := range fields {
		nfields[k] = v
	}

	return &stdLogger{opts: Options{
		Level:           l.opts.Level,
		Fields:          nfields,
		Out:             l.opts.Out,
		CallerSkipCount: l.opts.CallerSkipCount,
		Context:         l.opts.Context,
	}}
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// logCallerfilePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func logCallerfilePath(loggingFilePath string) string {
	// To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l *stdLogger) Printf(format string, v ...interface{}) {
	l.Logf(InfoLevel, format, v...)
}

func (l *stdLogger) Log(level Level, v ...interface{}) {
	// TODO decide does we need to write message if log level not used?
	var b strings.Builder
	b.WriteString("{")
	// level
	b.WriteString("\"level\": ")
	b.WriteString("\"")
	b.WriteString(level.String())
	b.WriteString("\"")
	b.WriteString(",")
	// ts
	b.WriteString("\"ts\": ")
	b.WriteString("\"")
	b.WriteString(time.Now().String())
	b.WriteString("\"")
	b.WriteString(",")

	// msg
	b.WriteString("\"msg\": ")
	b.WriteString("\"")
	b.WriteString(fmt.Sprint(v...))
	b.WriteString("\"")

	m := make(map[string]FiledValue, 0)
	size := len(l.fields)
	keys := make([]string, 0, size)

	for _, kv := range l.fields {
		keys = append(keys, kv.Key)
		m[kv.Key] = kv.Val
	}

	// CallerSkipCount
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {

		m["caller"] = FiledValue{
			V: fmt.Sprintf("%s:%d", logCallerfilePath(file), line),
			T: TypeString,
		}
		keys = append(keys, "caller")
	}

	sort.Strings(keys)

	for _, k := range keys {
		b.WriteString(",")
		b.WriteString(encode(k, m[k]))
	}

	b.WriteString("}")
	b.WriteString("\n")
	l.opts.Out.Write([]byte(b.String()))
}

func (l *stdLogger) Logf(level Level, format string, v ...interface{}) {

	//	 TODO decide does we need to write message if log level not used?

	var b strings.Builder
	b.WriteString("{")
	// level
	b.WriteString("\"level\": ")
	b.WriteString("\"")
	b.WriteString(level.String())
	b.WriteString("\"")
	b.WriteString(",")
	// ts
	b.WriteString("\"ts\": ")
	b.WriteString("\"")
	b.WriteString(time.Now().String())
	b.WriteString("\"")
	b.WriteString(",")

	// msg
	b.WriteString("\"msg\": ")
	b.WriteString("\"")
	b.WriteString(fmt.Sprintf(format, v...))
	b.WriteString("\"")

	m := make(map[string]FiledValue, 0)
	size := len(l.fields)
	keys := make([]string, 0, size)

	for _, kv := range l.fields {
		keys = append(keys, kv.Key)
		m[kv.Key] = kv.Val
	}

	// CallerSkipCount
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {

		m["caller"] = FiledValue{
			V: fmt.Sprintf("%s:%d", logCallerfilePath(file), line),
			T: TypeString,
		}
		keys = append(keys, "caller")
	}

	sort.Strings(keys)

	for _, k := range keys {
		b.WriteString(",")
		b.WriteString(encode(k, m[k]))
	}

	b.WriteString("}")
	b.WriteString("\n")
	l.opts.Out.Write([]byte(b.String()))
}

func (l *stdLogger) LogW(level Level, msg string, args ...Field) {

	//	 TODO decide does we need to write message if log level not used?

	/*if !l.opts.Level.Enabled(level) {
		return
	}*/

	var b strings.Builder
	b.WriteString("{")
	// level
	b.WriteString("\"level\": ")
	b.WriteString("\"")
	b.WriteString(level.String())
	b.WriteString("\"")
	b.WriteString(",")
	// ts
	b.WriteString("\"ts\": ")
	b.WriteString("\"")
	b.WriteString(time.Now().String())
	b.WriteString("\"")
	b.WriteString(",")

	// msg
	b.WriteString("\"msg\": ")
	b.WriteString("\"")
	b.WriteString(msg)
	b.WriteString("\"")

	m := make(map[string]FiledValue, len(args)+1)
	size := len(l.fields) + len(args)
	keys := make([]string, 0, size)

	for _, kv := range l.fields {
		keys = append(keys, kv.Key)
		m[kv.Key] = kv.Val
	}

	for _, kv := range args {
		keys = append(keys, kv.Key)
		m[kv.Key] = kv.Val
	}

	// CallerSkipCount
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {

		m["caller"] = FiledValue{
			V: fmt.Sprintf("%s:%d", logCallerfilePath(file), line),
			T: TypeString,
		}
		keys = append(keys, "caller")
	}

	sort.Strings(keys)

	for _, k := range keys {
		b.WriteString(",")
		b.WriteString(encode(k, m[k]))
	}

	b.WriteString("}")
	b.WriteString("\n")
	/*buf := l.pool.Get().(*bytes.Buffer)
	buf.WriteString(level.String())

	for k, v := range fields {
		_, _ = fmt.Fprintf(buf, "%s=%v \n", k, v)
	}*/

	l.opts.Out.Write([]byte(b.String()))
}

// 日志变更
func (l *stdLogger) OnChangeLevel(level Level) {
	l.opts.Level = level
	return
}

func encode(key string, val FiledValue) string {
	switch val.T {
	case TypeString:
		return fmt.Sprintf("\"%s\": \"%s\"", key, val.V)
	case TypeInt:
		return fmt.Sprintf("\"%s\": %d", key, val.V)
	case TypeFloat:
		return fmt.Sprintf("\"%s\": %f", key, val.V)
	case TypeBool:
		return fmt.Sprintf("\"%s\": %t", key, val.V)

	}

	return encodeAny(key, val.V)
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func encodeAny(key, data interface{}) string {
	v := indirect(data)
	if v == nil {
		return fmt.Sprintf("\"%s\": \"{}\"", key)
	}

	dt := reflect.TypeOf(data)
	kind := dt.Kind()

	switch kind {
	case reflect.Bool:
		return fmt.Sprintf("\"%s\": %t", key, v)
	case reflect.String:
		return fmt.Sprintf("\"%s\": \"%s\"", key, v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("\"%s\": %d", key, v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("\"%s\": %d", key, v)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("\"%s\": %f", key, v)
	}
	err, ok := data.(error)
	if ok {
		return fmt.Sprintf("\"%s\": \"%s\"", key, err.Error())
	}

	return fmt.Sprintf("\"%s\": %+v", key, v)
}

func (l *stdLogger) Options() Options {
	// not guard against options Context values
	l.RLock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	l.RUnlock()
	return opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           DebugLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stdout,
		CallerSkipCount: 1,
		Context:         context.Background(),
	}

	l := &stdLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(FatalLevel, err)
	}

	return l
}

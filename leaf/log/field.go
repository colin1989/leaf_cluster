package log

import (
	"fmt"
	"strings"
	"time"
)

type FiledType int

const (
	TypeString FiledType = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeAny
)

type Field struct {
	Key string
	Val FiledValue
}

type FiledValue struct {
	V interface{}
	T FiledType
}

func String(key, val string) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeString,
		},
	}
}

func Float64(key string, val float64) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeFloat,
		},
	}
}

func Bool(key string, val bool) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeBool,
		},
	}
}

func Int64(key string, val int64) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeInt,
		},
	}
}

func Int(key string, val int) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeInt,
		},
	}
}

func Int32(key string, val int32) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeInt,
		},
	}
}

func UInt(key string, val uint) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeInt,
		},
	}
}

func Duration(key string, val time.Duration) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeString,
		},
	}
}

func Any(key string, val interface{}) Field {
	return Field{
		Key: key,
		Val: FiledValue{
			V: val,
			T: TypeAny,
		},
	}
}

func FieldString(key, value string) Field {
	value = strings.Replace(value, " ", ".", -1)
	return String(key, value)
}

func FieldMod(value string) Field {
	return FieldString("mod", value)
}

func FieldErr(err error) Field {
	if err == nil {
		return Field{
			Key: "err",
			Val: FiledValue{
				V: "",
				T: TypeString,
			},
		}
	}
	return Field{
		Key: "err",
		Val: FiledValue{
			V: err.Error(),
			T: TypeString,
		},
	}
}

func FieldKey(value string) Field {
	return String("key", value)
}

func FieldAddr(value string) Field {
	return String("addr", value)
}

func FieldName(value string) Field {
	return String("name", value)
}

// FieldType ...
func FieldType(value string) Field {
	return String("type", value)
}

// FieldMessage ...
func FieldMessage(value string) Field {
	return String("message", value)
}

func FieldErrKind(value string) Field {
	return String("errKind", value)
}

// FieldMethod ...
func FieldMethod(value string) Field {
	return String("method", value)
}

// FieldEvent ...
func FieldEvent(value string) Field {
	return String("event", value)
}

// FieldCost 耗时时间
func FieldCost(value time.Duration) Field {
	return String("cost", fmt.Sprintf("%.3f(ms)", float64(value.Round(time.Microsecond))/float64(time.Millisecond)))
}

// FieldStack ...
func FieldStack(value []byte) Field {
	return String("stack", string(value))
}

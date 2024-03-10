package stringify

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// OptionFunction defines how to stringify certain types.
type OptionFunction func(any) string

// Options map contains OptionFunctions with the reflect.TypeOf.String() string keys.
var Options map[string]OptionFunction

type any interface{}

// Parse stringifies anything under the interface.
func Parse(obj any) string {
	response := ""
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	def, ok := Options[t.String()]
	if ok {
		return def(val.Interface())
	}
	switch t.Kind() {
	case reflect.Struct:
		response = response + parseStruct(val)
	case reflect.Slice:
		response = response + parseSlice(val)
	case reflect.Map:
		response = response + parseMap(val)
	case reflect.Ptr:
		response = response + parsePointer(val)
	case reflect.String:
		response = fmt.Sprintf(`"%v"`, reflect.ValueOf(obj).Interface())
	case reflect.Func:
		response = response + parseFunc(val)
	default:
		response = fmt.Sprintf("%v", reflect.ValueOf(obj).Interface())
	}
	return response
}

func parseStruct(val reflect.Value) string {
	response := "{"
	for i := 0; i < val.NumField(); i++ {
		if i != 0 {
			response = response + ", "
		}
		response = response + `"` + val.Type().Field(i).Name + `"`
		if val.Type().Field(i).IsExported() {
			response = response + ": " + Parse(val.Field(i).Interface())
		} else {
			response = response + ": " + Parse(getUnexportedField(val, i))
		}
	}
	return response + "}"
}

func parseSlice(val reflect.Value) string {
	if val.IsNil() {
		return "null"
	}
	response := "["
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i)
		if i == 0 {
			response = response + Parse(v.Interface())
		} else {
			response = response + ", " + Parse(v.Interface())
		}
	}
	return response + "]"
}

func parseMap(val reflect.Value) string {
	if val.IsNil() {
		return "nil"
	}
	response := "{"
	for i, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if i == 0 {
			response = response + Parse(e.Interface()) + ":" + Parse(v.Interface())
		} else {
			response = response + ", " + Parse(e.Interface()) + ": " + Parse(v.Interface())
		}
	}
	return response + "}"
}

func parsePointer(val reflect.Value) string {
	if val.IsNil() {
		return "nil"
	}
	return Parse(val.Elem().Interface())
}

func parseFunc(val reflect.Value) string {
	return `"` + runtime.FuncForPC(val.Pointer()).Name() + `"`
}

func getUnexportedField(val reflect.Value, field int) interface{} {
	rs2 := reflect.New(val.Type()).Elem()
	rs2.Set(val)
	rf := rs2.Field(field)
	return reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Interface()
}

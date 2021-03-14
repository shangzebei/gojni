package utils

import (
	"fmt"
	"reflect"
	"strings"

	"gitee.com/aifuturewell/gojni/jni"
)

var methods map[string]reflect.Method

func GetJNIMethods() map[string]reflect.Method {
	return methods
}

func GetMethodWithName(name string) *reflect.Method {
	if v, b := methods[strings.ToLower(name)]; b {
		return &v
	} else {
		return nil
	}
}

func init() {
	var vm jni.Env
	methods = make(map[string]reflect.Method)
	f := reflect.TypeOf(vm)
	for i := 0; i < f.NumMethod(); i++ {
		m := f.Method(i)
		methods[strings.ToLower(m.Name)] = m
	}
}

func CallJni(format string, args ...interface{}) reflect.Value {
	m := GetMethodWithName(format)
	if m == nil {
		panic(fmt.Errorf("method def is error %s", format))
	}
	var params []reflect.Value
	index := m.Type.NumIn() - 1
	for i, v := range args {
		of := reflect.ValueOf(v)
		if i >= index {
			params = append(params, reflect.ValueOf(uint64(JabValueToUint(of))))
		} else {
			params = append(params, of)
		}
	}
	//fmt.Println(params)
	ret := m.Func.Call(params)
	if len(ret) != 0 {
		return ret[0]
	}
	return reflect.Value{}
}

//TODO not impl
func JabValueToUint(r reflect.Value) uintptr {
	env := jni.AutoGetCurrentThreadEnv()
	switch r.Type().Kind() {
	case reflect.String:
		return env.NewString(r.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uintptr(r.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintptr(r.Uint())
	case reflect.Float32, reflect.Float64:
		//fmt.Println(reflect.ValueOf(20.65).Addr().Pointer())
		return uintptr(r.Float())
	default:
		panic(fmt.Sprintf("Return not support type %s", r.Kind().String()))
	}
}

func GetFormatCallFunc(format string, siged string) string {
	var orig = "object"
	if v, b := SigDecodeMap[siged]; b {
		orig = v
	}
	//fmt.Println(siged,orig)
	return fmt.Sprintf(format, orig)
}

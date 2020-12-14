package loader

import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"gitee.com/aifuturewell/gojni/jni"
)

//export a1
func a1(p uintptr) uintptr {
	return router("a1", p)
}

//export b1
func b1(p uintptr) uintptr {
	return router("b1", p)
}

//export c1
func c1(p uintptr) uintptr {
	return router("c1", p)
}

//export d1
func d1(p uintptr) uintptr {
	return router("d1", p)
}

//export e1
func e1(p uintptr) uintptr {
	return router("e1", p)
}

//export f1
func f1(p uintptr) uintptr {
	return router("f1", p)
}

//export g1
func g1(p uintptr) uintptr {
	return router("g1", p)
}

//############################################################################

//export a2
func a2(p1, p2 uintptr) uintptr {
	return router("a2", p1, p2)
}

//export b2
func b2(p1, p2 uintptr) uintptr {
	return router("b2", p1, p2)
}

//export c2
func c2(p1, p2 uintptr) uintptr {
	return router("c2", p1, p2)
}

//export d2
func d2(p1, p2 uintptr) uintptr {
	return router("d2", p1, p2)
}

//export e2
func e2(p1, p2 uintptr) uintptr {
	return router("e2", p1, p2)
}

//export f2
func f2(p1, p2 uintptr) uintptr {
	return router("f2", p1, p2)
}

//export g2
func g2(p1, p2 uintptr) uintptr {
	return router("g2", p1, p2)
}

//############################################################

//export a3
func a3(p1, p2, p3 uintptr) uintptr {
	return router("a3", p1, p2, p3)
}

//export b3
func b3(p1, p2, p3 uintptr) uintptr {
	return router("b3", p1, p2, p3)
}

//export c3
func c3(p1, p2, p3 uintptr) uintptr {
	return router("c3", p1, p2, p3)
}

//export d3
func d3(p1, p2, p3 uintptr) uintptr {
	return router("d3", p1, p2, p3)
}

//export e3
func e3(p1, p2, p3 uintptr) uintptr {
	return router("e3", p1, p2, p3)
}

//export f3
func f3(p1, p2, p3 uintptr) uintptr {
	return router("f3", p1, p2, p3)
}

//export g3
func g3(p1, p2, p3 uintptr) uintptr {
	return router("g3", p1, p2, p3)
}

//#################################################################
//export a4
func a4(p1, p2, p3, p4 uintptr) uintptr {
	return router("a4", p1, p2, p3, p4)
}

//export b4
func b4(p1, p2, p3, p4 uintptr) uintptr {
	return router("b4", p1, p2, p3, p4)
}

//export c4
func c4(p1, p2, p3, p4 uintptr) uintptr {
	return router("c4", p1, p2, p3, p4)
}

//export d4
func d4(p1, p2, p3, p4 uintptr) uintptr {
	return router("d4", p1, p2, p3, p4)
}

//export e4
func e4(p1, p2, p3, p4 uintptr) uintptr {
	return router("e4", p1, p2, p3, p4)
}

//export f4
func f4(p1, p2, p3, p4 uintptr) uintptr {
	return router("f4", p1, p2, p3, p4)
}

//export g4
func g4(p1, p2, p3, p4 uintptr) uintptr {
	return router("g4", p1, p2, p3, p4)
}

//##################################################################

//export a5
func a5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("a5", p1, p2, p3, p4, p5)
}

//export b5
func b5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("b5", p1, p2, p3, p4, p5)
}

//export c5
func c5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("c5", p1, p2, p3, p4, p5)
}

//export d5
func d5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("d5", p1, p2, p3, p4, p5)
}

//export e5
func e5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("e5", p1, p2, p3, p4, p5)
}

//export f5
func f5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("f5", p1, p2, p3, p4, p5)
}

//export g5
func g5(p1, p2, p3, p4, p5 uintptr) uintptr {
	return router("g5", p1, p2, p3, p4, p5)
}

//########################################################

//export a6
func a6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("a6", p1, p2, p3, p4, p5, p6)
}

//export b6
func b6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("b6", p1, p2, p3, p4, p5, p6)
}

//export c6
func c6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("c6", p1, p2, p3, p4, p5, p6)
}

//export d6
func d6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("d6", p1, p2, p3, p4, p5, p6)
}

//export e6
func e6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("e6", p1, p2, p3, p4, p5, p6)
}

//export f6
func f6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("f6", p1, p2, p3, p4, p5, p6)
}

//export g6
func g6(p1, p2, p3, p4, p5, p6 uintptr) uintptr {
	return router("g6", p1, p2, p3, p4, p5, p6)
}

//############################################################

//export a7
func a7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("a7", p1, p2, p3, p4, p5, p6, p7)
}

//export b7
func b7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("b7", p1, p2, p3, p4, p5, p6, p7)
}

//export c7
func c7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("c7", p1, p2, p3, p4, p5, p6, p7)
}

//export d7
func d7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("d7", p1, p2, p3, p4, p5, p6, p7)
}

//export e7
func e7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("e7", p1, p2, p3, p4, p5, p6, p7)
}

//export f7
func f7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("f7", p1, p2, p3, p4, p5, p6, p7)
}

//export g7
func g7(p1, p2, p3, p4, p5, p6, p7 uintptr) uintptr {
	return router("g7", p1, p2, p3, p4, p5, p6, p7)
}

//##############################################################

//export a8
func a8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("a8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export b8
func b8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("b8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export c8
func c8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("c8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export d8
func d8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("d8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export e8
func e8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("e8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export f8
func f8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("f8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//export g8
func g8(p1, p2, p3, p4, p5, p6, p7, p8 uintptr) uintptr {
	return router("g8", p1, p2, p3, p4, p5, p6, p7, p8)
}

//#############################################################

//export a9
func a9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("a9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export b9
func b9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("b9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export c9
func c9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("c9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export d9
func d9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("d9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export e9
func e9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("e9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export f9
func f9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("f9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//export g9
func g9(p1, p2, p3, p4, p5, p6, p7, p8, p9 uintptr) uintptr {
	return router("g9", 1, p2, p3, p4, p5, p6, p7, p8, p9)
}

//##############################################################

//export a10
func a10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("a10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export b10
func b10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("b10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export c10
func c10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("c10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export d10
func d10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("d10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export e10
func e10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("e10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export f10
func f10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("f10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//export g10
func g10(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 uintptr) uintptr {
	return router("g10", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
}

//in c type
func router(s string, p ...uintptr) uintptr {
	if f, b := _funcMapper[s]; b {
		rValues := reflect.ValueOf(f.fn).Call(convertParam(f, p...))
		if len(rValues) != 1 {
			return 0
		}
		return convertReturn(rValues[0])
	}
	return 0
}

//TODO not impl
func convertReturn(r reflect.Value) uintptr {
	env := jni.AutoGetCurrentThreadEnv()
	switch r.Type().Kind() {
	case reflect.String:
		return env.NewString(r.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintptr(r.Int())
	case reflect.Float32, reflect.Float64:
		//fmt.Println(reflect.ValueOf(20.65).Addr().Pointer())
		return uintptr(r.Float())
	default:
		panic(fmt.Sprintf("convertReturn not support type %s", r.Kind().String()))
	}
}

func convertParam(f method, params ...uintptr) []reflect.Value {
	var ret []reflect.Value
	lenP := len(f.sig)
	env := jni.AutoGetCurrentThreadEnv()
	for i := 0; i < lenP; i++ {
		s := f.sig[i]
		p := params[i+2]
		switch s.gSig.Kind() {
		case reflect.Int:
			ret = append(ret, reflect.ValueOf(int(p)))
		case reflect.Int64:
			ret = append(ret, reflect.ValueOf(int64(p)))
		case reflect.Float32, reflect.Float64:
			ret = append(ret, reflect.ValueOf(float64(p)))
		case reflect.String:
			jni.CheckNull(p, "jni input str is null")
			pkg := string(env.GetStringUTF(p))
			ret = append(ret, reflect.ValueOf(pkg))
		case reflect.Slice:
			ret = append(ret, convertParamSlice(env, s.gSig, p))
		default:
			panic("convertParam not support")
		}
	}
	return ret
}

func convertParamSlice(env jni.Env, Array reflect.Type, p uintptr) reflect.Value {

	iLen := env.GetArrayLength(p)
	item := Array.Elem()

	switch item.Kind() {
	case reflect.Int:
		iTypes := int(unsafe.Sizeof(C.long(0)))
		jBytes := iLen * iTypes
		ptr := env.GetLongArrayElements(p, true)
		reBytes := C.GoBytes(ptr, C.int(jBytes))
		env.ReleaseLongArrayElements(p, uintptr(ptr), 0)
		head := (*reflect.SliceHeader)(unsafe.Pointer(&reBytes))
		head.Cap /= iTypes
		head.Len /= iTypes
		return reflect.ValueOf(*(*[]int)(unsafe.Pointer(head)))
	case reflect.Int32:
		iTypes := int(unsafe.Sizeof(C.int(0)))
		jBytes := iLen * iTypes
		ptr := env.GetIntArrayElements(p, true)
		reBytes := C.GoBytes(ptr, C.int(jBytes))
		env.ReleaseIntArrayElements(p, uintptr(ptr), 0)
		head := (*reflect.SliceHeader)(unsafe.Pointer(&reBytes))
		head.Cap /= iTypes
		head.Len /= iTypes
		return reflect.ValueOf(*(*[]int32)(unsafe.Pointer(head)))
	case reflect.String:
		var temp []string = make([]string, iLen)
		for i := 0; i < iLen; i++ {
			temp[i] = string(env.GetStringUTF(env.GetObjectArrayElement(p, i)))
		}
		return reflect.ValueOf(temp)
	case reflect.Uint8:
		iTypes := 1
		jBytes := iLen * iTypes
		ptr := env.GetByteArrayElements(p, true)
		reBytes := C.GoBytes(ptr, C.int(jBytes))
		env.ReleaseByteArrayElements(p, uintptr(ptr), 0)
		head := (*reflect.SliceHeader)(unsafe.Pointer(&reBytes))
		head.Cap /= iTypes
		head.Len /= iTypes
		return reflect.ValueOf(*(*[]byte)(unsafe.Pointer(head)))
	case reflect.Float32:
		iTypes := int(unsafe.Sizeof(C.float(0.0)))
		jBytes := iLen * iTypes
		ptr := env.GetFloatArrayElements(p, true)
		reBytes := C.GoBytes(ptr, C.int(jBytes))
		env.ReleaseFloatArrayElements(p, uintptr(ptr), 0)
		head := (*reflect.SliceHeader)(unsafe.Pointer(&reBytes))
		head.Cap /= iTypes
		head.Len /= iTypes
		return reflect.ValueOf(*(*[]float32)(unsafe.Pointer(head)))
	case reflect.Float64:
		iTypes := int(unsafe.Sizeof(C.double(0.0)))
		jBytes := iLen * iTypes
		ptr := env.GetDoubleArrayElements(p, true)
		reBytes := C.GoBytes(ptr, C.int(jBytes))
		env.ReleaseDoubleArrayElements(p, uintptr(ptr), 0)
		head := (*reflect.SliceHeader)(unsafe.Pointer(&reBytes))
		head.Cap /= iTypes
		head.Len /= iTypes
		return reflect.ValueOf(*(*[]float64)(unsafe.Pointer(head)))
	default:
		panic(fmt.Sprintf("not support Array %s ", item))
	}
	return reflect.Value{}
}

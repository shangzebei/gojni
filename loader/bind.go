package loader

//extern void* a2(void *,void *);
//extern void* a3(void *,void *,void *);
//extern void* a4(void *,void *,void *,void *);
//extern void* a5(void *,void *,void *,void *,void *);
//extern void* a6(void *,void *,void *,void *,void *,void *);
//extern void* a7(void *,void *,void *,void *,void *,void *,void *);
//extern void* a8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* a9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* a10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* a11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* a12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* b2(void *,void *);
//extern void* b3(void *,void *,void *);
//extern void* b4(void *,void *,void *,void *);
//extern void* b5(void *,void *,void *,void *,void *);
//extern void* b6(void *,void *,void *,void *,void *,void *);
//extern void* b7(void *,void *,void *,void *,void *,void *,void *);
//extern void* b8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* b9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* b10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* b11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* b12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* c2(void *,void *);
//extern void* c3(void *,void *,void *);
//extern void* c4(void *,void *,void *,void *);
//extern void* c5(void *,void *,void *,void *,void *);
//extern void* c6(void *,void *,void *,void *,void *,void *);
//extern void* c7(void *,void *,void *,void *,void *,void *,void *);
//extern void* c8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* c9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* c10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* c11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* c12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* d2(void *,void *);
//extern void* d3(void *,void *,void *);
//extern void* d4(void *,void *,void *,void *);
//extern void* d5(void *,void *,void *,void *,void *);
//extern void* d6(void *,void *,void *,void *,void *,void *);
//extern void* d7(void *,void *,void *,void *,void *,void *,void *);
//extern void* d8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* d9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* d10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* d11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* d12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* e2(void *,void *);
//extern void* e3(void *,void *,void *);
//extern void* e4(void *,void *,void *,void *);
//extern void* e5(void *,void *,void *,void *,void *);
//extern void* e6(void *,void *,void *,void *,void *,void *);
//extern void* e7(void *,void *,void *,void *,void *,void *,void *);
//extern void* e8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* e9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* e10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* e11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* e12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* f2(void *,void *);
//extern void* f3(void *,void *,void *);
//extern void* f4(void *,void *,void *,void *);
//extern void* f5(void *,void *,void *,void *,void *);
//extern void* f6(void *,void *,void *,void *,void *,void *);
//extern void* f7(void *,void *,void *,void *,void *,void *,void *);
//extern void* f8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* f9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* f10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* f11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* f12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* g2(void *,void *);
//extern void* g3(void *,void *,void *);
//extern void* g4(void *,void *,void *,void *);
//extern void* g5(void *,void *,void *,void *,void *);
//extern void* g6(void *,void *,void *,void *,void *,void *);
//extern void* g7(void *,void *,void *,void *,void *,void *,void *);
//extern void* g8(void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* g9(void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* g10(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* g11(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
//extern void* g12(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
import "C"
import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
)

type native struct {
	sCls    string
	jCls    jni.Jclass
	env     jni.Env
	natives []jni.JNINativeMethod
}

type method struct {
	fn  interface{}
	sig []args
}

type args struct {
	jSig string
	gSig reflect.Type
}

var NMap = map[int][]unsafe.Pointer{
	2:  {C.a2, C.b2, C.c2, C.d2, C.e2, C.f2, C.g2},
	3:  {C.a3, C.b3, C.c3, C.d3, C.e3, C.f3, C.g3},
	4:  {C.a4, C.b4, C.c4, C.d4, C.e4, C.f4, C.g4},
	5:  {C.a5, C.b5, C.c5, C.d5, C.e5, C.f5, C.g5},
	6:  {C.a6, C.b6, C.c6, C.d6, C.e6, C.f6, C.g6},
	7:  {C.a7, C.b7, C.c7, C.d7, C.e7, C.f7, C.g7},
	8:  {C.a8, C.b8, C.c8, C.d8, C.e8, C.f8, C.g8},
	9:  {C.a9, C.b9, C.c9, C.d9, C.e9, C.f9, C.g9},
	10: {C.a10, C.b10, C.c10, C.e10, C.f10, C.g10},
	11: {C.a11, C.b11, C.c11, C.e11, C.f11, C.g11},
	12: {C.a12, C.b12, C.c12, C.e12, C.f12, C.g12},
}

var (
	fMappers   map[string]method
	statistics map[int]int
	MAX_INDEX  = 10
	MAX_DEP    = 6
)

func init() {
	statistics = make(map[int]int)
	fMappers = make(map[string]method)
}

func WithClass(cls string) *native {
	env := jni.AutoGetCurrentThreadEnv()
	jCls := env.FindClass(strings.ReplaceAll(cls, ".", "/"))
	return &native{jCls: jCls, sCls: cls, env: env}

}

func (n *native) WithClass(cls string) *native {
	n.Done()
	return WithClass(cls)
}

func (n *native) getPFunc(inNum int) (int, int, string) {

	if inNum >= MAX_INDEX {
		panic(fmt.Sprintf("function param overflow max %d numIN %d", MAX_INDEX, inNum))
	}
	dep := statistics[inNum]
	if dep >= MAX_DEP {
		inNum++
		return n.getPFunc(inNum)
	}
	code := fmt.Sprintf("%c%d", dep+97, inNum)
	return inNum, dep, code
}

func (n *native) BindNative(methodName string, def string, fun interface{}) *native {
	jni.CheckNull(n.jCls, fmt.Sprintf("not find class %s", n.sCls))
	ms := utils.EncodeToSig(def)
	//fmt.Println(ms.Sig)
	inNum := len(ms.ParamTyp) + 2
	goF := reflect.TypeOf(fun)
	if len(ms.ParamTyp) != goF.NumIn() {
		panic(fmt.Sprintf("method %s not match fun %s %d", methodName, ms.ParamTyp, goF.NumIn()))
	}
	newNum, dep, code := n.getPFunc(inNum)
	var mArgs []args
	for i := 0; i < goF.NumIn(); i++ {
		n.CheckType(i, methodName, def, ms.ParamTyp[i], goF.In(i))
		mArgs = append(mArgs, args{
			jSig: ms.ParamTyp[i],
			gSig: goF.In(i),
		})
	}
	if goF.NumOut() > 0 {
		n.CheckReturn(methodName, ms.RetTyp, goF.Out(0))
	}

	fMappers[code] = method{
		fn:  fun,
		sig: mArgs,
	}
	cf := NMap[newNum][dep]
	n.natives = append(n.natives, jni.JNINativeMethod{Name: methodName, Sig: ms.Sig, FnPtr: cf})
	statistics[newNum] += 1
	return n
}

var checkMap = map[string]reflect.Type{
	"[I":                  reflect.TypeOf([]int32{}),
	"[Ljava/lang/String;": reflect.TypeOf([]string{}),
	"[B":                  reflect.TypeOf([]byte{}),
	"[J":                  reflect.TypeOf([]int{}),
	"[F":                  reflect.TypeOf([]float32{}),
	"[D":                  reflect.TypeOf([]float64{}),
	"I":                   reflect.TypeOf(int32(0)),
	"Ljava/lang/String;":  reflect.TypeOf(""),
	"B":                   reflect.TypeOf(byte(0)),
	"J":                   reflect.TypeOf(int(1)),
	"F":                   reflect.TypeOf(float32(0.1)),
	"D":                   reflect.TypeOf(float64(0.1)),
}

func (n *native) CheckReturn(mName string, jsig string, gTyp reflect.Type) {
	if v, b := checkMap[jsig]; !b || v != gTyp {
		if b {
			panic(fmt.Sprintf("\n%s method %s return { %s  } not match go type {%s} \nmust use go type ==> %s",
				n.sCls, mName, jsig, gTyp, v))
		} else {
			panic(fmt.Sprintf("%s method %s return { %s  }  not support", n.sCls, mName, jsig))
		}
	}
}

func (n *native) CheckType(i int, mName string, def string, jsig string, gTyp reflect.Type) {
	if v, b := checkMap[jsig]; !b || v != gTyp {
		if b {
			panic(fmt.Sprintf("\n%s method %s definition { %s %d } not match go type {%s} \nmust use go type ==> %s",
				n.sCls, mName, def, i, gTyp, v))
		} else {
			panic(fmt.Sprintf("%s method %s definition { %s %d } sig %s not support", n.sCls, mName, def, i, jsig))
		}
	}
}

func (n *native) Done() {
	if n.env.RegisterNatives(n.jCls, n.natives) < 0 {
		fmt.Println("java class: ", n.sCls)
		n.printNative()
		panic("RegisterNatives error \nplease check java native define ")
	}
}

func (n *native) printNative() {
	for _, nativeMethod := range n.natives {
		fmt.Printf("%s %s\n", utils.Wp(nativeMethod.Name, 10), utils.Wp(nativeMethod.Sig, 100))
	}
}

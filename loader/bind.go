package loader

//extern void* v0();
//extern void* v1(void *);
//extern void* v2(void *,void *);
//extern void* v3(void *,void *,void *);
//extern void* v4(void *,void *,void *,void *);
//extern void* v5(void *,void *,void *,void *,void *);
import "C"
import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
	"strings"
)

type native struct {
	cls string
}

var _native *native

func Bind(cls string) *native {
	if _native == nil {
		_native = &native{cls: strings.ReplaceAll(cls, ".", "/")}
		return _native
	}
	return _native
}

func (n *native) BindNatives(methodName string, def string, fun interface{}) *native {
	env := jni.AutoGetCurrentThreadEnv()
	jni.CheckNull(uintptr(env), "not find env")
	jcls := env.FindClass(n.cls)
	jni.CheckNull(jcls, fmt.Sprintf("not find class %s", n.cls))
	var meds = []jni.JNINativeMethod{{methodName, utils.GetSig(def).Sig, C.v1}}
	if env.RegisterNatives(jcls, meds) < 0 {
		panic("RegisterNatives error")
	}
	return n
}

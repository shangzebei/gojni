package native

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
	"reflect"
	"strings"
)

type Class struct {
	mClass classMeta
	env    *jni.Env
}

type classMeta struct {
	jcls jni.Jclass
	scls string
}

func NewClassMeta(jclass jni.Jclass, cls string) classMeta {
	return classMeta{
		jcls: jclass,
		scls: cls,
	}
}

func LoadClass(name string) Class {
	env := jni.AutoGetCurrentThreadEnv()
	name = strings.ReplaceAll(name, ".", "/")
	jcls := env.FindClass(name)
	if jcls == 0 {
		jni.ThrowException(fmt.Sprintf("not find class %s", name))
	}
	class := Class{mClass: NewClassMeta(jcls, name), env: env}
	return class
}

func (cls Class) New(args ...string) Object {
	signed := EncodeToSig("void()")
	jMethodID := cls.env.GetMethodID(cls.mClass.jcls, "<init>", signed.Sig)
	ret := cls.env.NewObjectA(cls.mClass.jcls, jMethodID)
	v := Value{v: reflect.ValueOf(ret), vSig: *NewSig(cls.mClass.scls)}
	return Object{mClass: cls.mClass, v: v}
}

func (cls Class) StaticInvoke(name string, sig string, args ...interface{}) Value {
	env := jni.AutoGetCurrentThreadEnv()
	sm := EncodeToSig(sig)
	if len(sm.ParamTyp) != len(args) {
		jni.ThrowException("args length is not enough")
	}
	jMethod := env.GetStaticMethodID(cls.mClass.jcls, name, sm.Sig)
	if jMethod == 0 {
		jni.ThrowException(fmt.Sprintf("method %s not find maybe err %s", name, sm.Sig))
	}
	sType := sm.RetTyp.GetSigType()
	defArgs := []interface{}{
		*env, cls.mClass.jcls, jMethod,
	}
	defArgs = append(defArgs, args...)
	ret := utils.CallJni(utils.GetFormatCallFunc("CallStatic%sMethodA", sType), defArgs...)
	return Value{v: ret, vSig: sm.RetTyp}
}

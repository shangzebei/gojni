package native

import (
	"fmt"
	"github.com/shangzebei/gojni/jni"
	"github.com/shangzebei/gojni/utils"
	"reflect"
)

type Value struct {
	vSig Sig
	v    reflect.Value
}

func NewValue(vSig Sig, v reflect.Value) Value {
	return Value{
		vSig: vSig,
		v:    v,
	}
}

func (v Value) AsInt() int {
	return int(v.AsInt64())
}

func (v Value) AsObject() Object {
	r := v.vSig.GetSigType()
	if r[0] != 'L' {
		jni.ThrowException("is not object type")
	}
	cls := r[1 : len(r)-1]
	env := jni.AutoGetCurrentThreadEnv()
	jcls := env.FindClass(cls)
	if jcls == 0 {
		jni.ThrowException(fmt.Sprintf("not find class %s", cls))
	}
	return Object{v: v, mClass: NewClassMeta(jcls, cls)}
}

func (v Value) AsString() string {
	if v.vSig.GetSigType() != "Ljava/lang/String;" {
		jni.ThrowException(fmt.Sprintf("%s is not String type", v.v))
	}
	env := jni.AutoGetCurrentThreadEnv()
	return string(env.GetStringUTF(v.v.Interface().(uintptr)))
}

func (v Value) AsInt64() int64 {
	if v.vSig.GetSigType() != "I" {
		jni.ThrowException(fmt.Sprintf("%s is not int type", v.v))
	}
	return v.v.Int()
}

func (v Value) AsBytes() []byte {
	jni.ThrowException("not impl")
	return nil
}

type Object struct {
	mClass classMeta
	v      Value
}

func (obj Object) ToUintPtr() uintptr {
	i := obj.v.v.Interface().(uintptr)
	if i == 0 {
		jni.ThrowException("obj is null")
	}
	return i
}

func (obj Object) Invoke(name string, sig string, args ...interface{}) Value {
	u := obj.v.v.Interface().(uintptr)
	if u == 0 {
		jni.ThrowException("value is null")
	}
	return invoke(obj.mClass, u, "Call%sMethodA", name, sig, args...)
}

func invoke(cls classMeta, jobj jni.Jobject, callFormal string, name string, sig string, args ...interface{}) Value {
	env := jni.AutoGetCurrentThreadEnv()
	if cls.jcls == 0 {
		jni.ThrowException(fmt.Sprintf("not find class %s", cls.scls))
	}
	sm := EncodeToSig(sig)
	jMethod := env.GetMethodID(cls.jcls, name, sm.Sig)
	if jMethod == 0 {
		jni.ThrowException(fmt.Sprintf("method %s not find maybe err %s", name, sm.Sig))
	}
	sType := sm.RetTyp.GetSigType()
	defArgs := []interface{}{
		*env, jobj, jMethod,
	}
	defArgs = append(defArgs, args...)
	ret := utils.CallJni(utils.GetFormatCallFunc(callFormal, sType), defArgs...)
	return Value{vSig: sm.RetTyp, v: ret}
}

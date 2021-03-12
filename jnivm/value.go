package jnivm

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
	"reflect"
)

type value struct {
	sig *utils.MethodSig
	v   reflect.Value
}

func (v value) AsInt() int {
	return int(v.AsInt64())
}

func (v value) AsObject() Object {
	r := v.sig.RetTyp
	if r[0] != 'L' {
		panic("is not object type")
	}
	cls := r[1 : len(r)-1]
	return Object{v: v, jcls: cls}
}

func (v value) AsString() string {
	if v.sig.RetTyp != "Ljava/lang/String;" {
		panic(fmt.Errorf("%s is not String type", v.v))
	}
	env := jni.AutoGetCurrentThreadEnv()
	return string(env.GetStringUTF(v.v.Interface().(uintptr)))
}

func (v value) AsInt64() int64 {
	if v.sig.RetTyp != "I" {
		panic(fmt.Errorf("%s is not int type", v.v))
	}
	return v.v.Int()
}

func (v value) AsBytes() []byte {
	panic("not impl")
	return nil
}

type Object struct {
	jcls string
	v    value
}

func (obj Object) Invoke(name string, sig string, args ...interface{}) value {
	env := jni.AutoGetCurrentThreadEnv()
	cls := env.FindClass(obj.jcls)
	if cls == 0 {
		panic(fmt.Errorf("not find class %s", obj.jcls))
	}
	return invoke(cls, obj.v.v.Interface().(uintptr), "Call%sMethodA", name, sig, args...)

}

func invoke(cls jni.Jclass, jobj jni.Jobject, callFormal string, name string, sig string, args ...interface{}) value {
	env := jni.AutoGetCurrentThreadEnv()
	if cls == 0 {
		panic(fmt.Errorf("not find class %s", cls))
	}
	sm := utils.EncodeToSig(sig)
	jMethod := env.GetMethodID(cls, name, sm.Sig)
	if jMethod == 0 {
		panic(fmt.Errorf("method %s not find maybe err %s", name, sm.Sig))
	}
	sType := "Object"
	if s, b := utils.SigDecodeMap[sm.RetTyp]; b {
		sType = s
	}
	defArgs := []interface{}{
		env, jobj, jMethod,
	}
	defArgs = append(defArgs, args...)
	ret := callJni(fmt.Sprintf(callFormal, sType), defArgs...)
	return value{sig: sm, v: *ret}
}

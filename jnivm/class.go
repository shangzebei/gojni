package jnivm

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
	"strings"
)

type Class struct {
	jcls jni.Jclass
	env  jni.Env
}

func LoadClass(name string) Class {
	env := jni.AutoGetCurrentThreadEnv()
	name = strings.ReplaceAll(name, ".", "/")
	jcls := env.FindClass(name)
	if jcls == 0 {
		panic(fmt.Errorf("not find class %s", name))
	}
	return Class{jcls: jcls, env: env}
}

func (cls Class) New(args ...string) Object {
	panic("not impl")
	return Object{}
}

func (cls Class) StaticInvoke(name string, sig string, args ...interface{}) value {
	env := jni.AutoGetCurrentThreadEnv()
	sm := utils.EncodeToSig(sig)
	if len(sm.ParamTyp) != len(args) {
		panic("args length is not enough")
	}
	jMethod := env.GetStaticMethodID(cls.jcls, name, sm.Sig)
	if jMethod == 0 {
		panic(fmt.Errorf("method %s not find maybe err %s", name, sm.Sig))
	}
	sType := "Object"
	if s, b := utils.SigDecodeMap[sm.RetTyp]; b {
		sType = s
	}
	defArgs := []interface{}{
		env, cls.jcls, jMethod,
	}
	defArgs = append(defArgs, args...)
	ret := callJni(fmt.Sprintf("CallStatic%sMethodA", sType), defArgs...)
	return value{v: *ret, sig: sm}
}

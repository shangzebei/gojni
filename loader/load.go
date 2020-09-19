package loader

import "C"
import (
	"gitee.com/aifuturewell/gojni/jni"
)

var _onLoad []func(vm uintptr)
var _onUnLoad []func(vm uintptr)

//export JNI_OnLoad
func JNI_OnLoad(vm uintptr, reserved uintptr) int {
	jni.InitJNI(vm)
	for _, f := range _onLoad {
		f(vm)
	}
	if _, v := jni.VM(vm).GetEnv(jni.JNI_VERSION_1_6); v != jni.JNI_OK {
		panic("JNI_OnLoad error")
		return -1
	}
	return jni.JNI_VERSION_1_6
}

//export JNI_OnUnload
func JNI_OnUnload(vm uintptr, reserved uintptr) {
	for _, f := range _onUnLoad {
		f(vm)
	}
}

func OnLoad(f func(vm uintptr)) {
	_onLoad = append(_onLoad, f)
}

func OnUnload(f func(vm uintptr)) {
	_onUnLoad = append(_onUnLoad, f)
}

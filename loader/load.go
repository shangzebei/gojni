package loader

import "C"
import (
	"gitee.com/aifuturewell/gojni/jni"
)

var onLoads []func(vm uintptr)
var onUnLoads []func(vm uintptr)

//export JNI_OnLoad
func JNI_OnLoad(vm uintptr, reserved uintptr) int {
	jni.InitJNI(vm)
	for _, f := range onLoads {
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
	for _, f := range onUnLoads {
		f(vm)
	}
}

func OnLoad(f func(vm uintptr)) {
	onLoads = append(onLoads, f)
}

func OnUnload(f func(vm uintptr)) {
	onUnLoads = append(onUnLoads, f)
}

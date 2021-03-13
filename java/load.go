package java

import "C"
import (
	"gitee.com/aifuturewell/gojni/jni"
	"runtime"
)

var onLoads []func(reg Register)
var onUnLoads []func()

//export JNI_OnLoad
func JNI_OnLoad(vm uintptr, reserved uintptr) int {
	runtime.LockOSThread()

	jni.InitJNI(vm)
	r := Register{vm: jni.VM(vm)}
	for _, f := range onLoads {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					jni.ThrowException(err.(error).Error())
				}
			}()
			f(r)
		}()
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
		f()
	}
	runtime.UnlockOSThread()
}

func OnLoad(f func(reg Register)) {
	onLoads = append(onLoads, f)
}

func OnUnload(f func()) {
	onUnLoads = append(onUnLoads, f)
}

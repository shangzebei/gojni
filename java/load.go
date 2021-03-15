package java

import "C"
import (
	"gitee.com/aifuturewell/gojni/jni"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var onLoads []func(reg Register)
var onUnLoads []func()
var eg = `
######################### GOJNI ERROR ############################

func init() {
	java.OnLoad(func(reg java.Register) {
		...
	})
}

###################################################################

`

//export JNI_OnLoad
func JNI_OnLoad(vm uintptr, reserved uintptr) int {
	runtime.LockOSThread()

	jni.InitJNI(vm)
	r := Register{vm: jni.VM(vm)}
	if len(onLoads) == 0 {
		go func() {
			jni.JavaThrowException("you mast impl java.Onload on func init ." + "\n" + eg)
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}()
	}
	for _, f := range onLoads {
		go func(reg Register) {
			defer func() {
				if err := recover(); err != nil {
					msg := "error"
					if e, b := err.(error); b {
						msg = e.Error()
					}
					if e, b := err.(string); b {
						msg = e
					}
					jni.JavaThrowException(msg + "\n" + string(debug.Stack()))
				}
			}()
			f(reg)
		}(r)
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

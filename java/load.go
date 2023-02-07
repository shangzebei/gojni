package java

import "C"
import (
	"github.com/shangzebei/gojni/jni"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

type FRegister func(reg Register)

var onLoads []FRegister
var onMainLoads []FRegister
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
	var r Register = &registerImpl{vm: jni.VM(vm)}
	if len(onLoads) == 0 && len(onMainLoads) == 0 {
		go func() {
			jni.JavaThrowException("you mast impl java.Onload on func init ." + "\n" + eg)
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}()
	}

	//run on other thread
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
			reg.Done()
		}(r)
	}

	//run on main thread
	for _, load := range onMainLoads {
		load(r)
		r.Done()
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

// OnLoad used in golang code but running on other thread
func OnLoad(f FRegister) {
	OnMainLoad(f)
}

func OnAsyncLoad(f FRegister) {
	onLoads = append(onLoads, f)
}

// OnMainLoad used in golang code and running on main thread
func OnMainLoad(f FRegister) {
	onMainLoads = append(onMainLoads, f)
}

// OnUnload used in golang code when jvm OnUnload
func OnUnload(f func()) {
	onUnLoads = append(onUnLoads, f)
}

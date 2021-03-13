package jni

import "C"

import (
	"log"
)

func PushLocalFrame(nArgs int) Env {
	env := AutoGetCurrentThreadEnv()
	frameSize := 2*nArgs + 10
	if env.PushLocalFrame(frameSize) < 0 {
		log.Fatal("PushLocalFrame failed")
	}
	return env
}

func PopLocalFrame(env Env) {
	env.PopLocalFrame(0)
}

//export log_info
func log_info(v *C.char) {
	log.Println(C.GoString(v))
}

//export log_fatal
func log_fatal(v *C.char) {
	log.Fatal(C.GoString(v))
}

func SetVm(mVm VM) {

}

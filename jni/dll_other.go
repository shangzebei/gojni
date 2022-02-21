//go:build linux || android || darwin
// +build linux android darwin

package jni

/*
#cgo CFLAGS: -I ${SRCDIR}/include/
#cgo linux LDFLAGS: -ldl
#define _GNU_SOURCE
#include <dlfcn.h>
#include <jni.h>
#include <stdlib.h>
#include <pthread.h>
static JavaVM *jvm;
static pthread_key_t jnienvs;

extern void log_info(char *);
extern void log_fatal(char *);

static void env_destructor(void *env) {
	  if ((*jvm)->DetachCurrentThread(jvm) != JNI_OK) {
		log_fatal("failed to detach current thread");
	  }
}

static JNIEnv *go_seq_get_thread_env(void) {
	JNIEnv *env;
	if(jvm == 0){
		log_fatal("jvm is null");
	}
	jint ret = (*jvm)->GetEnv(jvm, (void **)&env, JNI_VERSION_1_6);
	if (ret != JNI_OK) {
		if (ret != JNI_EDETACHED) {
			log_fatal("failed to get thread env");
		}
		if ((*jvm)->AttachCurrentThread(jvm, (void **)&env, NULL) != JNI_OK) {
			log_fatal("failed to attach current thread");
		}
		pthread_setspecific(jnienvs, env);
	}
	return env;
}

static const char* getSoPath() {
	Dl_info dl_info;
	dladdr((void *)getSoPath, &dl_info);
    return dl_info.dli_fname;
}

static void init(JavaVM *j){
     jvm = j;
	 if (pthread_key_create(&jnienvs, env_destructor) != 0) {
	  	log_fatal("failed to initialize jnienvs thread local storage");
	 }
}
*/
import "C"
import (
	"unsafe"
)

func InitJNI(jvm uintptr) {
	SetVm(VM(jvm))
	C.init((*C.JavaVM)(unsafe.Pointer(jvm)))
}

func AutoGetCurrentThreadEnv() *Env {
	if wVm == 0 {
		panic("please invoke after onload")
	}
	e := Env(unsafe.Pointer(C.go_seq_get_thread_env()))
	return &e
}

func GetSelfPath() string {
	cs := C.GoString(C.getSoPath())
	return cs
}

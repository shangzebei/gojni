package jni

/*
#define __USE_GNU
#include <jni.h>
#include <stdlib.h>
#include <pthread.h>
//#include <dlfcn.h>
static JavaVM *jvm;
static pthread_key_t jnienvs;

extern void log_info(char *);
extern void log_fatal(char *);

static void env_destructor(void *env) {
	  if ((*jvm)->DetachCurrentThread(jvm) != JNI_OK) {
		log_info("failed to detach current thread");
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

static void init(JavaVM *j){
     //Dl_info dl_info;
     //dladdr((void *)init, &dl_info);
     //fprintf(stderr, "module %s loaded\n", dl_info.dli_fname);
     jvm = j;
	 if (pthread_key_create(&jnienvs, env_destructor) != 0) {
	   log_fatal("failed to initialize jnienvs thread local storage");
	 }
}
*/
import "C"

import (
	"log"
	"unsafe"
)

func InitJNI(jvm uintptr) {
	C.init((*C.JavaVM)(unsafe.Pointer(jvm)))
}

func AutoGetCurrentThreadEnv() Env {
	return Env(unsafe.Pointer(C.go_seq_get_thread_env()))
}

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

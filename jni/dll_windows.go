//+build windows

package jni

import (
	"reflect"
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

const (
	errnoERROR_IO_PENDING                        = 997
	GET_MODULE_HANDLE_EX_FLAG_PIN                = 0x00000001
	GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT = 0x00000002
	GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS       = 0x00000004
)

var (
	modkernel32            = syscall.NewLazyDLL("kernel32.dll")
	procGetModuleFileNameA = modkernel32.NewProc("GetModuleFileNameA")
	procGetModuleHandleExA = modkernel32.NewProc("GetModuleHandleExA")
)

func GetModuleFileNameW(hModule syscall.Handle, buf unsafe.Pointer, nSize uint32) (dWord uint32) {
	r0, _, _ := syscall.Syscall(procGetModuleFileNameA.Addr(), 3, uintptr(hModule), uintptr(buf), uintptr(nSize))
	dWord = uint32(r0)
	return
}

func GetModuleHandleExW(dwFlags uint32, lpModuleName unsafe.Pointer, hModule *syscall.Handle) (success bool) {
	r0, _, _ := syscall.Syscall(procGetModuleHandleExA.Addr(), 3, uintptr(dwFlags), uintptr(lpModuleName), uintptr(unsafe.Pointer(hModule)))
	success = r0 != 0
	return
}

func GetSelfPath() string {
	var buf [100]byte
	var hMd syscall.Handle
	var err = GetModuleHandleExW(GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS|GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT, unsafe.Pointer(reflect.ValueOf(GetSelfPath).Pointer()), &hMd)
	if !err {
		panic("windows GetSelfPath err")
	}
	size := GetModuleFileNameW(hMd, unsafe.Pointer(&buf), 100)
	return string(buf[0:size])
}

var wVm VM

func InitJNI(jvm uintptr) {
	wVm = VM(jvm)
}

func AutoGetCurrentThreadEnv() Env {
	if env, v := wVm.GetEnv(JNI_VERSION_1_6); v == JNI_OK {
		return env
	} else {
		panic("JNI_OnLoad error")
		return 0
	}
}

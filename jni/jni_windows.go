package jni

//
//import "syscall"
//
//var (
//	jvm                          *syscall.LazyDLL
//	AttachCurrentThread          *syscall.LazyProc
//	AttachCurrentThreadAsDaemon  *syscall.LazyProc
//	GetEnv                       *syscall.LazyProc
//	GetObjectRefType             *syscall.LazyProc
//	GetJavaVM                    *syscall.LazyProc
//	DestroyJavaVM                *syscall.LazyProc
//	DetachCurrentThread          *syscall.LazyProc
//	FindClass                    *syscall.LazyProc
//	GetVersion                   *syscall.LazyProc
//	FromReflectedMethod          *syscall.LazyProc
//	FromReflectedField           *syscall.LazyProc
//	ToReflectedMethod            *syscall.LazyProc
//	GetSuperclass                *syscall.LazyProc
//	IsAssignableFrom             *syscall.LazyProc
//	ToReflectedField             *syscall.LazyProc
//	Throw                        *syscall.LazyProc
//	ThrowNew                     *syscall.LazyProc
//	ExceptionOccurred            *syscall.LazyProc
//	ExceptionDescribe            *syscall.LazyProc
//	ExceptionClear               *syscall.LazyProc
//	FatalError                   *syscall.LazyProc
//	NewGlobalRef                 *syscall.LazyProc
//	DeleteGlobalRef              *syscall.LazyProc
//	DeleteLocalRef               *syscall.LazyProc
//	IsSameObject                 *syscall.LazyProc
//	NewObjectA                   *syscall.LazyProc
//	GetObjectClass               *syscall.LazyProc
//	IsInstanceOf                 *syscall.LazyProc
//	GetMethodID                  *syscall.LazyProc
//	CallObjectMethodA            *syscall.LazyProc
//	CallBooleanMethodA           *syscall.LazyProc
//	CallCharMethodA              *syscall.LazyProc
//	CallShortMethodA             *syscall.LazyProc
//	CallIntMethodA               *syscall.LazyProc
//	CallLongMethodA              *syscall.LazyProc
//	CallFloatMethodA             *syscall.LazyProc
//	CallDoubleMethodA            *syscall.LazyProc
//	CallVoidMethodA              *syscall.LazyProc
//	CallNonvirtualObjectMethodA  *syscall.LazyProc
//	CallNonvirtualBooleanMethodA *syscall.LazyProc
//	CallNonvirtualByteMethodA    *syscall.LazyProc
//	CallNonvirtualIntMethodA     *syscall.LazyProc
//	CallNonvirtualLongMethodA    *syscall.LazyProc
//	CallNonvirtualFloatMethodA   *syscall.LazyProc
//	CallNonvirtualDoubleMethodA  *syscall.LazyProc
//	CallNonvirtualCharMethodA    *syscall.LazyProc
//	GetFieldID                   *syscall.LazyProc
//	GetObjectField               *syscall.LazyProc
//	GetBooleanField              *syscall.LazyProc
//	GetByteField                 *syscall.LazyProc
//	GetCharField                 *syscall.LazyProc
//	GetShortField                *syscall.LazyProc
//	GetIntField                  *syscall.LazyProc
//	GetLongField                 *syscall.LazyProc
//	GetFloatField                *syscall.LazyProc
//	GetDoubleField               *syscall.LazyProc
//	SetObjectField               *syscall.LazyProc
//	SetBooleanField              *syscall.LazyProc
//	SetByteField                 *syscall.LazyProc
//	SetCharField                 *syscall.LazyProc
//	SetShortField                *syscall.LazyProc
//	SetIntField                  *syscall.LazyProc
//	SetLongField                 *syscall.LazyProc
//	SetFloatField                *syscall.LazyProc
//	SetDoubleField               *syscall.LazyProc
//)
//
//func init() {
//	test := syscall.NewLazyDLL("Win32Project1.dll")
//	AttachCurrentThread         = test.NewProc("AttachCurrentThread")
//	AttachCurrentThreadAsDaemon = test.NewProc("AttachCurrentThreadAsDaemon")
//	GetEnv                      = test.NewProc("GetEnv")
//	GetObjectRefType            = test.NewProc("GetObjectRefType")
//	GetJavaVM                   = test.NewProc("GetJavaVM")
//	DestroyJavaVM               = test.NewProc("DestroyJavaVM")
//	DetachCurrentThread         = test.NewProc("DetachCurrentThread")
//	FindClass                   = test.NewProc("FindClass")
//	GetVersion                  = test.NewProc("GetVersion")
//	FromReflectedMethod         = test.NewProc("FromReflectedMethod")
//	FromReflectedField          = test.NewProc("FromReflectedField")
//	ToReflectedMethod           = test.NewProc("ToReflectedMethod")
//	GetSuperclass               = test.NewProc("GetSuperclass")
//	IsAssignableFrom            = test.NewProc("IsAssignableFrom")
//	ToReflectedField            = test.NewProc("ToReflectedField")
//	Throw                       = test.NewProc("Throw")
//	ThrowNew                    = test.NewProc("ThrowNew")
//	ExceptionOccurred           = test.NewProc("ExceptionOccurred")
//	ExceptionDescribe           = test.NewProc("ExceptionDescribe")
//	ExceptionClear              = test.NewProc("ExceptionClear")
//	FatalError                  = test.NewProc("FatalError")
//	NewGlobalRef                = test.NewProc("NewGlobalRef")
//	DeleteGlobalRef             = test.NewProc("DeleteGlobalRef")
//	DeleteLocalRef              = test.NewProc("DeleteLocalRef")
//	IsSameObject                = test.NewProc("IsSameObject")
//	NewObjectA                  = test.NewProc("NewObjectA")
//	GetObjectClass              = test.NewProc("GetObjectClass")
//	IsInstanceOf                = test.NewProc("IsInstanceOf")
//	GetMethodID                 = test.NewProc("GetMethodID")
//	CallObjectMethodA           = test.NewProc("CallObjectMethodA")
//	CallBooleanMethodA          = test.NewProc("CallBooleanMethodA")
//	CallCharMethodA             = test.NewProc("CallCharMethodA")
//	CallShortMethodA            = test.NewProc("CallShortMethodA")
//	CallIntMethodA              = test.NewProc("CallIntMethodA")
//	CallLongMethodA             = test.NewProc("CallLongMethodA")
//	CallFloatMethodA            = test.NewProc("CallFloatMethodA")
//	CallDoubleMethodA           = test.NewProc("CallDoubleMethodA")
//	CallVoidMethodA             = test.NewProc("CallVoidMethodA")
//	CallNonvirtualObjectMethodA = test.NewProc("CallNonvirtualObjectMethodA")
//	CallNonvirtualBooleanMethodA= test.NewProc("CallNonvirtualBooleanMethodA")
//	CallNonvirtualByteMethodA   = test.NewProc("CallNonvirtualByteMethodA")
//	CallNonvirtualIntMethodA    = test.NewProc("CallNonvirtualIntMethodA")
//	CallNonvirtualLongMethodA   = test.NewProc("CallNonvirtualLongMethodA")
//	CallNonvirtualFloatMethodA  = test.NewProc("CallNonvirtualFloatMethodA")
//	CallNonvirtualDoubleMethodA = test.NewProc("CallNonvirtualDoubleMethodA")
//	CallNonvirtualCharMethodA   = test.NewProc("CallNonvirtualCharMethodA")
//	GetFieldID                  = test.NewProc("GetFieldID")
//	GetObjectField              = test.NewProc("GetObjectField")
//	GetBooleanField             = test.NewProc("GetBooleanField")
//	GetByteField                = test.NewProc("GetByteField")
//	GetCharField                = test.NewProc("GetCharField")
//	GetShortField               = test.NewProc("GetShortField")
//	GetIntField                 = test.NewProc("GetIntField")
//	GetLongField                = test.NewProc("GetLongField")
//	GetFloatField               = test.NewProc("GetFloatField")
//	GetDoubleField              = test.NewProc("GetDoubleField")
//	SetObjectField              = test.NewProc("SetObjectField")
//	SetBooleanField             = test.NewProc("SetBooleanField")
//	SetByteField                = test.NewProc("SetByteField")
//	SetCharField                = test.NewProc("SetCharField")
//	SetShortField               = test.NewProc("SetShortField")
//	SetIntField                 = test.NewProc("SetIntField")
//	SetLongField                = test.NewProc("SetLongField")
//	SetFloatField               = test.NewProc("SetFloatField")
//	SetDoubleField              = test.NewProc("SetDoubleField")
//}

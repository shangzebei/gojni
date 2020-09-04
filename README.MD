### use
```go
//export JNI_OnLoad
func JNI_OnLoad(vm uintptr, voidPtr uintptr) int {
	fmt.Println("JNI_OnLoad")
	JavaVM = jni.VM(vm)
	if _, v := jni.VM(vm).GetEnv(jni.JNI_VERSION_1_6); v != jni.JNI_OK {
		logrus.Error("exit vm")
		return -1
	}
	return jni.JNI_VERSION_1_6
}

//export Java_com_chain_cn_Miner_getHashRate
func Java_com_chain_cn_Miner_getHashRate(JNIEnv uintptr, jclass uintptr) uintptr {
	return jni.Env(JNIEnv).NewString(ethR.sHashRate)
}

```
### build
```
CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/linux" go build -tags 'jni' -buildmode=c-shared -o miner.so
```

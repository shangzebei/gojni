### use
import 
```go
import _ "gitee.com/aifuturewell/gojni/loader"
```
```go
//export Java_com_chain_cn_Miner_getHashRate
func Java_com_chain_cn_Miner_getHashRate(JNIEnv uintptr, jclass uintptr) uintptr {
	return jni.Env(JNIEnv).NewString(ethR.sHashRate)
}

```

```
### build
```
CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/linux" go build -tags 'jni' -buildmode=c-shared -o miner.so
```

###
{{$s=new java.lang.String}}

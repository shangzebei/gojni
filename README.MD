### use

```go
package main

import (
	"fmt"
	"unsafe"

	"gitee.com/aifuturewell/gojni/loader"
)

func main() {

}

func init() {
	loader.OnLoad(func(vm uintptr) {
		loader.Bind("com.szb.Hello").
			BindNative("stop", "void()", show).
			BindNative("args", "void(int[])", arg).
			BindNative("nice", "void(java.lang.String[])", nice).
			BindNative("bb", "void(byte[])", bys).
			BindNative("llll", "void(long[])", lllll).
			BindNative("fff", "void(float[])", fff).
			BindNative("ddd", "void(double[])", ddd).
			Done()
	})
}

func lllll(l []int) {
	fmt.Println(l)
}

func show() {
	fmt.Println("go show")
}

func arg(a []int32) {
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println("go out ", a[1])
}

func nice(ss []string) {
	fmt.Println(ss)
}

func bys(bb []byte) {
	fmt.Println(bb)
}

func fff(tt []float32) {
	fmt.Println(tt)
}

func ddd(ddd []float64) {
	fmt.Println(ddd)
}
```
### build
```
CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/linux" go build -tags 'jni' -buildmode=c-shared -o miner.so
```

###
{{$s=new java.lang.String}}

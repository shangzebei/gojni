### quick start
#### in go 
```go
package main
import (
	"fmt"
	"unsafe"
	"gitee.com/aifuturewell/gojni/loader"
)
func main() {}
func init() {
	loader.OnLoad(func(vm uintptr) {
		loader.WithClass("com.nk.Hello").
			BindNative("nice", "void(java.lang.String[])", nice).
			Done()
	})
}

func nice(ss []string) {
	fmt.Println(ss)
}
```
#### in java
```java
package com.nk;

public class Hello {
    static {
        System.loadLibrary("test");
    }
    public static void main(String[] args) {
        nice(new String[] { "come", "from", "gojni" });
    }
    public static native void nice(String[] sss);
}
```
### build and run
```
CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/linux" go build -buildmode=c-shared -o libtest.so
LD_LIBRARY_PATH=. java com.nk.Hello
```

###
```json
{{$s=new java.lang.String}}
```


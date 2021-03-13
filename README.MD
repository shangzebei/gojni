## quick start
### 1.java invoke native
in go 
```go
package main
import (
	"fmt"
	"unsafe"
	"gitee.com/aifuturewell/gojni/java"
)
func main() {}
func init() {
	java.OnLoad(func(reg java.Register) {
		reg.WithClass("com.nk.Hello").
			BindNative("nice", "void(java.lang.String[])", nice).
			Done()
	})
}

func nice(ss []string) {
	fmt.Println(ss)
}
```
in java
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

### 2.native invoke java
```go
native.LoadClass("com.szb.Jvm").New().Invoke("attach", "void()")

fmt.Println(vm.RunSource("java.lang.Thread.currentThread[java.lang.Thread()]();").
				AsObject().
				Invoke("getName", "java.lang.String()").
				AsString())
```

```json
{{$s=new java.lang.String}}
```


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
	java.OnMainLoad(func(reg java.Register) {
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
or
```shell
go env -w CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/$(go env GOHOSTOS) -g -O2"
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
# Android build
```
GOOS=android \
GOARCH=arm64 \
CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang \
CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang++ \
CGO_ENABLED=1 \
go build -buildmode=c-shared -o aaa.so
```

# IOS build
```
GOOS=darwin \
GOARCH=arm64 \
CC=/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang \
CXX=/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang++ \
CGO_CFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.7.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" \
CGO_CXXFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.7.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" \
CGO_LDFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.7.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" \
CGO_ENABLED=1 \
go build -tags ios -v -x -buildmode=c-archive -o aaa.a
```

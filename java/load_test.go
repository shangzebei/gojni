package java

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestCode(t *testing.T) {
	fmt.Println(unsafe.Sizeof(float32(0.0)))
}

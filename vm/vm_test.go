package vm

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/native"
	"testing"
)

func TestOjbHelp(t *testing.T) {
	var o native.ojbHeap
	o.Push(native.ojb{})
	fmt.Println(o.Len())
	o.Push(native.ojb{})
	fmt.Println(o.Len())

}

func TestFloat(t *testing.T) {

}

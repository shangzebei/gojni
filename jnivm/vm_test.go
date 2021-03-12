package jnivm

import (
	"fmt"
	"testing"
)

func TestOjbHelp(t *testing.T) {
	var o ojbHeap
	o.Push(ojb{})
	fmt.Println(o.Len())
	o.Push(ojb{})
	fmt.Println(o.Len())

}

func TestFloat(t *testing.T) {

}

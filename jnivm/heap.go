package jnivm

import (
	"gitee.com/aifuturewell/gojni/jni"
)

type ojb struct {
	varName string
	jCls    jni.Jclass
	jObj    jni.Jobject
}

//heap
type ojbHeap []ojb

func (h *ojbHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *ojbHeap) Len() int {
	return len(*h)
}

func (h *ojbHeap) Pop() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *ojbHeap) Push(v ojb) {
	*h = append(*h, v)
}

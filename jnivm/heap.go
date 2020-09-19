package jnivm

import (
	"gitee.com/aifuturewell/gojni/jni"
)

type ojb struct {
	jCls jni.Jclass
	jObj jni.Jobject
}

type ojbHelp []ojb

func (h *ojbHelp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *ojbHelp) Len() int {
	return len(*h)
}

func (h *ojbHelp) Pop() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *ojbHelp) Push(v ojb) {
	*h = append(*h, v)
}

package jparser

import (
	"testing"
)

var sigTest = map[string]string{
	"void(int)":                               "(I)V",
	"java.lang.String(int)":                   "(I)Ljava/lang/String;",
	"int(java.lang.String)":                   "(Ljava/lang/String;)I",
	"void(java.lang.String,java.lang.String)": "(Ljava/lang/String;Ljava/lang/String;)V",
	"void()":                  "()V",
	"void(int[])":             "([I)V",
	"java.lang.String[](int)": "(I)[Ljava/lang/String;",
}

func TestGetSig(t *testing.T) {
	for s, s2 := range sigTest {
		res := GetSig(s).String()
		if res != s2 {
			t.Fatalf("result [%s] != %s", res, s2)
		}
	}
}

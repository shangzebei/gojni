package utils

import (
	"fmt"
	"testing"
)

var sigTest = map[string]string{
	"void(int,int)":                           "(II)V",
	"java.lang.String(int)":                   "(I)Ljava/lang/String;",
	"int(java.lang.String,int)":               "(Ljava/lang/String;I)I",
	"void(java.lang.String,java.lang.String)": "(Ljava/lang/String;Ljava/lang/String;)V",
	"void()":                   "()V",
	"void(int[])":              "([I)V",
	"java.lang.String[](int)":  "(I)[Ljava/lang/String;",
	"void(java.lang.String[])": "([Ljava/lang/String;)V",
}

func TestGetSig(t *testing.T) {
	for s, s2 := range sigTest {
		res := GetSig(s)
		fmt.Printf("%s %s \n", s, res.ParamTyp)
		if res.Sig != s2 {
			t.Fatalf("result {%s} != %s", res.Sig, s2)
		}
	}
}

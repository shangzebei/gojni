package native

import (
	"fmt"
	"testing"
)

func TestSig(t *testing.T) {
	fmt.Println(NewSig("com.android.thread").GetSigType())
	fmt.Println(NewSig("void").GetSigType())
	fmt.Println(NewSig("int[]").GetSigType())
	fmt.Println(NewSig("com.android.thread[]").GetSigType())
	fmt.Println(SigOf("Lcom/android/thread;").GetType())
	fmt.Println(SigOf("[Lcom/android/thread;").GetType())
}

func TestSigned(t *testing.T) {
	//int(java.lang.String)
	fmt.Println(EncodeToSig("int(java.lang.String)").ParamTyp)
	fmt.Println(EncodeToSig("void(java.lang.String)"))
}

var sigTest = map[string]string{
	"void(int,int)":                           "(II)V",
	"java.lang.String(int)":                   "(I)Ljava/lang/String;",
	"int(java.lang.String,int)":               "(Ljava/lang/String;I)I",
	"void(java.lang.String,java.lang.String)": "(Ljava/lang/String;Ljava/lang/String;)V",
	"void()":                   "()V",
	"void(int[])":              "([I)V",
	"java.lang.String[](int)":  "(I)[Ljava/lang/String;",
	"void(java.lang.String[])": "([Ljava/lang/String;)V",
	"int[]()":                  "()[I",
}

func TestEncodeToSig(t *testing.T) {
	for s, s2 := range sigTest {
		res := EncodeToSig(s)
		fmt.Printf("%s %s \n", s, res.ParamTyp)
		if res.Sig != s2 {
			t.Fatalf("result {%s} != %s", res.Sig, s2)
		}
	}
}

package native

import (
	"fmt"
	"testing"
)

func TestSig(t *testing.T) {
	fmt.Println(NewSig("com.android.thread").GetSigType())
	fmt.Println(NewSig("void").GetSigType())
	fmt.Println(SigOf("Lcom/android/thread;").GetType())
}

func TestSigned(t *testing.T) {
	//int(java.lang.String)
	fmt.Println(EncodeToSig("int(java.lang.String)").ParamTyp)
}

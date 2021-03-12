package utils

import (
	"fmt"
	"strings"
)

//jni sig
var SigEncodeMap = map[string]string{
	"int":     "I",
	"boolean": "Z",
	"byte":    "B",
	"char":    "C",
	"short":   "S",
	"long":    "J",
	"float":   "F",
	"double":  "D",
	"void":    "V",
}

var SigDecodeMap = map[string]string{
	"I": "int",
	"Z": "boolean",
	"B": "byte",
	"C": "char",
	"S": "short",
	"J": "long",
	"F": "float",
	"D": "double",
	"V": "void",
}

type MethodSig struct {
	RetTyp   string
	ParamTyp []string
	Sig      string
}

func (a MethodSig) String() string {
	return fmt.Sprintf("sig %s", a.Sig)
}

//example void()
func EncodeToSig(oSig string) *MethodSig { //(I)V
	m := MethodSig{}
	if oSig == "" {
		m.RetTyp = "VOID"
		m.Sig = "()V"
		return &m
	}
	pos := strings.Index(oSig, "(")
	ret := oSig[:pos]
	retV := ""
	if strings.Contains(ret, "[]") {
		retV = "["
		ret = strings.ReplaceAll(ret, "[]", "")
	}
	if v, b := SigEncodeMap[ret]; b {
		retV += v
	} else if strings.Contains(ret, ".") { //class
		retV += "L" + strings.ReplaceAll(ret, ".", "/") + ";"
	} else {
		panic(fmt.Sprintf("not find ret sig [%s] orig %s", ret, oSig))
	}
	m.RetTyp = retV
	inputV := ""
	input := oSig[pos+1 : len(oSig)-1]
	if input != "" {
		kk := strings.Split(input, ",")
		//l := len(kk) - 1
		for i := 0; i < len(kk); i++ {
			value := kk[i]
			temp := ""
			if strings.Contains(value, "[]") {
				temp = "["
				value = strings.ReplaceAll(value, "[]", "")
			}
			if v, b := SigEncodeMap[value]; b {
				temp += v
			} else {
				s := strings.ReplaceAll(value, ".", "/") + ";"
				temp += "L" + s

			}
			m.ParamTyp = append(m.ParamTyp, temp)
			inputV += temp
		}
	}
	m.Sig = fmt.Sprintf("(%s)%s", inputV, retV)
	return &m
}

/*
	stop ()V
	args ([I)V
	nice ([Ljava/lang/String;)V
	bb ([B)V
	llll ([J)V
	fff ([F)V
	ddd ([D)V
*/
func SigToJavaNative(name string, sig string) {

}


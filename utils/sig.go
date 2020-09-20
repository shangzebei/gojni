package utils

import (
	"fmt"
	"strings"
)

var SigMap = map[string]string{
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

var SV = map[string]string{
	"I": "Int",
	"Z": "Boolean",
	"B": "Byte",
	"C": "Char",
	"S": "Short",
	"J": "Float",
	"F": "Float",
	"D": "Double",
	"V": "Void",
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
func GetSig(oSig string) *MethodSig { //(I)V
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
	if v, b := SigMap[ret]; b {
		retV += v
	} else if strings.Contains(ret, ".") { //class
		retV += "L" + strings.ReplaceAll(ret, ".", "/") + ";"
	} else {
		panic(fmt.Sprintf("not find ret sig %s", ret))
	}
	m.RetTyp = retV
	inputV := ""
	input := oSig[pos+1 : len(oSig)-1]
	if input != "" {
		kk := strings.Split(input, ",")
		//l := len(kk) - 1
		for i := 0; i < len(kk); i++ {
			value := kk[i]
			isArray := false
			if strings.Contains(value, "[]") {
				inputV = "["
				value = strings.ReplaceAll(value, "[]", "")
				isArray = true
			}
			if v, b := SigMap[value]; b {
				if isArray {
					m.ParamTyp = append(m.ParamTyp, "["+v)
				} else {
					m.ParamTyp = append(m.ParamTyp, v)
				}
				inputV += v
				//if i != l {
				//	inputV += ","
				//}
			} else {
				s := strings.ReplaceAll(value, ".", "/") + ";"
				inputV += "L" + s
				m.ParamTyp = append(m.ParamTyp, s)
			}
		}
	}
	m.Sig = fmt.Sprintf("(%s)%s", inputV, retV)
	return &m
}

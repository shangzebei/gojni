package jparser

import (
	"fmt"
	"strings"
)

var sigMap = map[string]string{
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

var sv = map[string]string{
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

func GetSig(oSig string) MethodSig { //(I)V
	m := MethodSig{}
	if oSig == "" {
		m.RetTyp = "VOID"
		m.Sig = "()V"
		return m
	}
	pos := strings.Index(oSig, "(")
	ret := oSig[:pos]
	retV := ""
	if strings.Contains(ret, "[]") {
		retV = "["
		ret = strings.ReplaceAll(ret, "[]", "")
	}
	if v, b := sigMap[ret]; b {
		retV += v
	} else { //class
		retV += "L" + strings.ReplaceAll(ret, ".", "/") + ";"
	}
	m.RetTyp = retV
	inputV := ""
	input := oSig[pos+1 : len(oSig)-1]
	if input != "" {
		kk := strings.Split(input, ",")
		l := len(kk) - 1
		for i := 0; i < len(kk); i++ {
			value := kk[i]
			if strings.Contains(value, "[]") {
				inputV = "["
				value = strings.ReplaceAll(value, "[]", "")
			}
			if v, b := sigMap[value]; b {
				inputV += v
				if i != l {
					inputV += ","
				}
			} else {
				inputV += "L" + strings.ReplaceAll(value, ".", "/") + ";"
			}
		}
	}
	m.Sig = fmt.Sprintf("(%s)%s", inputV, retV)
	return m
}

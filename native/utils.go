package native

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/jni"
	"gitee.com/aifuturewell/gojni/utils"
	"strings"
)

//example void()
func EncodeToSig(oSig string) *MethodSig { //(I)V
	var retTyp, sig string
	var paramTyp []Sig
	if oSig == "" {
		retTyp = "VOID"
		sig = "()V"
		return &MethodSig{
			RetTyp:   *SigOf(retTyp),
			ParamTyp: paramTyp,
			Sig:      sig,
		}
	}
	pos := strings.Index(oSig, "(")
	ret := oSig[:pos]
	retV := ""
	if strings.Contains(ret, "[]") {
		retV = "["
		ret = strings.ReplaceAll(ret, "[]", "")
	}
	if v, b := utils.SigEncodeMap[ret]; b {
		retV += v
	} else if strings.Contains(ret, ".") { //class
		retV += "L" + strings.ReplaceAll(ret, ".", "/") + ";"
	} else {
		jni.ThrowException(fmt.Sprintf("not find ret sig [%s] orig %s", ret, oSig))
	}
	retTyp = retV
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
			if v, b := utils.SigEncodeMap[value]; b {
				temp += v
			} else {
				s := strings.ReplaceAll(value, ".", "/") + ";"
				temp += "L" + s

			}
			paramTyp = append(paramTyp, *SigOf(temp))
			inputV += temp
		}
	}
	sig = fmt.Sprintf("(%s)%s", inputV, retV)
	return &MethodSig{
		RetTyp:   *SigOf(retTyp),
		ParamTyp: paramTyp,
		Sig:      sig,
	}
}

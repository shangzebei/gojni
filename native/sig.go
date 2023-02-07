package native

import (
	"fmt"
	"github.com/shangzebei/gojni/jni"
	"github.com/shangzebei/gojni/utils"
	"strings"
)

type Sig struct {
	tName string
	sName string
}

func NewSig(typeName string) *Sig {
	var styp string
	if s, b := utils.SigEncodeMap[typeName]; b {
		styp = s
	} else if strings.HasSuffix(typeName, "[]") {
		pkg := strings.ReplaceAll(typeName[:len(typeName)-2], ".", "/")
		styp = "[L" + pkg + ";"
	} else {
		pkg := strings.ReplaceAll(typeName, ".", "/")
		styp = "L" + pkg + ";"
	}
	return &Sig{
		tName: typeName,
		sName: styp,
	}
}

// SigOf Lcom/android/thread;
func SigOf(siged string) *Sig {
	var otyp string
	if s, b := utils.SigDecodeMap[siged]; b {
		otyp = s
	} else {
		//object
		if siged[0] == 'L' { //Ljava/lang/String; => java.lang.String
			otyp = strings.ReplaceAll(siged[1:len(siged)-1], "/", ".")
		} else if siged[0] == '[' { //[Ljava/lang/String; => java.lang.string[]
			otyp = strings.ReplaceAll(siged[2:len(siged)-1], "/", ".") + "[]"
		} else {
			jni.ThrowException(fmt.Sprintf("decode err %s", siged))
		}
	}
	return &Sig{
		tName: otyp,
		sName: siged,
	}
}

func (g *Sig) GetType() string {
	return g.tName
}

func (g *Sig) GetSigType() string {
	return g.sName
}

func (a *Sig) String() string {
	return fmt.Sprintf("Sig %s", a.sName)
}

type MethodSig struct {
	RetTyp   Sig
	ParamTyp []Sig
	Sig      string
}

func (a MethodSig) String() string {
	return fmt.Sprintf("Sig %s", a.Sig)
}

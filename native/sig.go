package native

import (
	"fmt"
	"gitee.com/aifuturewell/gojni/utils"
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
	} else {
		pkg := strings.ReplaceAll(typeName, ".", "/")
		styp = "L" + pkg + ";"
	}
	return &Sig{
		tName: typeName,
		sName: styp,
	}
}

//Lcom/android/thread;
func SigOf(siged string) *Sig {
	var otyp string
	if s, b := utils.SigDecodeMap[siged]; b {
		otyp = s
	} else {
		if siged[0] == 'L' && strings.HasSuffix(siged, ";") {
			otyp = strings.ReplaceAll(siged[1:len(siged)-1], "/", ".")
		} else {
			panic(fmt.Errorf("decode err %s", siged))
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
